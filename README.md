# memoguard-cli

Command-line interface for running MemoGuard locally and in CI. The implementation language is Go.

## Owns

- `scan` command, flags, stdin/file input, and configuration discovery.
- Human and JSON output derived from engine findings.
- Stable exit codes and release binaries.

## Does not own

Rule definitions, XDR parsing, or GitHub Action annotations.

## Run

```bash
memoguard scan --xdr transaction.xdr
memoguard scan --xdr - --format json < transaction.xdr
memoguard scan --simulation simulation.json --policy policy.json
```

Use exactly one of `--xdr`, `--memo`, `--simulation`, or `--json`. `--xdr -` reads standard input. The report contains rule names and field paths, never matched values or the original payload. Exit code `0` means no blocking finding, `1` means blocked, `2` means invalid input or policy, and `3` means an operational failure. Warnings alone do not block. For sensitive memos, scan transaction XDR from a file or stdin; a `--memo` shell argument may be retained in history.

This tool scans supported decoded fields; it cannot prove that arbitrary contracts or opaque bytes contain no private data. Run it before submitting a transaction. Go 1.26 and access to tagged private `memoguard-rules` and `memoguard-engine` modules are required to build from source.

Product PRD and architecture live in the parent `memguard/docs` folder in the local workspace.
