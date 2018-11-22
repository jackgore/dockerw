package dockerw

// Writer contains specifiers for writing a Dockerfile
type Writer struct {
	// Commands contains an order list of commands
	Commands []Command
}

// AddCommand adds the provided command c to the list of commands
func (w *Writer) addCommand(c Command) {
	w.Commands = append(w.Commands, c)
}
