# MemoGuard build plan

## Phase 0 — contracts and fixtures

- Finalize Go module paths, license, support policy, and baseline CI.
- Create sanitized fixture envelopes covering memo types, operations, malformed XDR, and Soroban simulation artifacts.
- Lock the finding JSON schema and exit-code table.

**Done when:** all four repositories have purpose-specific READMEs, CI, and a local workspace that builds without unpublished dependency hacks.

## Phase 1 — rules and engine

- Implement policy parsing/validation and the initial rule pack in `memoguard-rules`.
- Implement XDR extraction and scanning in `memoguard-engine`.
- Add tests for redaction, false positives, malformed inputs, and common Stellar transaction shapes.

**Done when:** a Go caller receives deterministic findings from an unsafe transaction and a clean report from a safe one.

## Phase 2 — CLI

- Implement file and stdin inputs, JSON/human output, policy flag, and stable exit codes.
- Publish Go binaries for supported operating systems.

**Done when:** a developer can scan a transaction locally with one command and CI can read the JSON report.

## Phase 3 — GitHub Action and pilot

- Implement Action manifest and pinned CLI installation.
- Turn findings into safe annotations and fail the job on blocking rules.
- Run a pilot with a real Stellar development repository and tune noisy rules.

**Done when:** CI catches a deliberate test leak without printing the sensitive value.

## Definition of done for every repository

Code, tests, README usage example, changelog entry, CI checks, and tagged release. A downstream repository upgrades only to a released upstream version.
