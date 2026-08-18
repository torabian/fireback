# By default we only make the binary for testing - use make server for building all the packages
# for release

default:
	rm -rf app && cd cmd/fireback && make dev

# Quick local dev reset: builds the app, drops and recreates the Postgres
# database this checkout's DB_DSN (see .env) points at for a genuinely fresh
# start, applies migrations, seeds the passport methods (via `seeders` - see
# abac.PassportMethodSyncSeeders, wired as SeedersSync in cmd/fireback/main.go)
# so "email" is enabled and the signin screen has something to show, creates a
# root user via `auth --in-root` (the same non-interactive flow
# e2e/cypress/support/setup.js uses), prints `ws view` to confirm it's
# authorized, seeds a few mock users via `user mock`, and sets gin-mode to
# release plus the server port - so once this finishes, `make default &&
# ./app start` serves on APP_PORT and lets you log into the UI right away
# with EMAIL/PASSWORD (override any of these on the command line, e.g.
# `make devsetup EMAIL=me@x.com PASSWORD=secret APP_PORT=8080`).
EMAIL ?= a@a.com
PASSWORD ?= 123321
MOCK_COUNT ?= 5
APP_PORT ?= 4500

demo: default
	@DSN=$$(./app config db-dsn get); \
	VENDOR=$$(./app config db-vendor get); \
	if [ "$$VENDOR" != "postgres" ]; then \
		echo "devsetup only knows how to reset a postgres database right now (DB_VENDOR is '$$VENDOR')"; \
		exit 1; \
	fi; \
	HOST=$$(echo "$$DSN" | grep -oE 'host=[^ ]*' | cut -d= -f2); \
	DBPORT=$$(echo "$$DSN" | grep -oE 'port=[^ ]*' | cut -d= -f2); \
	PGUSER=$$(echo "$$DSN" | grep -oE 'user=[^ ]*' | cut -d= -f2); \
	PGPASS=$$(echo "$$DSN" | grep -oE 'password=[^ ]*' | cut -d= -f2); \
	DBNAME=$$(echo "$$DSN" | grep -oE 'dbname=[^ ]*' | cut -d= -f2); \
	echo "Dropping and recreating database '$$DBNAME' on $$HOST:$$DBPORT ..."; \
	PGPASSWORD=$$PGPASS psql -U $$PGUSER -h $$HOST -p $$DBPORT -d postgres -v ON_ERROR_STOP=1 -c \
		"SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '$$DBNAME' AND pid <> pg_backend_pid();" && \
	PGPASSWORD=$$PGPASS psql -U $$PGUSER -h $$HOST -p $$DBPORT -d postgres -v ON_ERROR_STOP=1 -c "DROP DATABASE IF EXISTS $$DBNAME;" && \
	PGPASSWORD=$$PGPASS psql -U $$PGUSER -h $$HOST -p $$DBPORT -d postgres -v ON_ERROR_STOP=1 -c "CREATE DATABASE $$DBNAME;" && \
	./app migration apply && \
	./app seeders && \
	./app auth --in-root=true --value=$(EMAIL) --type=email --password=$(PASSWORD) \
		--workspace-type-id=root --first-name=Root --last-name=User && \
	./app ws view && \
	./app user mock --count=$(MOCK_COUNT) && \
	./app config gin-mode set release && \
	./app config port set $(APP_PORT) && \
	echo "" && \
	echo "Root user ready - log in with:" && \
	echo "  email:    $(EMAIL)" && \
	echo "  password: $(PASSWORD)" && \
	echo "  serve on: http://localhost:$(APP_PORT) (gin-mode release)"

# Compiles the project into wasm, which would run in browser
wasm:
	cd cmd/fireback-wasm && make

# Builds the wasm-demo app (ui/src/apps/wasm-demo) - a fireback server compiled
# to wasm, running entirely client-side against an in-browser Postgres
# (pglite), no network backend at all. Depends on `wasm` for a fresh
# fireback.wasm (gitignored, ui/public/wasm_exec.js is already committed).
# Output lands in ui/dist, ready to serve as plain static files - see
# .github/workflows/fireback-build.yml's deploy_wasm_demo_pages job, and
# e2e/cypress/e2e/wasm-demo.cy.ts, which drives these same two steps itself.
wasm-demo: wasm
	cd ui && npm run wasm-demo:build

# Compiles everything, zips, and packages for all targets that we are planning.
server:
	cd cmd/fireback && make everything

# Fireback has some sdks on some projects which are commited due to fact I want it
# be ready to use without any builds tools right away. They often get old over changes we make
# to typescript builder for example, and forget to update the codegen projects.
# this function need to do that, and before making any release we need to make
# sure, that running this command on main (or release tag) make any code diff.


# Every Emi action - hand-declared or entity-synthesized alike - is meant to have a
# dedicated <Name>Action_test.go next to its generated wiring file. checkendpointtests
# (see tools/checkendpointtests) reads each module's own preprocessed.yml (the
# "preprocessor" emi target's output - already includes entity CRUD actions, since
# core.ReadEmiFromString runs Preprocess() before writing it) to get the real, complete
# list of actions, and reports which ones have no test file yet.
#
# This is report-only for now (exit 0 regardless of gaps) since there is real, sizable
# test debt today - see the printed count. Once coverage has caught up, switch this to
# `go run ./tools/checkendpointtests --strict` (exits 1 on any gap) and wire it into
# `prepare`/CI so new endpoints can't ship without a test.
checkendpointtests:
	go run ./tools/checkendpointtests

# Runs the Go test suite (GOEXPERIMENT=jsonv2 is required - see modules/backup/Exec.go's
# own comment - and cmd/fireback-wasm has an unrelated, pre-existing build error, so this
# scopes to ./modules/... rather than ./... to stay actually runnable).
#
# modules/abac/tests are black-box tests against a real, already-running `./app start`
# (see modules/abac/tests/testconfig.go) - every one of them skips cleanly, rather than
# failing, when there's no server reachable at ABAC_TEST_BASE_URL (default
# http://localhost:4500). The ones that call an authenticated endpoint also need
# ABAC_TEST_CLI_TOKEN set to a real session token (e.g. from `./app auth` or signing in
# and copying data.item.session.token) - they skip too without one:
#
#   ./app start &
#   TOKEN=$$(curl -s -X POST http://localhost:4500/passports/signin/classic \
#     -H 'Content-Type: application/json' \
#     -d '{"value":"<email>","password":"<password>"}' | \
#     python3 -c 'import json,sys;print(json.load(sys.stdin)["data"]["item"]["session"]["token"])')
#   ABAC_TEST_CLI_TOKEN=$$TOKEN make test
test:
	GOEXPERIMENT=jsonv2 go test ./modules/... -v

# Builds the disk image for docker hub, as fireback can be installed as disk image
dockerbuild:
	docker build -t fireback . 

dockerpublish:
	make dockerbuild && docker tag fireback fireback/fireback:latest && docker push fireback/fireback:latest

# Regenerates ui/packages/js-remote-ctx - the single copy of the framework-agnostic
# runtime (fetch/WebSocket/SSE wrappers, response envelopes, react hooks) that emi's js
# compiler otherwise would have to embed a private copy of into every module's own
# generated output. Every js target below carries the "no-sdk" tag and points
# js-sdk-location at "@fireback/js-remote-ctx" (see each *.emi.yml) specifically so this
# is the one place that content is written - everything else only imports it by name.
# Writes into a throwaway .gen/ subfolder first (js:sdk always nests its output under a
# "sdk/" folder of its own) so we can flatten it up to the package root without ever
# touching js-remote-ctx/package.json.
defs-sdk:
	rm -rf ui/packages/js-remote-ctx/.gen && \
	./app emi js:sdk --output ui/packages/js-remote-ctx/.gen --tags typescript && \
	rm -rf ui/packages/js-remote-ctx/common ui/packages/js-remote-ctx/react ui/packages/js-remote-ctx/js ui/packages/js-remote-ctx/envelopes && \
	cp -r ui/packages/js-remote-ctx/.gen/sdk/. ui/packages/js-remote-ctx/ && \
	rm -rf ui/packages/js-remote-ctx/.gen

# Recompiles the definitions using emi compiler. defs-sdk must run first so every
# module's js target has an up to date @fireback/js-remote-ctx to point at.
defs: defs-sdk
	./app emi compile --path modules/fireback/Fireback.emi.yml && \
	./app emi compile --path modules/abac/Abac.emi.yml && \
	./app emi compile --path modules/abac/messaging/Messaging.emi.yml && \
	./app emi compile --path modules/abac/interfacetools/InterfaceTools.emi.yml && \
	./app emi compile --path modules/eventbus/EventBus.emi.yml && \
	./app emi compile --path modules/reactivesearch/ReactiveSearch.emi.yml && \
	./app emi compile --path modules/backup/Backup.emi.yml && \
	./app emi compile --path modules/internalstats/InternalStats.emi.yml

# Packs every ui/packages/* workspace package (see ui/packages/README.md) into a
# normal, npm-installable tarball under artifacts/fireback-packages/ - so a brand new
# project can `npm install <path-or-url>.tgz` directly (from a local build, or a
# GitHub release asset - see .github/workflows/fireback-build.yml's
# build-ui-packages/deploy_github_release jobs) instead of vendoring this whole repo
# as a git submodule. Raw TS/TSX source in, same raw TS/TSX source out (see
# ui/packages/README.md on why these don't ship a dist/) - the consuming project's own
# bundler still does the compiling, exactly like today.
ui-packages-pack:
	rm -rf artifacts/fireback-packages && mkdir -p artifacts/fireback-packages && \
	cd ui && npm pack --workspaces --pack-destination ../artifacts/fireback-packages

interface: interface-manage interface-ss

interface-manage:
	cd ui && \
	npm run manage:build && \
	cd - && \
	rm -rf modules/interfaces/fireback-manage && \
	cp -rf ui/dist modules/interfaces/fireback-manage && \
	git checkout modules/interfaces/fireback-manage/index.go

interface-ss:
	cd ui && \
	npm run self-service:build && \
	cd - && \
	rm -rf modules/interfaces/selfservice && \
	cp -rf ui/dist modules/interfaces/selfservice && \
	git checkout modules/interfaces/selfservice/index.go