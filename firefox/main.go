package firefox

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/vegarsti/tabs"
	"github.com/vegarsti/tabs/firefox/mozlz4"
)

type Firefox struct {
	file io.Reader
}

func New(r io.Reader) *Firefox {
	return &Firefox{file: r}
}

func (f *Firefox) Read() []tabs.Tab {
	bs := new(bytes.Buffer)
	if err := mozlz4.Decompress(f.file, bs); err != nil {
		fmt.Fprintf(os.Stderr, "decompress: %v\n", err)
		os.Exit(1)
	}
	// log.Println(string(bs.Bytes()[:10]))
	return []tabs.Tab{}
}
