package cli

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	engine "github.com/memoguard8876/memoguard-engine"
	rules "github.com/memoguard8876/memoguard-rules"
)

const (
	ExitClean   = 0
	ExitBlocked = 1
	ExitInput   = 2
	ExitFailure = 3
)

func Run(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) == 0 || args[0] != "scan" {
		fmt.Fprintln(stderr, "usage: memoguard scan [--xdr file | --memo text | --simulation file | --json file] [--policy file] [--format human|json]")
		return ExitInput
	}
	flags := flag.NewFlagSet("scan", flag.ContinueOnError)
	flags.SetOutput(stderr)
	xdrPath := flags.String("xdr", "", "base64 transaction envelope XDR file, or - for stdin")
	memo := flags.String("memo", "", "memo text")
	simulationPath := flags.String("simulation", "", "Stellar RPC simulateTransaction JSON file")
	jsonPath := flags.String("json", "", "decoded public-transaction JSON file")
	policyPath := flags.String("policy", "", "policy JSON file")
	format := flags.String("format", "human", "human or json")
	if err := flags.Parse(args[1:]); err != nil || flags.NArg() != 0 {
		return ExitInput
	}
	provided := 0
	for _, candidate := range []string{*xdrPath, *memo, *simulationPath, *jsonPath} {
		if candidate != "" {
			provided++
		}
	}
	if provided != 1 || (*format != "human" && *format != "json") {
		fmt.Fprintln(stderr, "choose exactly one input and --format human or json")
		return ExitInput
	}
	policy := rules.Default()
	if *policyPath != "" {
		file, err := os.Open(*policyPath)
		if err != nil {
			fmt.Fprintln(stderr, "cannot open policy file")
			return ExitInput
		}
		defer file.Close()
		policy, err = rules.Load(file)
		if err != nil {
			fmt.Fprintln(stderr, "invalid policy:", err)
			return ExitInput
		}
	}
	input, err := readInput(*xdrPath, *memo, *simulationPath, *jsonPath, stdin)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return ExitInput
	}
	scanner, err := engine.New(policy)
	if err != nil {
		fmt.Fprintln(stderr, "invalid policy:", err)
		return ExitInput
	}
	report, err := scanner.Scan(ctx, input)
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			fmt.Fprintln(stderr, "scan canceled")
			return ExitFailure
		}
		fmt.Fprintln(stderr, "invalid input:", err)
		return ExitInput
	}
	if *format == "json" {
		if err := json.NewEncoder(stdout).Encode(report); err != nil {
			fmt.Fprintln(stderr, "cannot write report")
			return ExitFailure
		}
	} else if err := writeHuman(stdout, report); err != nil {
		fmt.Fprintln(stderr, "cannot write report")
		return ExitFailure
	}
	if report.Blocked() {
		return ExitBlocked
	}
	return ExitClean
}

func readInput(xdrPath, memo, simulationPath, jsonPath string, stdin io.Reader) (engine.Input, error) {
	if memo != "" {
		return engine.Input{Kind: engine.MemoText, Payload: []byte(memo)}, nil
	}
	kind := engine.EnvelopeXDR
	path := xdrPath
	if simulationPath != "" {
		kind, path = engine.SorobanSimulation, simulationPath
	}
	if jsonPath != "" {
		kind, path = engine.DecodedJSON, jsonPath
	}
	var reader io.Reader = stdin
	if path != "-" {
		file, err := os.Open(path)
		if err != nil {
			return engine.Input{}, errors.New("cannot open input file")
		}
		defer file.Close()
		reader = file
	}
	data, err := io.ReadAll(io.LimitReader(reader, engine.MaxInputBytes+1))
	if err != nil {
		return engine.Input{}, errors.New("cannot read input")
	}
	if len(data) > engine.MaxInputBytes {
		return engine.Input{}, errors.New("input exceeds size limit")
	}
	return engine.Input{Kind: kind, Payload: data}, nil
}

func writeHuman(writer io.Writer, report engine.Report) error {
	if len(report.Findings) == 0 {
		_, err := fmt.Fprintln(writer, "MemoGuard: clean")
		return err
	}
	if _, err := fmt.Fprintf(writer, "MemoGuard: %d finding(s)\n", len(report.Findings)); err != nil {
		return err
	}
	for _, finding := range report.Findings {
		if _, err := fmt.Fprintf(writer, "%s %s at %s: %s. %s\n",
			strings.ToUpper(string(finding.Severity)), finding.RuleID, finding.FieldPath,
			finding.Description, finding.Remediation); err != nil {
			return err
		}
	}
	return nil
}
