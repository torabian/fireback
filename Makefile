default:
	rm -rf app && cd cmd/fireback && make dev

wasm:
	cd cmd/fireback-wasm && make

mac-pkg:
	cd cmd/fireback && make mac-pkg

server:
	cd cmd/fireback && make everything

test:
	FIREBACK_SDK_LOCATION=$(PWD) ./artifacts/fireback/f tests run

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

dockerbuild:
	docker build -t fireback . 

dockerpublish:
	make dockerbuild && docker tag fireback fireback/fireback:latest && docker push fireback/fireback:latest

defs:
	./app emi compile --path modules/fireback/Fireback.emi.yml && \
	./app emi compile --path modules/abac/Abac.emi.yml && \
	./app emi compile --path modules/eventbus/EventBus.emi.yml && \
	./app emi compile --path modules/reactivesearch/ReactiveSearch.emi.yml

interface: interface-manage interface-ss

interface-manage:
	cd modules/project-generator/react-new && \
	npm run manage:build && \
	cd - && \
	rm -rf modules/interfaces/fireback-manage && \
	cp -rf modules/project-generator/react-new/dist modules/interfaces/fireback-manage && \
	git checkout modules/interfaces/fireback-manage/index.go

interface-ss:
	cd modules/project-generator/react-new && \
	npm run self-service:build && \
	cd - && \
	rm -rf modules/interfaces/selfservice && \
	cp -rf modules/project-generator/react-new/dist modules/interfaces/selfservice && \
	git checkout modules/interfaces/selfservice/index.go