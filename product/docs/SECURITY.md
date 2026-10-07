# MemoGuard security and privacy requirements

MemoGuard handles data that may already be sensitive. The scanner must minimize what it stores and reveals.

## Required controls

- Run the scan locally by default; do not send transaction contents to a hosted API.
- Never log raw matched values, full XDR, secrets, or environment variables in standard output.
- Do not persist input unless the user explicitly chooses an output file.
- Bound input size, rule count, regex complexity, and execution time.
- Use vetted parsers for Stellar XDR and reject malformed data cleanly.
- Sign release artifacts where practical and publish checksums for CLI binaries.
- Pin the CLI version and checksum in the GitHub Action.
- Keep Action permissions at `contents: read` unless a specific feature needs more.

## Threats to test

| Threat | Control |
| --- | --- |
| Sensitive value leaks through an error or CI annotation | Redacted report contract and regression fixtures |
| Malicious rule causes excessive processing | Rule validation, limits, timeout tests |
| Crafted XDR crashes parser | Fuzz tests and input limits |
| Action installs untrusted binary | Pinned release and checksum verification |
| False positive blocks legitimate work | Severity tuning and scoped, expiring exceptions |

## Reporting a vulnerability

Do not put exploit details or live sensitive data in a public issue. Use the repository's private vulnerability-reporting feature once it is enabled.
