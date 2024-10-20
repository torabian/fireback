default:
	cd cmd/fireback && make dev

mac-pkg:
	cd cmd/fireback && make mac-pkg

server:
	cd cmd/fireback && make everything

desktop:
	cd cmd/fireback-desktop && make

mocks:
	f mock write && rm -rf clients/fireback-react/public/md && cp -R ./artifacts/md clients/fireback-react/public/md

github:
	cd clients/fireback-react && npm run github

capacitor:
	cd clients/fireback-react && npm run capacitor && cd - && rm -rf clients/fireback-capacitor/www && cp -R clients/fireback-react/build clients/fireback-capacitor/www

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