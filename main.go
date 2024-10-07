package main

import (
	"flag"
	"fmt"
	"slices"

	"github.com/elhananf/prepare-dump-file/cmd"
)

func main() {
	fmt.Println("Prepare dump file restore")
	flag.Parse()
	var args = flag.Args()
	if len(args) != 2 || slices.Contains(args, "--help") {
		cmd.PrintHelp()
	}
}
