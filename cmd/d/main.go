package main

import (
	"fmt"
	"os"

	"github.com/vegarsti/tabs"
	"github.com/vegarsti/tabs/firefox"
	"github.com/vegarsti/tabs/sqlite"
)

func main() {
	firefoxPath := os.Getenv("FIREFOX_TABS_PATH")
	if firefoxPath == "" {
		fmt.Fprintln(os.Stderr, "set environment variable FIREFOX_TABS_PATH")
		os.Exit(1)
	}
	dbPath := os.Getenv("TABS_DB_PATH")
	if dbPath == "" {
		fmt.Fprintln(os.Stderr, "set environment variable TABS_DB_PATH")
		os.Exit(1)
	}
	file, err := os.Open(firefoxPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file '%s': %v\n", firefoxPath, err)
		os.Exit(1)
	}

	var t tabs.TabService

	t, err = firefox.NewTabService(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create firefox tab service: %v\n", err)
		os.Exit(1)
	}
	tabs, err := t.ReadTabs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	t.Close()
	t, err = sqlite.NewTabService(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create sqlite tab service: %v\n", err)
		os.Exit(1)
	}
	defer t.Close()
	if err := t.WriteTabs(tabs); err != nil {
		fmt.Fprintf(os.Stderr, "sqlite write tabs: %v\n", err)
		os.Exit(1)
	}
	if _, err := t.ReadTabs(); err != nil {
		fmt.Fprintf(os.Stderr, "sqlite read tabs: %v\n", err)
		os.Exit(1)
	}
}
