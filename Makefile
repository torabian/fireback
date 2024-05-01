default:
	cd cmd/fireback-server && make dev

mac-pkg:
	cd cmd/fireback-server && make pkg

desktop:
	cd cmd/fireback-desktop && make

mocks:
	f mock write && rm -rf clients/fireback-react/public/md && cp -R ./artifacts/md clients/fireback-react/public/md

github:
	cd clients/fireback-react && npm run github

capacitor:
	cd clients/fireback-react && npm run capacitor && cd - && rm -rf clients/fireback-capacitor/www && cp -R clients/fireback-react/build clients/fireback-capacitor/www
