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

Before v0.1.0, run Go tests, `go vet`, fuzz smoke tests, and an end-to-end Action workflow against a synthetic unsafe fixture. A pilot team should review false positives before enabling blocking in production CI.
