# go-whosonfirst-format-wasm

Go package for building Who's On First formatting tools as WebAssembly (WASM and WASI) binaries.

## Important

This is work in progress and not properly documented. Method signatures may still change.

This requires Go 1.21.0 or higher.

## Example

### WASI

```
$> make wasip
GOARCH=wasm GOOS=wasip1 go build -mod vendor -ldflags="-s -w" -o static/wasip/wof-format.wasm cmd/wof-format-wasip/main.go

$> wasmer static/wasip/wof-format.wasm '{"type":"Feature","properties":{"wof:name":"Test","wof:id": 1, "wof:placetype":"custom", "wof:repo":"whosonfirst-data-example"}, "geometry": {"type":"Point", "coordinates": [0.0, 0.0]}}'
{
  "id": 0,
  "type": "Feature",
  "properties": {
    "wof:id": 1,
    "wof:name": "Test",
    "wof:placetype": "custom",
    "wof:repo": "whosonfirst-data-example"
  },
  "bbox": null,
  "geometry": {"coordinates":[0,0],"type":"Point"}
}
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-format