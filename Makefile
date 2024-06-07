default:
	cd cmd/fireback && make dev

mac-pkg:
	cd cmd/fireback && make pkg

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
	