package wasip

import (
	"embed"
)

//go:embed *.wasm
var FS embed.FS
