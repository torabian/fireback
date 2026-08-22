# wallet / walletpublic

Two-module wallet system: `wallet` is the internal engine (entities, gateway
abstraction, root/admin actions, and the in-process Go purchase API other nima modules
call); `walletpublic` is the owner-facing API (create/list/topup/history for a wallet's
own user or workspace), returning public-projection DTOs so internal fields never leak.

## Entities (`wallet`)

- **wallet** — owned by exactly one user _or_ one workspace (`ownerType` +
  `userId`/`workspaceId`), one currency, `balance` as a minor-units decimal string.
  No public create/update action — see "Money" and "Restricted actions" below.
- **walletCurrency** — supported currencies (fiat or crypto), each with a `decimals`
  precision (2 for USD, 8 for BTC, 18 for ETH, ...).
- **walletTransaction** — append-only balance-history ledger. One row per balance change,
  written in the same DB transaction as the balance update itself.
- **walletGateway** — a registered payment provider (code + config). Behavior lives in a
  `GatewayAdapter` Go implementation registered under that code (see below), not in the
  emi schema.
- **walletPaymentAttempt** — one attempt to move money through a gateway (topup/purchase/
  withdrawal), with status, gateway reference, and links back to the resulting ledger
  entry once it succeeds.
- **walletEvent** — raw inbound gateway webhook/event log, for audit and replay.
- **walletConfig** — root-only wallet limits (max wallets per user/workspace, optionally
  per currency). No generic entity actions at all (`features.actions: false`) — only
  reachable through `getWalletConfig`/`updateWalletConfig`.
- **walletProviderConfig** — root-only, DB-backed config for one provider adapter
  (`providerType`, matching a `GatewayAdapter.Code()`) in one `region` (an ISO-3166-ish
  code, or `"global"` for every region), plus `isEnabled` and a non-secret `config` JSON
  blob. A composite unique index on `(providerType, region)` enforces at most one row per
  pair. Unlike `walletConfig`, this one keeps the full generated CRUD (create/get/browse/
  update/delete) — enabling/disabling a provider or editing its config is just an update
  like any other field, gated root-only (see `WalletProviderConfigImplementation.go`).
  Not yet consulted by `topup`/the webhook handler for provider selection - see "Not yet
  done".

## Money: minor-units strings, never float

Every amount field (`wallet.balance`, `walletTransaction.amount`,
`walletPaymentAttempt.amount`) is a **string holding an integer count of minor units** at
the currency's declared `decimals` precision — e.g. USD `"10050"` = $100.50, BTC
`"150000000"` = 1.5 BTC. `float32`/`float64` are never used for money: float64 only
carries ~15-17 significant decimal digits, fine for USD-scale amounts but silently lossy
at 18-decimal crypto scale, and a ledger must never round something nobody asked for.
Server-side arithmetic goes exclusively through `math/big.Int` via `Money.go`.

## Concurrency & idempotency: `Purchase.go`

Every balance mutation — `purchase`, `adjustBalance`, and a topup's `gatewayWebhook`
credit — goes through one function, `applyLedgerEntry` (`Purchase.go`):

1. If a `walletTransaction` with the given `idempotencyKey` already exists, return it
   unchanged instead of re-applying the change (safe retry).
2. Inside one DB transaction, take a row lock on the wallet
   (`SELECT ... FOR UPDATE`), so concurrent calls against the _same_ wallet serialize.
3. Validate status/amount via `math/big`, compute the new balance.
4. Update the wallet's balance and insert the `walletTransaction` row together, in the
   same transaction.
5. If two concurrent callers both race past step 1 with the same `idempotencyKey`, the
   unique index on `walletTransaction.idempotencyKey` rejects the loser's insert; that's
   caught and turned back into "return the existing row" too.

`Purchase_test.go` proves this with a real (sqlite, single-connection) DB: sequential
debits + insufficient-balance rejection, idempotent-retry-never-double-debits, and a
concurrency test where 25 goroutines race to debit a wallet that can only afford 10 —
exactly 10 must succeed, the final balance must land at exactly 0, and the ledger must
have exactly 10 rows. Run with `go test ./modules/wallet/... -race`.

**Note on the test DB:** every generated entity's `uniqueId` column declares
`default:gen_random_uuid()` (Postgres/pgcrypto) — not valid SQLite syntax, so
`AutoMigrate` can't be used against sqlite as-is. `Purchase_test.go` hand-creates the two
tables it needs with SQLite-safe DDL instead. This repo's entity modules are otherwise
only ever exercised against real Postgres (see `docker-compose.yml`) — that remains the
path for a full authenticated CLI/HTTP walkthrough (`docker-compose up -d postgres`, then
`authorize --in-root ...` to get a root session).

## Internal Go API for other modules

```go
import wallet "github.com/torabian/fireback/modules/finance/wallet"

entry, err := wallet.Purchase(wallet.PurchaseInput{
    WalletUniqueId: walletId,
    Amount:         "1999",              // minor units
    ReferenceType:  "course-purchase",   // your module's own identity
    ReferenceId:    orderId,
    IdempotencyKey: orderId,             // make retries safe
})
```

No HTTP round-trip needed. This is the same function the `purchase` HTTP action
(gated behind `PERM_ROOT_WALLET_PURCHASE`, for trusted service callers) calls.

## Gateway abstraction

`GatewayAdapter.go` defines the interface (`InitiatePayment`, `VerifyWebhook`) every
concrete provider implements. `MockGatewayAdapter` (auto-succeeds) and
`ManualGatewayAdapter` (bank-transfer-style, admin marks it succeeded by hand) ship by
default for CLI/dev use — pass real ones via `WalletModuleConfig.Gateways`. The webhook
receiver (`GatewayWebhookHandler`) is a hand-wired plain Gin route
(`Any /wallet/gateway/:code/webhook`, registered for every HTTP method — see below), not
an emi action — a webhook needs the raw request body/headers for the gateway's own
signature verification and has no caller session, neither of which fits emi's
typed-action model.

### Real providers: `modules/wallet/providers/{stripe,przelewy24,blik,zarinpal}`

Four self-contained packages, each implementing `wallet.GatewayAdapter` against a real
provider's actual API, injected in `cmd/nima-server/main.go`:

| Package      | Provider                         | Env vars                                                                     | Callback shape                                    |
| ------------ | -------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------- |
| `stripe`     | Stripe Payment Intents           | `STRIPE_SECRET_KEY`, `STRIPE_WEBHOOK_SECRET`                                 | Signed POST (`Stripe-Signature`, HMAC-SHA256)     |
| `przelewy24` | Przelewy24 (P24) REST API v1     | `P24_MERCHANT_ID`, `P24_POS_ID`, `P24_API_KEY`, `P24_CRC_KEY`, `P24_SANDBOX` | Signed POST (SHA384) + a synchronous verify call  |
| `blik`       | BLIK, via a **generic acquirer** | `BLIK_ACQUIRER_BASE_URL`, `BLIK_API_KEY`, `BLIK_MERCHANT_ID`                 | Signed POST (HMAC-SHA256)                         |
| `zarinpal`   | ZarinPal REST API v4 (Iran)      | `ZARINPAL_MERCHANT_ID`, `ZARINPAL_SANDBOX`                                   | Unsigned GET redirect + a synchronous verify call |

All four also read `WALLET_PUBLIC_BASE_URL` (this server's own public URL) to build the
callback URL they hand to the provider. `New(Config{})` never fails on missing
env/config — an unconfigured adapter just returns a clear error from
`InitiatePayment`/`VerifyWebhook`, so wiring all four unconditionally in `main.go` is
safe; only the ones actually configured work. A matching `walletGateway` row (`wallet
gateway create --code stripe ...`) is still required per provider — not auto-seeded,
same as currencies.

**BLIK has no vendor-neutral direct-to-merchant API** — every real integration goes
through some acquirer/PSP (Stripe and Przelewy24 both support BLIK as one of _their own_
payment methods too). The `blik` package is a generic acquirer adapter you point at
whichever contract the business actually holds via `BLIK_ACQUIRER_BASE_URL`; see its
package doc comment.

**ZarinPal has no signed server-to-server webhook** — it redirects the payer's browser
back via GET (`?Authority=...&Status=OK|NOK`) instead, and authenticity comes from a
synchronous `verify.json` call rather than a shared-secret signature. Rather than adding
a second interface method for this, `GatewayWebhookHandler` normalizes _any_ HTTP method
into the same `(rawBody []byte, headers http.Header)` shape every `VerifyWebhook`
expects (for GET, the query string becomes a flat JSON object) — the interface itself
stayed exactly one method-set for all four providers. See
`GatewayWebhookImplementation.go`'s doc comment and `providers/zarinpal`'s.

Every provider's request-building/response-parsing is unit-tested against an
`httptest.Server` (no live credentials), plus deterministic signature-algorithm tests
(`stripe`: HMAC-SHA256 test vectors; `przelewy24`: SHA384). `P24`'s notification-sign
field order is flagged in its own package doc comment as worth a sandbox check before
production use — P24's public docs don't fully spell it out.

## Restricted actions

`wallet`, `walletTransaction`, `walletPaymentAttempt`, `walletEvent` don't expose the
generic create/update/delete actions a normal emi entity gets (see each entity's
`features` override in `Wallet.emi.yml`) — balance/ledger/attempt state only changes
through the paths above, never a generic CRUD endpoint a client could point balance-writes
at directly.

## CLI verification

```
wallet currency create ...          # register a currency
wallet gateway create ...           # register a gateway (code must match a GatewayAdapter)
wallet-public create ...            # owner creates a wallet
wallet-public topup ...             # start a topup (mock/manual adapters auto-succeed)
wallet-public my                    # list my wallets
wallet purchase ...                 # (service-permission) debit a wallet
wallet-public history ...           # view the resulting ledger
```

Every action above has a `{name}Gin`/`{name}CliHandler` pair generated from the exact
same `*Implementation.go` function, so a CLI-verified action's HTTP surface is exercising
identical logic — see `RouterManifest.go`/`*Module.go` in each package.

## Not yet done

- Frontend (`ui/src/modules/wallet`, `ui/src/modules/walletpublic`) and translations
  (`strings-{en,fa,ru,pl}.yml`) — planned but not yet implemented.
- The wallet owner's email isn't threaded through to `topup` yet, which Przelewy24's
  register call wants (see the `TODO` in `providers/przelewy24/Przelewy24.go`).
- On-chain crypto gateway adapters.
- `walletProviderConfig` rows aren't consulted anywhere yet - `topup` still takes an
  explicit `gatewayCode` and looks up `walletGateway` directly. Wiring region-aware
  provider selection (e.g. "pick the enabled provider for the caller's region, falling
  back to a `region: global` row") through `topup`/`GatewayWebhookHandler` is a natural
  next step, kept out of this change since it wasn't asked for yet.
