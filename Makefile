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

vscode:
	cd clients/fireback-tools-vs-code-extension && npm run package
	

test:
	FIREBACK_SDK_LOCATION=$(PWD) ./artifacts/fireback/f tests run

bed:
	rm -rf ../fbtest && cd .. && ./fireback/artifacts/fireback/f new --name fbtest --ui --mobile --replace-fb ../fireback --module github.com/torabian/fireback/testbed

test_rebuild:
	node e2e/scripts/rebuild.js $(PWD)

refresh:
	./artifacts/fireback/f gen gof --def modules/workspaces/WorkspaceModule3.yml --relative-to . --gof-module github.com/torabian/fireback --no-cache true && \
	./artifacts/fireback/f gen gof --def modules/geo/GeoModule3.yml --relative-to . --gof-module github.com/torabian/fireback --no-cache true && \
	./artifacts/fireback/f gen gof --def modules/licenses/LicenseModule3.yml --relative-to . --gof-module github.com/torabian/fireback --no-cache true && \
	./artifacts/fireback/f gen gof --def modules/worldtimezone/WorldTimeZoneModule3.yml --relative-to . --gof-module github.com/torabian/fireback --no-cache true && \
	./artifacts/fireback/f gen gof --def modules/currency/CurrencyModule3.yml --relative-to . --gof-module github.com/torabian/fireback --no-cache true && \
	make


# Fireback has some sdks on some projects which are commited due to fact I want it
# be ready to use without any builds tools right away. They often get old over changes we make 
# to typescript builder for example, and forget to update the codegen projects.
# this function need to do that, and before making any release we need to make
# sure, that running this command on main (or release tag) make any code diff.

rebuild-sdks:
	./app gen react --path modules/workspaces/codegen/react-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-new && npm run build
	./app gen react --path modules/workspaces/codegen/react-native-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-native-new 

## This is different because we use the fireback built on ci-cd for this purpose.
rebuild-sdks-ci:
	./app gen react --path modules/workspaces/codegen/react-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-new && npm run build
	./app gen react --path modules/workspaces/codegen/react-native-new/src/modules/fireback/sdk --no-cache true && \
	cd modules/workspaces/codegen/react-native-new 
