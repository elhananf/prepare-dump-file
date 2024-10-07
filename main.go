package main

import (
	"flag"
	"fmt"
	"log"
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

	err := cmd.PrepareDump(inputPath, outputPath)
	if err != nil {
		log.Fatal(err)
	}
}
