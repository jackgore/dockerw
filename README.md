# Dockerw

Dockerw is a set of utilities for programmtically writing a Dockerfile.

```
package main

import (
	"os"

	"github.com/jackgore/dockerw"
)

func main() {
	writer := dockerw.Writer{
		Writer: os.Stdout,
	}

	writer.From("golang:1.11").
		WorkDir("/go/github.com/jackgore/dockerw").
		Copy(".", "/go/github.com/jackgore/dockerw")
	writer.Write()
}
```
