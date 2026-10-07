# Why MemoGuard is a Stellar tool

MemoGuard addresses an irreversible Stellar-specific mistake: a developer can publish private information inside a transaction memo, operation field, or contract interaction. Once the transaction is accepted, that information is public and cannot be removed from the ledger.

The engine understands Stellar transaction-envelope XDR and names the exact public field at risk. A generic text scanner can find an email address in a file, but it does not know whether that text will become part of a Stellar transaction. The GitHub Action and Go API let teams check this at the point where they create or review transactions.

The first product does not need a token, wallet, or smart contract. Its Web3 value comes from protecting users of Stellar applications before a chain write occurs.

## Demonstration

1. Build a testnet transaction with an email address in the memo.
2. Run `memoguard scan` and show a blocking finding for `transaction.memo.text` with no email address printed.
3. Replace the email with an opaque customer reference and show a clean result.
4. Show the same check failing and passing in GitHub Actions.
