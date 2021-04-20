package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // sqlite driver
	"github.com/vegarsti/tabs"
)

type TabService struct {
	file string
	db   *sql.DB
}

func NewTabService(file string) (*TabService, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("sqlite open '%s': %w", file, err)
	}
	migration := "CREATE TABLE firefox (url text not null, at integer not null);"
	if _, err := db.Exec(migration); err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}
	return &TabService{
		file: file,
		db:   db,
	}, nil
}

func (s *TabService) ReadTabs() ([]tabs.Tab, error) {
	return nil, nil
}

func (s *TabService) WriteTabs([]tabs.Tab) error {
	return nil
}
