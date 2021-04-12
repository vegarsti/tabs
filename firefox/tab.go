package firefox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vegarsti/tabs"
	"github.com/vegarsti/tabs/firefox/mozlz4"
)

type TabService struct {
	file io.Reader
}

func NewTabService(r io.Reader) *TabService {
	return &TabService{file: r}
}

func (f *TabService) ReadTabs() ([]tabs.Tab, error) {
	bs := new(bytes.Buffer)
	if err := mozlz4.Decompress(f.file, bs); err != nil {
		return nil, fmt.Errorf("decompress: %w", err)
	}
	p := Payload{}
	if err := json.Unmarshal(bs.Bytes(), &p); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	tt := make([]tabs.Tab, 0)
	for i, w := range p.Windows {
		for j, t := range w.Tabs {
			current := t.Entries[len(t.Entries)-1]
			tt = append(tt, tabs.Tab{
				Title:        current.Title,
				URL:          current.URL,
				LastAccessed: t.LastAccessed,
				WindowIndex:  i,
				TabIndex:     j,
			})
		}
	}
	return tt, nil
}

type Payload struct {
	Windows []struct {
		Tabs []struct {
			LastAccessed int `json:"lastAccessed"`
			Entries      []struct {
				CacheKey           int    `json:"cacheKey"`
				DocIdentifier      int    `json:"docIdentifier"`
				DocshellUUID       string `json:"docshellUUID"`
				HasUserInteraction bool   `json:"hasUserInteraction"`
				Persist            bool   `json:"persist"`
				OriginalURI        string `json:"originalURI"`
				ResultPrincipalURI string `json:"resultPrincipalURI"`
				Title              string `json:"title"`
				URL                string `json:"url"`
			} `json:"entries"`
		} `json:"tabs"`
	} `json:"windows"`
}
