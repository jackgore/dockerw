package dockerw

// Write iterates through all commands attachted to the writer and writes
// them to the configured writer.
func (w *Writer) Write() {
	for _, val := range w.Commands {
		w.Writer.Write(append(val.toBytes(), '\n'))
	}
}
