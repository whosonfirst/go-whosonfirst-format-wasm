package main

import (
	"flag"
	"fmt"

	wof_format "github.com/whosonfirst/go-whosonfirst-format"
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args() {
		fmt.Println(format(raw))
	}
}

//export format
func format(raw string) string {

	out, err := wof_format.FormatBytes([]byte(raw))

	if err != nil {
		return fmt.Sprintf("Failed to format data, %v", err)
	}

	return string(out)
}
