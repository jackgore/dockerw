package dockerw

import (
	"fmt"
	"strings"
)

type Cmd struct {
	value string
}

func (c *Cmd) verify() error {
	return nil
}

func (c *Cmd) toBytes() []byte {
	return []byte(c.value)
}

func (c *Cmd) toString() string {
	return c.value
}

// Cmd consumes commands to add to a cmd directive
func (w *Writer) Cmd(cmds []string) *Writer {
	s := strings.Builder{}
	s.WriteString("CMD [")

	for i, cmd := range cmds {
		s.WriteString(fmt.Sprintf("\"%v\"", cmd))
		if i < len(cmds)-1 {
			s.WriteString(",")
		}
	}

	s.WriteString("]")

	w.addCommand(&Cmd{s.String()})

	return w
}
