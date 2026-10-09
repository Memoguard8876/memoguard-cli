# MemoGuard Wave readiness

Checked 9 October 2026. This is an evidence record, not a claim of Drips approval. Recheck live services before submitting.

## Fit and limits

MemoGuard scans supported, decoded Stellar transaction fields before submission. A Go rules library defines policy; a Go engine decodes Stellar XDR and extracts fields; a CLI presents redacted findings; a GitHub Action runs the pinned CLI in CI. Stellar XDR and transaction semantics are load-bearing: a generic secret scanner does not know which bytes will be published as a memo, operation field, contract argument, or simulation result. The [Stellar transaction docs](https://developers.stellar.org/docs/learn/fundamentals/transactions/operations-and-transactions) describe memos and operations, and the [Stellar simulation docs](https://developers.stellar.org/docs/data/apis/rpc/api-reference/methods/simulateTransaction) describe pre-submit simulation. MemoGuard has no Soroban contract, wallet, web app, stateful backend, or database. A clean report is not proof that opaque or unsupported content is safe.

The [live Stellar Wave repo page](https://www.drips.network/wave/stellar/repos) reported 824 approved repos when checked. The visible examples span payments, escrow, DeFi, marketplaces, and developer tooling. That sample cannot prove MemoGuard is unique across all 824. A focused web search found Stellar XDR parsers and transaction-inspection tools but did not establish a directly equivalent pre-submit privacy scanner; this is a research lead, not proof of white space. The [SDF grants page](https://stellar.org/grants-and-funding) supports projects growing Stellar and Soroban; [SCF 7.0](https://stellar.org/blog/ecosystem/introducing-scf-v7) has an Open Track and a separate RFP track for requested tooling. SCF and Drips Wave are different programs.

## Phase check against the builder playbook

| Phase | Status for MemoGuard | Evidence or remaining work |
| --- | --- | --- |
| 1. Ecosystem reconnaissance | Partial | Live Wave and SDF sources checked. No exhaustive categorization or uniqueness proof across approved repos. |
| 2–3. Ideas and critical review | Selected idea documented; full historical exercise not recorded | [PRD](PRD.md), [decisions](DECISIONS.md), and [why Stellar](WHY_STELLAR.md) explain the problem and Stellar fit. The main weak spot is incomplete field coverage and false-positive/false-negative behavior; a consenting pilot must measure both. |
| 4. Name and repo structure | Complete | The user selected MemoGuard. Four repositories have separate boundaries and release order in [REPOSITORIES.md](REPOSITORIES.md). |
| 5–7. Contract and app system prompts | Not applicable | No contract, frontend, or indexer is needed for this Go scanning tool. Go package and CLI interfaces are in [ARCHITECTURE.md](ARCHITECTURE.md) and [INTERFACES.md](INTERFACES.md). |
| 8. Local build and deployment | Complete for distributable tool | [VALIDATION.md](VALIDATION.md) records tests, fuzz smoke runs, public CLI binaries, and Action integration on three operating systems. No on-chain deployment exists or is required. |
| 9. Hosting topology | Complete for current design | Users run the CLI locally or the Action in GitHub CI; Go apps can import the engine. There is no always-on service or database. |
| 10. Repo hygiene | Check live before submission | MIT, CONTRIBUTING, SECURITY, logos, CI, releases/tags, and six scoped issues exist. Confirm current topics, branch rules, checks, and issue bodies on GitHub. The owner asked that Wave labels remain untouched. |
| 11. Documentation site | Versioned Markdown docs complete; hosted docs site absent | Product docs live in `memoguard-cli/product/docs`; a separate GitBook or Pages site has not been published. |
| 12. Submission | Not submitted | See [SUBMISSION.md](SUBMISSION.md). A screen-recorded demo and a real-user pilot do not yet exist. CLI release and Action URLs are the live product artifacts; a web-app URL and contract ID do not apply. |
| 13. Post-approval iteration | Not applicable yet | Requires Wave acceptance and contributor applications. |

## Pilot boundary

Before recommending MemoGuard as a mandatory CI gate, run it with a consenting Stellar team against sanitized examples from their own transaction flow. Measure missed fields and false alarms, tune jurisdiction-specific rules, and document unsupported data. No formal security audit or comprehensive customer pilot has been completed.
