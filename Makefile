# By default we only make the binary for testing - use make server for building all the packages
# for release

default:
	rm -rf app && cd cmd/fireback && make dev

# Compiles the project into wasm, which would run in browser
wasm:
	cd cmd/fireback-wasm && make

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

# Recompiles the definitions using emi compiler.
defs:
	./app emi compile --path modules/fireback/Fireback.emi.yml && \
	./app emi compile --path modules/abac/Abac.emi.yml && \
	./app emi compile --path modules/abac/messaging/Messaging.emi.yml && \
	./app emi compile --path modules/abac/interfacetools/InterfaceTools.emi.yml && \
	./app emi compile --path modules/eventbus/EventBus.emi.yml && \
	./app emi compile --path modules/reactivesearch/ReactiveSearch.emi.yml

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