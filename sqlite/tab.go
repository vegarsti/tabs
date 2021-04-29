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
	migration := `
		DROP TABLE IF EXISTS firefox;
		CREATE TABLE firefox (url text not null, title text not null, at integer not null);
	`
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

func (s *TabService) WriteTabs(tt []tabs.Tab) error {
	for _, t := range tt {
		insert := "INSERT INTO firefox (url, title, at) VALUES ($1, $2, $3);"
		result, err := s.db.Exec(insert, t.URL, t.Title, t.LastAccessed)
		if err != nil {
			return fmt.Errorf("insert tabs: %w", err)
		}
		n, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("rows affected: %w", err)
		}
		if int(n) != 1 {
			return fmt.Errorf("expected to write a row")
		}
	}
	return nil
}

func (s *TabService) Close() error {
	return nil
}
