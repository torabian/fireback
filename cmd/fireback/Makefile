project = fireback
projectBinary = f

deb:
	bash ./fireback-deb.sh
dev:
	go build -ldflags "-s -w" -o ../../artifacts/$(project)/$(projectBinary) . && \
	cp ../../artifacts/$(project)/$(projectBinary) ../../app && \
	../../app gen module3spec --out ../../modules/workspaces/module3-json-schema.json

# Ui2 also copies the fireback-manage, in case it has a different built required
# make sure you change it here too.

ui2:
	cd ../../modules/workspaces/codegen/react-new && npm run build && cd - && \
<<<<<<< HEAD
	rm -rf ui && cp -R ../../modules/workspaces/codegen/react-new/build ui && \
	rm -rf ../../modules/workspaces/codegen/fireback-manage && \
	cp -R ../../modules/workspaces/codegen/react-new/build ../../modules/workspaces/codegen/fireback-manage && \
	git checkout ../../modules/workspaces/codegen/fireback-manage/index.go
=======
	rm -rf ui && cp -R ../../modules/workspaces/codegen/react-new/build ui
	cd ../../modules/workspaces/codegen/react-new && npm run self-service:build && cd - && \
	rm -rf ../../modules/workspaces/codegen/selfservice && cp -R ../../modules/workspaces/codegen/react-new/build ../../modules/workspaces/codegen/selfservice && \
	git checkout ../../modules/workspaces/codegen/selfservice/index.go 

>>>>>>> @{-1}

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
	make ui2 && \
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

npm:
	make npmc npmp

npmc:
	cp ../../artifacts/$(project)-server-all/$(project)_amd64_darwin ../fireback-npm/cli-darwin-x64/bin/fireback && \
	cp ../../artifacts/$(project)-server-all/$(project)_arm64_darwin ../fireback-npm/cli-darwin-arm64/bin/fireback && \
	cp ../../artifacts/$(project)-server-all/$(project)_amd64_linux ../fireback-npm/cli-linux-x64/bin/fireback && \
	cp ../../artifacts/$(project)-server-all/$(project)_amd64_windows.exe ../fireback-npm/cli-win32-x64/bin/fireback.exe

npmp:
	cd ../fireback-npm/cli && npm publish && cd - && \
	cd ../fireback-npm/cli-darwin-x64 && npm publish && cd - && \
	cd ../fireback-npm/cli-darwin-arm64 && npm publish && cd - && \
	cd ../fireback-npm/cli-linux-x64 && npm publish && cd - && \
	cd ../fireback-npm/cli-win32-x64 && npm publish

clean:
	rm -rf ../../artifacts/$(project)-server-all && mkdir -p ../../artifacts/$(project)-server-all