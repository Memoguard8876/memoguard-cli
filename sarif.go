package cli

import (
	"encoding/json"
	"io"
	"sort"

	engine "github.com/memoguard8876/memoguard-engine"
	rules "github.com/memoguard8876/memoguard-rules"
)

// SARIF contains locations within decoded transaction data, not byte offsets in
// the source XDR. Matched values and the original input are intentionally absent.
func writeSARIF(writer io.Writer, report engine.Report) error {
	type text struct {
		Text string `json:"text"`
	}
	type rule struct {
		ID               string `json:"id"`
		ShortDescription text   `json:"shortDescription"`
	}
	type location struct {
		LogicalLocations []struct {
			FullyQualifiedName string `json:"fullyQualifiedName"`
		} `json:"logicalLocations"`
	}
	type result struct {
		RuleID    string     `json:"ruleId"`
		Level     string     `json:"level"`
		Message   text       `json:"message"`
		Locations []location `json:"locations"`
	}
	ruleByID := make(map[string]rule)
	results := make([]result, 0, len(report.Findings))
	for _, finding := range report.Findings {
		ruleByID[finding.RuleID] = rule{ID: finding.RuleID, ShortDescription: text{Text: finding.Description}}
		level := "warning"
		if finding.Severity == rules.Block {
			level = "error"
		}
		loc := location{}
		loc.LogicalLocations = append(loc.LogicalLocations, struct {
			FullyQualifiedName string `json:"fullyQualifiedName"`
		}{FullyQualifiedName: finding.FieldPath})
		results = append(results, result{RuleID: finding.RuleID, Level: level,
			Message: text{Text: finding.Description + ". " + finding.Remediation}, Locations: []location{loc}})
	}
	ids := make([]string, 0, len(ruleByID))
	for id := range ruleByID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	rulesList := make([]rule, 0, len(ids))
	for _, id := range ids {
		rulesList = append(rulesList, ruleByID[id])
	}
	document := struct {
		Schema  string `json:"$schema"`
		Version string `json:"version"`
		Runs    []struct {
			Tool struct {
				Driver struct {
					Name  string `json:"name"`
					Rules []rule `json:"rules"`
				} `json:"driver"`
			} `json:"tool"`
			Results []result `json:"results"`
		} `json:"runs"`
	}{Schema: "https://json.schemastore.org/sarif-2.1.0.json", Version: "2.1.0"}
	document.Runs = append(document.Runs, struct {
		Tool struct {
			Driver struct {
				Name  string `json:"name"`
				Rules []rule `json:"rules"`
			} `json:"driver"`
		} `json:"tool"`
		Results []result `json:"results"`
	}{})
	document.Runs[0].Tool.Driver.Name = "MemoGuard"
	document.Runs[0].Tool.Driver.Rules = rulesList
	document.Runs[0].Results = results
	return json.NewEncoder(writer).Encode(document)
}
