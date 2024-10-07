package main

import (
	"flag"
	"fmt"
	"slices"

	"github.com/elhananf/prepare-dump-file/cmd"
)

func main() {
	fmt.Println("Prepare dump file")
	flag.Parse()
	var args = flag.Args()

	if len(args) != 2 || slices.Contains(args, "--help") {
		cmd.PrintHelp()
		return
	}

	inputPath, outputPath := args[0], args[1]
	fmt.Printf("%s -> %s\n", inputPath, outputPath)
}
