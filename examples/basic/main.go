package main

import (
    "github.com/JonathonGore/dockerw"
)

func main() {
    writer := dockerw.Writer{}
    writer.From("go:10")
    writer.Write()
}
