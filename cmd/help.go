package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintHelp() {
	var executablePath = os.Args[0]
	var executableFileName = filepath.Base(executablePath)
	fmt.Printf("Usage: %s <input-path> <output-path> [--help]\n", executableFileName)
}
