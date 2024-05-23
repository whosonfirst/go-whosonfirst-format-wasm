package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/whosonfirst/go-whosonfirst-format-wasm/static/javascript"
	"github.com/whosonfirst/go-whosonfirst-format-wasm/static/wasm"
)

//go:embed *.html
var FS embed.FS

func main() {

	var host string
	var port int

	flag.StringVar(&host, "host", "localhost", "...")
	flag.IntVar(&port, "port", 8080, "...")

	flag.Parse()

	html_fs := http.FS(FS)
	html_handler := http.FileServer(html_fs)

	wasm_fs := http.FS(wasm.FS)
	wasm_handler := http.FileServer(wasm_fs)

	javascript_fs := http.FS(javascript.FS)
	javascript_handler := http.FileServer(javascript_fs)

	mux := http.NewServeMux()

	mux.Handle("/javascript/", http.StripPrefix("/javascript/", javascript_handler))
	mux.Handle("/wasm/", http.StripPrefix("/wasm/", wasm_handler))
	mux.Handle("/", html_handler)

	addr := fmt.Sprintf("%s:%d", host, port)

	http_server := http.Server{
		Addr: addr,
	}

	http_server.Handler = mux

	fmt.Printf("Listening for requests at %s\n", addr)
	err := http_server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to serve requests, %w", err)
	}

}
