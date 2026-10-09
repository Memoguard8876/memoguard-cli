<p align="center"><img src="assets/logo.svg" alt="MemoGuard logo" width="112"></p>

# memoguard-cli

[![Go CI](https://github.com/Memoguard8876/memoguard-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/Memoguard8876/memoguard-cli/actions/workflows/ci.yml)

Command-line interface for running MemoGuard locally and in CI. The implementation language is Go.

## Owns

- `scan` command, flags, stdin/file input, and configuration discovery.
- Human and JSON output derived from engine findings.
- Stable exit codes and release binaries.

## Does not own

Rule definitions, XDR parsing, or GitHub Action annotations.

## Run

Install with Go 1.26 or download a binary and `SHA256SUMS` from the [v0.2.2 release](https://github.com/Memoguard8876/memoguard-cli/releases/tag/v0.2.2):

```bash
go install github.com/memoguard8876/memoguard-cli/cmd/memoguard@v0.2.2
```

```bash
memoguard scan --xdr transaction.xdr
memoguard scan --xdr - --format json < transaction.xdr
memoguard scan --simulation simulation.json --policy policy.json
memoguard scan --xdr transaction.xdr --format sarif --fail-on warning
```

Use exactly one of `--xdr`, `--memo`, `--simulation`, or `--json`. `--xdr -` reads standard input. Output formats are `human`, `json`, and SARIF 2.1.0; reports contain rule names and field paths, never matched values or the original payload. Exit code `0` means no finding at or above the chosen failure threshold, `1` means the threshold was met, `2` means invalid input or policy, and `3` means an operational failure. `--fail-on` accepts `block` (default), `warning`, or `none`. For sensitive memos, scan transaction XDR from a file or stdin; a `--memo` shell argument may be retained in history.

This tool scans supported decoded fields; it cannot prove that arbitrary contracts or opaque bytes contain no private data. Inputs beyond supported scan limits return exit code `2` rather than a clean result. Run it before submitting a transaction. Go 1.26 and the tagged `memoguard-rules` and `memoguard-engine` modules are required to build from source.

The shared [product requirements](product/docs/PRD.md), [architecture](product/docs/ARCHITECTURE.md), [validation record](product/docs/VALIDATION.md), and vector brand files are versioned in `product/`. The parent `memguard` folder also keeps a local workspace copy.

The [documentation index](product/docs/README.md), [demo outline](product/docs/DEMO.md), [Wave readiness record](product/docs/WAVE_READINESS.md), and [submission packet](product/docs/SUBMISSION.md) separate validated behavior from remaining pilot and program work.

[Wave application steps and six contributor issues](product/docs/WAVE.md) are documented separately.

See [CONTRIBUTING.md](CONTRIBUTING.md) for changes, [SECURITY.md](SECURITY.md) for private vulnerability reports, and [LICENSE](LICENSE) for MIT terms.

Maintainers: [Memoguard8876](https://github.com/Memoguard8876). Discuss public work in [issues](https://github.com/Memoguard8876/memoguard-cli/issues); report vulnerabilities privately through SECURITY.md.
