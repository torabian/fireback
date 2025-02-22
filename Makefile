default:
	cd cmd/fireback && make dev

mac-pkg:
	cd cmd/fireback && make mac-pkg

server:
	cd cmd/fireback && make everything

desktop:
	cd cmd/fireback-desktop && make

npm:
	cd cmd/fireback && make npm

npmp:
	cd cmd/fireback && make npmp

test:
	FIREBACK_SDK_LOCATION=$(PWD) ./artifacts/fireback/f tests run

bed:
	rm -rf ../fbtest && cd .. && ./fireback/artifacts/fireback/f new --name fbtest --ui --mobile --replace-fb ../fireback --module github.com/torabian/fireback/testbed

test_rebuild:
	node e2e/scripts/rebuild.js $(PWD)

refresh:
	./artifacts/fireback/f gen gof --def modules/workspaces/WorkspaceModule3.yml --relative-to . --gof-module github.com/torabian/fireback --no-cache true && \
	make


# Fireback has some sdks on some projects which are commited due to fact I want it
# be ready to use without any builds tools right away. They often get old over changes we make 
# to typescript builder for example, and forget to update the codegen projects.
# this function need to do that, and before making any release we need to make
# sure, that running this command on main (or release tag) make any code diff.

rebuild-sdks:
	rm -rf gen react --path e2e/react-bed/src/sdk && \
	rm -rf modules/workspaces/codegen/react-new/src/modules/fireback/sdk && \
	rm -rf modules/workspaces/codegen/react-native-new/src/modules/fireback/sdk && \
	./app gen react --path e2e/react-bed/src/sdk --no-cache true && \
	./app gen react --path modules/workspaces/codegen/react-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-new && npm run build
	./app gen react --path modules/workspaces/codegen/react-native-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-native-new 

## This is different because we use the fireback built on ci-cd for this purpose.
rebuild-sdks-ci:
	rm -rf gen react --path e2e/react-bed/src/sdk && \
	rm -rf modules/workspaces/codegen/react-new/src/modules/fireback/sdk && \
	rm -rf modules/workspaces/codegen/react-native-new/src/modules/fireback/sdk && \
	fireback gen react --path e2e/react-bed/src/sdk --no-cache true && \
	fireback gen react --path modules/workspaces/codegen/react-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-new && npm run build
	fireback gen react --path modules/workspaces/codegen/react-native-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-native-new 