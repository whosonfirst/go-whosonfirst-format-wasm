GOROOT=$(shell go env GOROOT)
GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

wasmjs:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" static/javascript/

wasip:
	GOARCH=wasm GOOS=wasip1 \
		go build -mod $(GOMOD) -ldflags="-s -w" \
		-o static/wasip/wof-format.wasm
		cmd/wof-format-wasip/main.go

wasm:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="-s -w" \
		-o static/wasm/wof-format.wasm \
		cmd/wof-format-wasm/main.go
