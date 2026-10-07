# MemoGuard architecture

## Data flow

```text
XDR / memo / simulation artifact
              │
              ▼
      parse and normalize
              │
              ▼
   public-field inventory
              │
              ▼
     versioned rule policy
              │
              ▼
   safe structured findings
          ┌───┴────┐
          ▼        ▼
        CLI     Go library
          │
          ▼
    GitHub Action
```

## Core contracts

The rules package exports a `Policy`, rule IDs, severities, and validation. The engine accepts a typed input plus a policy and returns a report. Every finding has a stable rule ID, severity, Stellar field path, and remediation hint. It must not carry the matched value by default.

The CLI accepts `scan --xdr <file>` or stdin, `--policy <file>`, and an output mode. Exit codes should distinguish clean, blocked, invalid input, and internal failure. The Action maps CLI findings to GitHub annotations and uses the same release binary; it must not reimplement detection.

## Stellar specifics

- Decode transaction-envelope XDR with the official Go XDR types. Scan the memo and any operation fields that become public.
- Treat opaque byte arrays carefully: decode text only when the protocol meaning and encoding are known.
- A Soroban simulation result is a separate input type. Simulated events and return values can be checked before submission when available.
- A historic event scan is an audit mode only. It cannot prevent publication and must be labeled accordingly.
- Never submit a transaction, hold keys, or require a wallet inside the scanner.

## Policy lifecycle

Built-in policies are versioned and tested against representative safe and unsafe fixtures. Teams can extend them with local rules. An exception must be narrow: rule ID, exact field scope, reason, and expiration date. A broad `disable-all` option must never be a default workflow.

## Output contract

JSON is the stable machine interface. Human output is optimized for local use. Both are derived from the same report. CI annotations include a rule, field path, and fix. Matched content is redacted. Future SARIF output should use the same finding IDs.

## Reliability

Scanning is deterministic and offline. No network request is required to inspect input. Large or malformed XDR must fail safely with bounded memory use, input-size limits, and parser errors rather than panics.
