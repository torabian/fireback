project = fireback
projectBinary = f

wasm:
	GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o ../../artifacts/$(project)-server-all/$(project).wasm . && \
	cp ../../artifacts/$(project)-server-all/$(project).wasm ../../ui/public/$(project).wasm
