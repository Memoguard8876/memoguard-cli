# MemoGuard product requirements

Status: planning baseline, 7 October 2026

## Problem

Stellar transaction data is public and durable. A team can accidentally place a person's email address, phone number, customer reference, or other private data in a memo or contract call. Once submitted, a warning after the fact cannot remove it. Existing generic secret scanners do not understand Stellar XDR structures or which fields can become public.

## Product promise

MemoGuard gives developers a clear pass or fail **before submission**. It points to the unsafe field and explains the rule without repeating the private value in logs.

## Primary users

- Stellar application teams building payment or contract integrations.
- CI and security teams reviewing transaction templates and fixture files.
- SDK maintainers who want to add a pre-submit check to a Go service.

## Jobs to be done

1. Before sending a transaction, inspect its decoded public fields for accidental private information.
2. In CI, reject a fixture or generated XDR file that would expose private data.
3. Explain exactly which rule fired and where, so the developer can correct the source.
4. Allow documented exceptions for deliberate public data without disabling the whole scan.

## MVP scope

### Required inputs

- Base64 Stellar transaction-envelope XDR.
- A plain memo or decoded transaction JSON for local development.
- Soroban invocation arguments and simulation output when the caller provides them.

### Required behavior

- Decode XDR using the official Stellar Go library and scan only fields that would be public on submission.
- Apply a versioned policy pack for common personal data patterns and user-defined patterns.
- Return finding ID, severity, field path, safe description, and suggested fix.
- Never include matched raw private values in normal reports, CI annotations, or telemetry.
- Exit with a nonzero code when a blocking finding is present.
- Support JSON and human-readable output; SARIF can follow once the core path is stable.
- Provide a reusable Go API and a GitHub Action that invokes the released CLI.

### Initial policy categories

- Email addresses.
- Phone numbers, with region-aware tuning to reduce false positives.
- Government or customer IDs only when explicitly configured, because formats vary.
- Credentials or tokens accidentally inserted into memo or contract input.

Rules must identify a field and a confidence level. A raw regex match alone is not sufficient to block a payment without policy context.

## Explicit boundaries

- MemoGuard cannot erase data already submitted to Stellar.
- Runtime contract events are only inspectable before submission if the team provides simulation results; real events already emitted on-chain can only be audited.
- MemoGuard does not promise legal compliance or complete detection of every private data format.
- No website, wallet, custody, or contract deployment is required for the MVP.

## User stories and acceptance criteria

| Story | Acceptance criteria |
| --- | --- |
| Scan a transaction before sending it | Valid XDR is decoded; a known sensitive memo yields a blocking finding with its field path; clean XDR passes. |
| Review findings safely in CI | Action fails on blocking findings; annotation contains rule and location but not the matched value. |
| Add an organization rule | A versioned policy file loads, validates, and rejects invalid or duplicate rule IDs. |
| Allow a deliberate exception | An allowlist entry requires a rule ID, scope, reason, and expiry; it suppresses only the matching finding. |
| Handle malformed input | CLI returns a clear parsing error and a distinct exit code without panicking. |

## Success measures

- A pilot team catches at least one intentional test leak before submission.
- No matched private value appears in standard CLI or Action output in test fixtures.
- The same input and policy version produce the same findings across CLI and Go API.
- Developers can add the Action to a repository with a short workflow file and no hosted account.

## Open questions to resolve before implementation

- Which Stellar Go SDK version and XDR envelope types to support first.
- Whether policy files use YAML or JSON; JSON is the default until user research suggests otherwise.
- Which simulation artifact format should be accepted for Soroban event checks.
- Whether any pilot team needs language bindings beyond Go.
