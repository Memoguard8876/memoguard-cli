package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestBlockAndRedact(t *testing.T) {
	var output, errors bytes.Buffer
	code := Run(context.Background(), []string{"scan", "--memo", "person@example.com", "--format", "json"}, strings.NewReader(""), &output, &errors)
	if code != ExitBlocked {
		t.Fatalf("exit=%d stderr=%s", code, errors.String())
	}
	if strings.Contains(output.String(), "person@example.com") || !strings.Contains(output.String(), "transaction.memo.text") {
		t.Fatalf("unsafe or incomplete output: %s", output.String())
	}
}

func TestClean(t *testing.T) {
	var output, errors bytes.Buffer
	code := Run(context.Background(), []string{"scan", "--memo", "reference-123"}, strings.NewReader(""), &output, &errors)
	if code != ExitClean || !strings.Contains(output.String(), "clean") {
		t.Fatalf("exit=%d stdout=%s stderr=%s", code, output.String(), errors.String())
	}
}

func TestInvalidInput(t *testing.T) {
	var output, errors bytes.Buffer
	code := Run(context.Background(), []string{"scan", "--memo", "test", "--xdr", "file"}, strings.NewReader(""), &output, &errors)
	if code != ExitInput {
		t.Fatalf("exit=%d", code)
	}
}

func TestFailureThresholdAndSARIFRedaction(t *testing.T) {
	for _, test := range []struct {
		name, memo, threshold string
		wantExit              int
	}{
		{"warning does not block by default", "+12125550123", "block", ExitClean},
		{"warning threshold blocks", "+12125550123", "warning", ExitBlocked},
		{"none disables failure", "person@example.com", "none", ExitClean},
	} {
		t.Run(test.name, func(t *testing.T) {
			var output, errors bytes.Buffer
			code := Run(context.Background(), []string{"scan", "--memo", test.memo, "--format", "sarif", "--fail-on", test.threshold}, strings.NewReader(""), &output, &errors)
			if code != test.wantExit {
				t.Fatalf("exit=%d stderr=%s", code, errors.String())
			}
			if strings.Contains(output.String(), test.memo) {
				t.Fatal("SARIF disclosed the matched value")
			}
			var doc map[string]any
			if err := json.Unmarshal(output.Bytes(), &doc); err != nil || doc["version"] != "2.1.0" {
				t.Fatalf("invalid SARIF: %v", err)
			}
		})
	}
}
