package dockerw

import (
	"strings"
)

type WorkDir struct {
	value string
}

func (w *WorkDir) verify() error {
	return nil
}

func (w *WorkDir) toString() string {
	return w.value
}

// WorkDir consumes a directory to add to the docker writer as a WORKDIR
// directive.
func (w *Writer) WorkDir(dir string) *Writer {
	dir = strings.TrimSpace(dir)

	if !strings.HasPrefix(dir, "WORKDIR") {
		dir = "WORKDIR " + dir
	}

	w.addCommand(&From{dir})

	return w
}
