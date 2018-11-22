package dockerw

import (
    "fmt"
)

// Write iterates through all commands attachted to the writer and writes
// them to the configured writer.
func (w *Writer) Write() {
    for _, val := range w.Commands {
        fmt.Printf("%v\n", val.toString())
    }
}
