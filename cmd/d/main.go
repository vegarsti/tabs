package main

import (
	"fmt"
	"os"

	"github.com/vegarsti/tabs/firefox"
	"github.com/vegarsti/tabs/sqlite"
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

	t1, err := firefox.NewTabService(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create firefox tab service: %v\n", err)
		os.Exit(1)
	}
	tabs, err := t1.ReadTabs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	t2, err := sqlite.NewTabService("lol.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create sqlite tab service: %v\n", err)
		os.Exit(1)
	}
	if err := t2.WriteTabs(tabs); err != nil {
		fmt.Fprintf(os.Stderr, "sqlite write tabs: %v\n", err)
		os.Exit(1)
	}
	if _, err := t2.ReadTabs(); err != nil {
		fmt.Fprintf(os.Stderr, "sqlite read tabs: %v\n", err)
		os.Exit(1)
	}
}
