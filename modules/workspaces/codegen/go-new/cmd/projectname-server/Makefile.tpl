project = {{ .ctx.Name }}
projectBinary = {{ .ctx.Name }}

dev:
	go build -ldflags "-s -w" -o ../../artifacts/$(project)-server/$(project) . && \
	echo "Binary has been built in: ./artifacts/$(project)-server/$(project) and ./app" && \
	cp ../../artifacts/$(project)-server/$(project) ../../app

everything:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project)_amd64_darwin .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_darwin.zip $(project)_amd64_darwin && cd -
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project)_arm64_darwin .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_darwin.zip $(project)_arm64_darwin && cd -
	GOARCH=arm64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project)_arm64_windows.exe .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_windows.zip $(project)_arm64_windows.exe && cd -
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project)_amd64_windows.exe .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_windows.zip $(project)_amd64_windows.exe && cd -
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project)_arm64_linux .
	cd ../../artifacts/$(project)-server-all && zip $(project)_arm64_linux.zip $(project)_arm64_linux && cd -
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project)_amd64_linux .
	cd ../../artifacts/$(project)-server-all && zip $(project)_amd64_linux.zip $(project)_amd64_linux && cd -
