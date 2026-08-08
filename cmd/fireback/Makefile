project = fireback
projectBinary = f

# modules/backup embeds wal-g in-process (see modules/backup/Exec.go), which
# depends on the experimental encoding/json/v2 stdlib package - every build
# target needs this, same as cmd/nima-server/Makefile.
export GOEXPERIMENT = jsonv2

deb:
	bash ./fireback-deb.sh

dev:
	go build --tags "fts5" -ldflags "-s -w" -o ../../artifacts/$(project)/$(projectBinary) . && \
	cp ../../artifacts/$(project)/$(projectBinary) ../../app && \
	cp ../../artifacts/$(project)/$(projectBinary) ../../artifacts/$(project)/$(projectBinary).exe

# Ui2 also copies the fireback-manage, in case it has a different built required
# make sure you change it here too.
wasm:
	GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).wasm .
linux-arm64:
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
linux-amd64:
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .

wininstaller:

	GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).exe . && \
	../../thirdparty/wix314-binaries/candle.exe msi/Product.wxs && \
	../../thirdparty/wix314-binaries/light.exe -out ../../artifacts/$(project)-server-all/fireback_win_arm64_installer.msi ./Product.wixobj 

	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).exe . && \
	../../thirdparty/wix314-binaries/candle.exe msi/Product.wxs && \
	../../thirdparty/wix314-binaries/light.exe -out ../../artifacts/$(project)-server-all/fireback_win_amd64_installer.msi ./Product.wixobj 


mac-pkg:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_darwin.zip $(project) && cd -
	packagesbuild --project ./macos-installer.pkgproj && mv ../../artifacts/$(project)-server-all/$(project).pkg ../../artifacts/$(project)-server-all/$(project)_intel_amd64.pkg
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_darwin.zip $(project) && cd -
	packagesbuild --project ./macos-installer.pkgproj && mv ../../artifacts/$(project)-server-all/$(project).pkg ../../artifacts/$(project)-server-all/$(project)_silicon_arm64.pkg
everything:
	make clean && \
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_darwin.zip $(project) && cd -
	packagesbuild --project ./macos-installer.pkgproj && mv ../../artifacts/$(project)-server-all/$(project).pkg ../../artifacts/$(project)-server-all/$(project)_intel_amd64.pkg
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_darwin.zip $(project) && cd -
	packagesbuild --project ./macos-installer.pkgproj && mv ../../artifacts/$(project)-server-all/$(project).pkg ../../artifacts/$(project)-server-all/$(project)_silicon_arm64.pkg
	GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).exe .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_windows.zip $(project).exe && cd -
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).exe .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_windows.zip $(project).exe && cd -
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_linux.zip $(project) && cd -
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_linux.zip $(project) && cd -
	make deb

windows-zip:
	GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).exe .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_windows.zip $(project).exe && cd -
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).exe .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_windows.zip $(project).exe && cd -

linux-zip:
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_linux.zip $(project) && cd -
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project) .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_linux.zip $(project) && cd -

pkg:
	go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-mac/$(project) . && \
	packagesbuild --project ./macos-installer.pkgproj

clean:
	rm -rf ../../artifacts/$(project)-server-all && mkdir -p ../../artifacts/$(project)-server-all