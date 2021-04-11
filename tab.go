package tabs

type Tab struct {
	URL          string
	Title        string
	LastAccessed int
	WindowIndex  int
	TabIndex     int
}

type Service interface {
	Read() ([]Tab, error)
}
