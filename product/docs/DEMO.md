# MemoGuard demo outline

Use only the checked-in synthetic fixtures. Do not record real transaction data or secret keys.

1. Download and verify the [CLI v0.2.2 release](https://github.com/Memoguard8876/memoguard-cli/releases/tag/v0.2.2), or build from the public source with Go 1.26.3 or newer.
2. From a `memoguard-cli` checkout beside `memoguard-engine`, run `go run ./cmd/memoguard scan --xdr ../memoguard-engine/testdata/unsafe-email.xdr`. The expected output is one `personal.email` block at `transaction.memo.text`, with no matched email printed; the command exits 1.
3. Run `go run ./cmd/memoguard scan --xdr ../memoguard-engine/testdata/clean.xdr`. The expected output is `MemoGuard: clean` with exit code 0. Explain that clean means only the supported fields were inspected.
4. Show the `memoguard-action` integration workflow checking clean, warning, and blocked synthetic fixtures on Linux, macOS, and Windows. Explain that the Action runs the pinned CLI rather than a second scanner.
5. End with the limitation: opaque bytes and unsupported shapes need coverage work, and a consenting pilot must tune the policy before it becomes a required gate.

The [validation record](VALIDATION.md) links the CLI release workflow. A video based on these steps has not yet been recorded.
