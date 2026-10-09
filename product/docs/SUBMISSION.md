# MemoGuard Stellar Wave submission packet

This is draft material for the organization owner. Do not claim approval or a real-user pilot until either is verified.

## Project description

MemoGuard is a Go tool that checks Stellar transaction data for accidental disclosure before a transaction is submitted. It decodes supported Stellar XDR and Soroban-related fields, applies a configurable privacy policy, and reports the rule and field path without repeating the sensitive value. Teams can run it locally through a CLI or in GitHub Actions. It helps catch mistakes while they can still be fixed; it cannot remove data already published on Stellar or guarantee that opaque content is safe.

## Why Stellar

Transaction memos, operation fields, and contract interactions have Stellar-specific encodings and publication paths. MemoGuard uses the official Stellar Go SDK to decode those shapes and locate the public field at risk. This is an off-chain, pre-submit check; it requires no token or Soroban contract.

## Repository relationship

- [memoguard-rules](https://github.com/Memoguard8876/memoguard-rules) owns validated policies, built-in rules, and scoped exceptions.
- [memoguard-engine](https://github.com/Memoguard8876/memoguard-engine) imports the rules, decodes Stellar data, and emits redacted findings.
- [memoguard-cli](https://github.com/Memoguard8876/memoguard-cli) imports the engine and rules, handles files and flags, and publishes binaries. It also versions these product docs.
- [memoguard-action](https://github.com/Memoguard8876/memoguard-action) downloads a checksum-pinned CLI release and turns its safe report into CI annotations. It does not duplicate detection logic.

## Planned contributor work

Six scoped issues are open: policy [overlays](https://github.com/Memoguard8876/memoguard-rules/issues/1) and [schema fixtures](https://github.com/Memoguard8876/memoguard-rules/issues/2); engine [Soroban XDR fixtures](https://github.com/Memoguard8876/memoguard-engine/issues/1) and [opaque-field coverage notes](https://github.com/Memoguard8876/memoguard-engine/issues/2); CLI [multi-file scanning](https://github.com/Memoguard8876/memoguard-cli/issues/1); and Action [optional SARIF upload](https://github.com/Memoguard8876/memoguard-action/issues/1). The owner asked that no `Stellar Wave` label be applied by this preparation work. The program workflow can add approved issues later.

## Evidence and missing links

| Item | Current evidence | Status |
| --- | --- | --- |
| Source and docs | Four public repositories, MIT licenses, READMEs, CI, [product docs](README.md), and releases for [rules](https://github.com/Memoguard8876/memoguard-rules/releases/tag/v0.1.1), [engine](https://github.com/Memoguard8876/memoguard-engine/releases/tag/v0.2.1), [CLI](https://github.com/Memoguard8876/memoguard-cli/releases/tag/v0.2.2), and [Action](https://github.com/Memoguard8876/memoguard-action/releases/tag/v0.2.1) | Available |
| Downloadable CLI | [v0.2.2 release](https://github.com/Memoguard8876/memoguard-cli/releases/tag/v0.2.2) with binaries and checksums | Available |
| Reusable Action | [memoguard-action](https://github.com/Memoguard8876/memoguard-action), pinned to CLI v0.2.2 | Available |
| Demo video | No screen recording of unsafe fixture, redacted finding, corrected fixture, and Action check | Missing |
| Real-user pilot | No consenting team has tuned rules against its sanitized transaction shapes | Missing |
| Live app URL or contract explorer | No web app or contract in this product | Not applicable |
| Drips application and approval | Not recorded | Owner action and organizer decision |

## Owner submission sequence

1. Search the [live approved-repo list](https://www.drips.network/wave/stellar/repos) for MemoGuard immediately before applying. The list can change.
2. Record a short demo using synthetic fixtures: show a blocked finding without a leaked value, a clean corrected fixture, and the Action check. See [DEMO.md](DEMO.md).
3. Run a consenting pilot if possible and record false alarms and uncovered fields. Do not claim production-grade completeness before that test.
4. Sign in to [Drips Wave](https://www.drips.network/wave/stellar), install its GitHub App on `Memoguard8876`, sync the four repos, and apply each relevant repo to the Stellar Wave Program.
5. Wait for organizer approval. Then select issues through the program workflow; leave Wave labels to Stellar as requested.
