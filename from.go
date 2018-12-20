package dockerw

import (
	"strings"
)

type From struct {
	value string
}

func (f *From) verify() error {
	return nil
}

func (f *From) toBytes() []byte {
	return []byte(f.value)
}

func (f *From) toString() string {
	return f.value
}

// FromString directly assigns the string to the From statement of the
// Dockerfile prepends 'FROM' if needed.
func (w *Writer) From(src string) *Writer {
	src = strings.TrimSpace(src)

	if !strings.HasPrefix(src, "FROM") {
		src = "FROM " + src
	}

	w.addCommand(&From{src})

	return w
}
