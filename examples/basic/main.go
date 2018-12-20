package main

import (
    "os"

	"github.com/JonathonGore/dockerw"
)

func main() {
	writer := dockerw.Writer{
        Writer: os.Stdout,
    }

	writer.From("golang:1.11").
		WorkDir("/go/github.com/JonathonGore/dockerw").
		Copy(".", "/go/github.com/JonathonGore/dockerw")
	writer.Write()
}
