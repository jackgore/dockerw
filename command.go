package dockerw

// Command is a general interface implemented by different Dockerfile directives
type Command interface {
	verify() error
	toString() string
}
