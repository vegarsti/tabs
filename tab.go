package tabs

type Tab struct {
	URL   string `json:"originalURI"`
	Title string `json:"title"`
}

type Read interface {
	Tabs() []Tab
}
