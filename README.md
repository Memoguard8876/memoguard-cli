<p align="center"><img src="assets/logo.svg" alt="MemoGuard logo" width="112"></p>

<h1 align="center">memoguard-cli</h1>

<p align="center"><b>Offline command-line scanner that checks Stellar transactions for private data before you submit.</b></p>

<p align="center">
  <a href="https://github.com/Memoguard8876/memoguard-cli/actions/workflows/ci.yml"><img src="https://github.com/Memoguard8876/memoguard-cli/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/Memoguard8876/memoguard-cli?color=blue" alt="License: MIT"></a>
  <a href="https://github.com/Memoguard8876/memoguard-cli/tags"><img src="https://img.shields.io/github/v/tag/Memoguard8876/memoguard-cli?label=release&color=brightgreen" alt="Latest release"></a>
  <img src="https://img.shields.io/github/go-mod/go-version/Memoguard8876/memoguard-cli?color=00ADD8" alt="Go version">
  <a href="https://github.com/Memoguard8876/memoguard-cli/issues"><img src="https://img.shields.io/github/issues/Memoguard8876/memoguard-cli?color=orange" alt="Open issues"></a>
  <a href="https://github.com/Memoguard8876/memoguard-cli/issues?q=is%3Aopen+label%3A%22help+wanted%22"><img src="https://img.shields.io/badge/help%20wanted-welcome-8A2BE2" alt="Help wanted"></a>
  <img src="https://img.shields.io/badge/built%20for-Stellar-black" alt="Built for Stellar">
</p>

<p align="center">
  <a href="https://cjay-1.gitbook.io/memoguard-docs/">Documentation</a> ·
  <a href="https://github.com/Memoguard8876/memoguard-cli/releases">Releases</a> ·
  <a href="https://github.com/Memoguard8876/memoguard-cli/issues">Issues</a> ·
  <a href="CONTRIBUTING.md">Contributing</a> ·
  <a href="SECURITY.md">Security</a>
</p>

---

## What it is

`memoguard` is the command-line front end. Give it a transaction envelope, a simulation result, decoded JSON or a memo, and it prints a redacted report and sets an exit code your shell or CI can act on. It runs locally. Nothing is sent anywhere.

## Features

- Four inputs: base64 XDR (file or stdin), simulation JSON, decoded JSON, memo text.
- Three outputs: human, JSON and SARIF 2.1.0.
- Stable exit codes, and a `--fail-on` threshold (`block`, `warning`, `none`).
- Custom policies with `--policy`.
- Reports show the rule and field path, **never the matched value or the original payload**.
- Release binaries for Linux, macOS and Windows with `SHA256SUMS`.

## Install

```bash
go install github.com/memoguard8876/memoguard-cli/cmd/memoguard@v0.2.3
```

Or download a binary from the [v0.2.3 release](https://github.com/Memoguard8876/memoguard-cli/releases/tag/v0.2.3) and check it:

```bash
sha256sum --check --ignore-missing SHA256SUMS
```

| Platform | Binary |
| --- | --- |
| Linux x86_64 / arm64 | `memoguard-linux-amd64` / `memoguard-linux-arm64` |
| macOS x86_64 / arm64 | `memoguard-darwin-amd64` / `memoguard-darwin-arm64` |
| Windows x86_64 | `memoguard-windows-amd64.exe` |

## Quick start

```bash
git clone https://github.com/Memoguard8876/memoguard-engine
memoguard scan --xdr memoguard-engine/testdata/unsafe-email.xdr
```

```text
MemoGuard: 1 finding(s)
BLOCK personal.email at transaction.memo.text: Email address in public transaction data. Use an opaque reference instead of an email address
```

The exit code is `1`, and the email address is not printed. A clean file prints `MemoGuard: clean` and exits `0`.

## Usage

```bash
memoguard scan --xdr transaction.xdr
memoguard scan --xdr - --format json < transaction.xdr
memoguard scan --simulation simulation.json --policy policy.json
memoguard scan --json transaction.json
memoguard scan --xdr transaction.xdr --format sarif --fail-on warning
```

Use exactly one of `--xdr`, `--memo`, `--simulation` or `--json`.

| Flag | Default | Meaning |
| --- | --- | --- |
| `--xdr <file>` | | Base64 envelope XDR. `-` reads stdin |
| `--simulation <file>` | | A `simulateTransaction` reply: the full JSON-RPC reply or just its `result`. Unrecognized input exits 2 |
| `--json <file>` | | Decoded transaction JSON |
| `--memo <text>` | | Memo text. It may stay in shell history, so use a file for real data |
| `--policy <file>` | built-in | JSON policy. **Replaces** the built-in rules |
| `--format` | `human` | `human`, `json` or `sarif` |
| `--fail-on` | `block` | Lowest severity that fails: `block`, `warning` or `none` |

## Exit codes

| Code | Meaning |
| --- | --- |
| `0` | Nothing at or above the `--fail-on` threshold |
| `1` | Threshold met |
| `2` | Invalid input, flags or policy, or input beyond a scan limit |
| `3` | Operational failure |

A scan that cannot finish never reports clean.

## JSON report

```json
{"policy_version":"v1","findings":[{"rule_id":"personal.email","severity":"block","confidence":"high","field_path":"transaction.memo.text","description":"Email address in public transaction data","remediation":"Use an opaque reference instead of an email address"}]}
```

SARIF results use `error` for block and `warning` for warning, with the field path as a logical location. Locations are paths inside the decoded transaction, not file lines.

## Build from source

```bash
git clone https://github.com/Memoguard8876/memoguard-cli
cd memoguard-cli
go build ./cmd/memoguard
go test ./...
go vet ./...
```

Needs Go 1.26.3 or newer. The `memoguard-rules` and `memoguard-engine` modules are public.

## Limits

The tool scans supported decoded fields only. It cannot prove that arbitrary contracts or opaque bytes hold no private data. Run it before you submit a transaction. Over-limit input returns exit code `2`, never a clean result. See [What is scanned](https://cjay-1.gitbook.io/memoguard-docs/core-concepts/what-is-scanned) and [Limitations](https://cjay-1.gitbook.io/memoguard-docs/limitations).

Product requirements, architecture, validation record, demo outline and submission notes are versioned in [`product/docs`](product/docs/README.md).

## The MemoGuard family

MemoGuard is four independent Go repositories. Each builds from tagged releases of the one before it.

```text
memoguard-rules ──► memoguard-engine ──► memoguard-cli ──► memoguard-action
```

| Repository | Role | Latest |
| --- | --- | --- |
| [memoguard-rules](https://github.com/Memoguard8876/memoguard-rules) | Policy schema, validation, built-in rules, expiring exceptions | v0.1.1 |
| [memoguard-engine](https://github.com/Memoguard8876/memoguard-engine) | Stellar XDR decoding, field extraction, scanning, redacted findings | v0.2.1 |
| [memoguard-cli](https://github.com/Memoguard8876/memoguard-cli) | `memoguard scan` command, output formats, exit codes, release binaries | v0.2.3 |
| [memoguard-action](https://github.com/Memoguard8876/memoguard-action) | GitHub Action: pinned CLI, annotations, failure threshold | v0.2.1 |

Full guides, the field-path reference and walkthroughs are in **[MemoGuard Docs](https://cjay-1.gitbook.io/memoguard-docs/)**. Product requirements and architecture are versioned in [memoguard-cli/product/docs](https://github.com/Memoguard8876/memoguard-cli/tree/main/product/docs).

## Contributing

Contributions are welcome. Start with [CONTRIBUTING.md](CONTRIBUTING.md).

- Browse [open issues](https://github.com/Memoguard8876/memoguard-cli/issues). Labels show the type (`enhancement`, `documentation`, `testing`), `help wanted`, and a `complexity` level.
- Comment on an issue and wait to be assigned before you start.
- Use **synthetic data only** in tests and fixtures. Never commit real customer data, keys, tokens or real transactions.
- Add tests for every behaviour change, including malformed input, and a line in [CHANGELOG.md](CHANGELOG.md).

## Security

A clean scan is not a guarantee, and there has been no formal security audit. Report vulnerabilities privately through GitHub's private vulnerability reporting for this repository. See [SECURITY.md](SECURITY.md). Do not post exploit details or real private data in a public issue.

## Maintainers

| Maintainer | Role | Contact |
| --- | --- | --- |
| [Memoguard8876](https://github.com/Memoguard8876) | Organization owner, releases | [GitHub issues](https://github.com/Memoguard8876/memoguard-cli/issues) |

## Community

Ask questions and propose changes in [GitHub issues](https://github.com/Memoguard8876/memoguard-cli/issues). Read the [documentation](https://cjay-1.gitbook.io/memoguard-docs/) first; the [FAQ](https://cjay-1.gitbook.io/memoguard-docs/faq) answers the common questions.

## Contributors

<a href="https://github.com/Memoguard8876/memoguard-cli/graphs/contributors"><img src="https://contrib.rocks/image?repo=Memoguard8876/memoguard-cli" alt="Contributors"></a>

## License

[MIT](LICENSE) © MemoGuard contributors.
