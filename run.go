package dockerw

import (
	"fmt"
	"strings"
)

type Run struct {
	value string
}

func (c *Run) verify() error {
	return nil
}

func (c *Run) toBytes() []byte {
	return []byte(c.value)
}

func (c *Run) toString() string {
	return c.value
}

// Run consumes commands to add to a run directive
func (w *Writer) Run(cmds []string) *Writer {
	s := strings.Builder{}
	s.WriteString("RUN [")

	for i, cmd := range cmds {
		s.WriteString(fmt.Sprintf("\"%v\"", cmd))
		if i < len(cmds)-1 {
			s.WriteString(",")
		}
	}

	s.WriteString("]")

	w.addCommand(&Run{s.String()})

	return w
}
