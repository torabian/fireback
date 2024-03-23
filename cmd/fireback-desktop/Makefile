default:
	cd ../../clients/fireback-react && rm -rf build && npm run build:desktop && \
	cd - && rm -rf ./frontend/dist && cp -R ../../clients/fireback-react/build ./frontend/dist && \
	~/go/bin/wails build;

pkg:
	rm -rf ../../artifacts/academy-desktop-mac && mkdir -p ../../artifacts/academy-desktop-mac && \
	packagesbuild --project ./macos-installer.pkgproj && \
	cp ./build/academy.pkg ../../artifacts/academy-desktop-mac/academy_desktop_mac_latest.pkg;
	rm -rf ./build/academy.pkg;
