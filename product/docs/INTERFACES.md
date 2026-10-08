# MemoGuard interface contract

This describes the released v0.2 interface. The Go packages remain the source of truth for exact types.

## Go engine

```go
type Input struct {
    Kind    Kind // envelope_xdr, memo_text, soroban_simulation, or decoded_json
    Payload []byte
}

type Finding struct {
    RuleID      string
    Severity    string
    Confidence  string
    FieldPath   string
    Description string
    Remediation string
}

type Report struct {
    PolicyVersion string
    Findings      []Finding
}
```

Matched raw content is deliberately absent from `Finding`. The engine must be deterministic for a fixed input and policy version.

## CLI

```text
memoguard scan --xdr transaction.xdr --format json
memoguard scan --memo "invoice 123" --format human
memoguard scan --simulation simulation.json --policy policy.json
memoguard scan --xdr transaction.xdr --format sarif --fail-on warning
```

Exit codes: `0` no finding at the selected threshold, `1` threshold met, `2` invalid or incomplete input or policy, `3` operational error. Warnings do not fail by default; `--fail-on warning` opts in, and `--fail-on none` emits findings without failing.

## JSON report example

```json
{
  "policy_version": "v1",
  "findings": [
    {
      "rule_id": "personal.email",
      "severity": "block",
      "confidence": "high",
      "field_path": "transaction.memo.text",
      "description": "Email address in a public memo",
      "remediation": "Use an opaque reference instead of the email address"
    }
  ]
}
```

## GitHub Action

The Action receives an input path, kind, optional policy, and failure threshold. Its release pins one CLI version and its platform SHA-256 digests. It downloads the public binary, verifies the digest, runs the scan, and emits annotations from the JSON report. It does not include input values in annotations.
