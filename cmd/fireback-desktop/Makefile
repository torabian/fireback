new: 
	~/go/bin/wails build -o fireback;
	
default:
	cd ../../ui && rm -rf build && npm run ac:build:desktop && cd ../cmd/academy-desktop && \
	cp -R ../academy-core/embed-content/* ../../ui/build && \
	rm -rf ./frontend/dist && mkdir -p ./cmd/academy-desktop/frontend && cp -R ../../ui/build ./frontend/dist && \
	cp ./assets/appicon.png ./build/appicon.png && \
	~/go/bin/wails build -o academydesktop;

pkg:
	rm -rf ../../artifacts/academy-desktop-mac && mkdir -p ../../artifacts/academy-desktop-mac && \
	packagesbuild --project ./macos-installer.pkgproj && \
	cp ./build/academy.pkg ../../artifacts/academy-desktop-mac/academy_desktop_mac_latest.pkg;
	rm -rf ./build/academy.pkg;
