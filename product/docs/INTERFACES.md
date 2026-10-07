# MemoGuard interface contract

This is the first implementation target. Package names and flag spelling may evolve before v0.1.0, but the behavior should stay consistent.

## Go engine

```go
type Input struct {
    Kind    string // envelope_xdr, memo_text, or soroban_simulation
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
```

Planned exit codes: `0` clean, `1` blocking finding, `2` invalid input or policy, `3` internal error. Warnings do not fail by default; a strict mode may opt in later.

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

The Action receives an input path, policy path, and CLI version. It downloads the matching CLI release, verifies its checksum, runs the scan, and emits annotations from the JSON report. It must not include input values in annotations.
