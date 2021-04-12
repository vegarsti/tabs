package main

import (
	"fmt"
	"os"

	"github.com/vegarsti/tabs"
	"github.com/vegarsti/tabs/firefox"
)

func main() {
	// Compress and uncompress an input string.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [mozlz4 file]\n", "main")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file '%s': %v\n", os.Args[1], err)
		os.Exit(1)
	}

	var t tabs.TabService
	t = firefox.NewTabService(file)
	if _, err := t.ReadTabs(); err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
}
