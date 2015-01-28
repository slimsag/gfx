package gfx

import "errors"

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidValue     = errors.New("invalid value")
	ErrOutOfMemory      = errors.New("out of memory")
)

// Context is a graphics context. Unlike traditional OpenGL contexts, it is not
// tied to a specific OS thread (but still must be accessed from only one
// goroutine/thread at a time).
type Context interface {
	// NewBuffer returns a new buffer object that contains data such as
	// vertices or colors.
	NewBuffer() Buffer

	// Attributes returns a structure filled with this context's in-use
	// attributes.
	Attributes(a ContextAttributes) ContextAttributes

	// Error returns the last error that occured in the context, or nil.
	Error() error
}
