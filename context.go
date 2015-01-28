package gfx

import "errors"

// Panics that Context.Error may generate.
var (
	OutOfMemory                 = errors.New("out of memory")
	InvalidEnum                 = errors.New("invalid enum")
	InvalidOperation            = errors.New("invalid operation")
	InvalidFramebufferOperation = errors.New("invalid framebuffer operation")
	InvalidValue                = errors.New("invalid value")
	ContextLost                 = errors.New("context lost")
)

// Context is a graphics context. Unlike traditional OpenGL contexts, it is not
// tied to a specific OS thread (but still must be accessed from only one
// goroutine/thread at a time).
type Context interface {
	// Framebuffer is the default framebuffer of the context, i.e. the one that
	// when drawn to, appears on the window.
	Framebuffer

	// NewBuffer returns a new buffer object that contains data such as
	// vertices or colors.
	NewBuffer() Buffer

	// Check checks that no errors have occured in the context. If an error
	// occurs, it is either a programmer error (passing an invalid value, etc)
	// or a serious device error (running out of memory, losing the context).
	Check()
}
