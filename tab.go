package tabs

type Tab struct {
	URL          string
	Title        string
	LastAccessed int
	WindowIndex  int
	TabIndex     int
}

type TabService interface {
	ReadTabs() ([]Tab, error)
	WriteTabs([]Tab) error
}
