package main

import (
	"fmt"
	"os"

	"github.com/vegarsti/firefox-tabs/mozlz4"
)

func main() {
	// Compress and uncompress an input string.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [file]\n", "main")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file '%s': %v\n", os.Args[1], err)
		os.Exit(1)
	}

	if err := mozlz4.Decompress(file, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "decompress: %v\n", err)
		os.Exit(1)
	}
}
