default:
	cd cmd/fireback-server && make dev

desktop:
	cd cmd/fireback-desktop && make

mocks:
	f mock write && rm -rf clients/fireback-react/public/md && cp -R ./artifacts/md clients/fireback-react/public/md

github:
	cd clients/fireback-react && npm run github