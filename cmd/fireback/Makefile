project = fireback
projectBinary = f

dev:
	go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-mac/$(project) . && \
	mv ../../artifacts/$(project)-server-mac/$(project) /usr/local/bin/$(project)

everything:
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
