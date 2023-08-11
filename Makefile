GOROOT=$(shell go env GOROOT)
GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

wasip:
	GOARCH=wasm GOOS=wasip1 go build -mod $(GOMOD) -ldflags="-s -w" -o static/wasip/wof-format.wasm cmd/wof-format-wasip/main.go
