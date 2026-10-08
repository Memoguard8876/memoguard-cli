# MemoGuard test plan

## Fixtures

Use synthetic data only. Include safe and unsafe Stellar transaction envelopes for text, ID, hash, and return memos; operation fields; Soroban invocation data; and malformed XDR. Store expected findings by stable rule ID and field path.

## Required checks

| Area | Check |
| --- | --- |
| Detection | Unsafe memo and supported public fields produce the expected finding. |
| Safe input | Opaque references and ordinary transaction data do not trigger blocking rules. |
| Redaction | Raw matched values never appear in JSON, human output, annotations, or error text. |
| Parsing | Truncated, oversized, and malformed XDR return controlled errors. |
| Policy | Invalid regex, duplicate rule IDs, expired exceptions, and broad exceptions are rejected or handled explicitly. |
| Consistency | Engine, CLI, and Action agree on the same report and blocking status. |
| Fuzzing | Parser and rule evaluation do not panic on arbitrary bytes. |

Current automated coverage includes stored synthetic clean and unsafe text-memo XDR, generated ID/hash/return memo cases, ManageData, simulation events and results, malformed XDR, oversized and deeply nested input, policy loading, CLI redaction and exit codes, and a three-OS Action fixture matrix. Soroban invoke/constructor fixtures, broader operation shapes, opaque-field coverage reporting, and a consented false-positive pilot remain open. Run Go tests, `go vet`, the fuzz smoke targets, and the Action integration workflow before each release. A pilot team should review false positives before enabling blocking in production CI.
