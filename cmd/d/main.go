package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	interruptChannel := make(chan os.Signal, 1)
	defer close(interruptChannel)
	signal.Notify(interruptChannel, os.Interrupt, syscall.SIGTERM)
	ticks := make(chan struct{}, 1)
	defer close(ticks)
	go sendTick(ticks, 500*time.Millisecond)

	for {
		select {
		case <-interruptChannel:
			fmt.Println()
			return
		case <-ticks:
			if err := run(firefoxPath, dbPath); err != nil {
				close(interruptChannel)
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}

func run(firefoxPath string, dbPath string) error {
	var t tabs.TabService
	t, err := firefox.NewTabService(firefoxPath)
	if err != nil {
		return fmt.Errorf("couldn't create firefox tab service: %w", err)
	}
	tabs, err := t.ReadTabs()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	t.Close()
	t, err = sqlite.NewTabService(dbPath)
	if err != nil {
		return fmt.Errorf("couldn't create sqlite tab service: %w", err)
	}
	defer t.Close()
	if err := t.WriteTabs(tabs); err != nil {
		return fmt.Errorf("sqlite write tabs: %w", err)
	}
	if _, err := t.ReadTabs(); err != nil {
		return fmt.Errorf("sqlite read tabs: %w", err)
	}
	return nil
}

func sendTick(ch chan<- struct{}, rate time.Duration) {
	ch <- struct{}{}
	for range time.Tick(rate) {
		ch <- struct{}{}
	}
}
