package mozlz4

import (
	"fmt"
	"io"

	"github.com/frioux/leatherman/pkg/mozlz4"
)

// Decompress a mozlz4 compressed file r into w
func Decompress(r io.Reader, w io.Writer) error {
	lr, err := mozlz4.NewReader(r)
	if err != nil {
		return fmt.Errorf("mozlz4.NewReader: %w", err)
	}
	_, err = io.Copy(w, lr)
	if err != nil {
		return fmt.Errorf("Couldn't copy: %w", err)
	}
	return nil
}
