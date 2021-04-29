package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

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
	migration := "DROP TABLE IF EXISTS firefox; CREATE TABLE firefox (url text not null, at integer not null);"
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
	if len(tt) == 0 {
		return nil
	}
	vv := make([]string, len(tt))
	for i, t := range tt {
		vv[i] = fmt.Sprintf("('%s', %d)", t.URL, t.LastAccessed)
	}
	values := strings.Join(vv, ", ")
	result, err := s.db.Exec("INSERT INTO firefox (url, at) VALUES " + values + ";")
	if err != nil {
		return fmt.Errorf("insert tabs: %w", err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if int(n) != len(tt) {
		return fmt.Errorf("wrote %d rows, but expected to write %d", n, len(tt))
	}
	return nil
}

func (s *TabService) Close() error {
	return nil
}
