default:
	make prepare && \
	~/go/bin/wails build -o projectname-desktop;

clean:
	rm -rf ../../artifacts/projectname-desktop && mkdir -p ../../artifacts/projectname-desktop

prepare:
	echo "Call the html/react/angular project build and copy content to frontend/dist";
	# cd ../../ui && rm -rf build && npm run build:desktop && cd ../cmd/projectname-desktop && \
	# rm -rf ./frontend/dist && mkdir -p ./cmd/projectname/frontend && cp -R ../../ui/build ./frontend/dist && \
	# cp ./assets/appicon.png ./build/appicon.png
pkg:
	rm -rf ../../artifacts/projectname-desktop && mkdir -p ../../artifacts/projectname-desktop && \
	packagesbuild --project ./macos-installer.pkgproj && \
	cp ./build/academy.pkg ../../artifacts/projectname-mac/academy_desktop_amd_mac_latest.pkg;
	rm -rf ./build/academy.pkg;

darwin:
	make clean && \
	make prepare && \
	~/go/bin/wails build -platform darwin/amd64 && \
	cd build/bin && zip -9 -y -r -q projectname_desktop_stanalone_intel_amd64.zip "Raspberry Pie Studio.app" && cd - && \
	mv build/bin/projectname_desktop_stanalone_intel_amd64.zip ../../artifacts/projectname-desktop/
	packagesbuild --project ./macos-installer.pkgproj && \
	mv "build/Raspberry Pie Studio Desktop.pkg" ../../artifacts/projectname-desktop/projectname_desktop_mac_amd64_intel_installer.pkg
	~/go/bin/wails build -platform darwin/arm64 && \
	cd build/bin && zip -9 -y -r -q projectname_desktop_stanalone_mac_silicon_arm64.zip "Raspberry Pie Studio.app" && cd - && \
	mv build/bin/projectname_desktop_stanalone_mac_silicon_arm64.zip ../../artifacts/projectname-desktop/
	packagesbuild --project ./macos-installer.pkgproj && \
	mv "build/Raspberry Pie Studio Desktop.pkg" ../../artifacts/projectname-desktop/projectname_desktop_mac_silicon_arm64_installer.pkg

windows:
	make clean && \
	make prepare && \
	mkdir -p ../../artifacts/projectname-desktop && \
	~/go/bin/wails build -platform windows/arm64 -o projectname_windows_arm64.exe
	mv "build/bin/projectname_windows_arm64.exe" ../../artifacts/projectname-desktop/projectname.exe && \
	cd ../../artifacts/projectname-desktop && zip projectname_windows_arm64.zip projectname.exe && cd -

	~/go/bin/wails build -platform windows/amd64 -o projectname_windows_amd64.exe
	mv "build/bin/projectname_windows_amd64.exe" ../../artifacts/projectname-desktop/projectname.exe
	cd ../../artifacts/projectname-desktop && zip projectname_windows_amd64.zip projectname.exe

linux:
	# make clean && \
	# make prepare && \
	~/go/bin/wails build -platform linux/arm64 -o projectname_linux_arm64
	mv "build/bin/projectname_linux_arm64" ../../artifacts/projectname-desktop/projectname_linux_arm64
	~/go/bin/wails build -platform linux/amd64 -o projectname_linux_amd64
	mv "build/bin/projectname_linux_amd64" ../../artifacts/projectname-desktop/projectname_linux_amd64

all:
	make linux darwin windows