package dockerw

import (
    "io"
)

// Writer contains specifiers for writing a Dockerfile
type Writer struct {
    // Specifies where the Dockerfile will be written to
    Writer io.Writer

	// Commands contains an ordered list of commands
	Commands []Command
}

// AddCommand adds the provided command c to the list of commands
func (w *Writer) addCommand(c Command) {
	w.Commands = append(w.Commands, c)
}
