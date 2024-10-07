package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello, Mars!")
	flag.Parse()
	var args = flag.Args()
	fmt.Printf("%s\n", args)
}
