package gfx

import "errors"

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidValue     = errors.New("invalid value")
	ErrOutOfMemory      = errors.New("out of memory")
)

// ContextAttributes is a set of properties used to initialize the context for
// the first time.
//
// Each property is a preference, and might not be configurable. To query which
// properties you did receive, check Context.GetAttributes()
type ContextAttributes struct {
	// Alpha is whether the drawing buffer has an alpha channel for alpha
	// compositing operations.
	Alpha bool

	// Antialias is whether the rendering context is antialiased.
	Antialias bool

	// Depth is whether or not the drawing buffer has an depth buffer.
	Depth bool

	// PremultipliedAlpha is whether or not the colors in the drawing buffer
	// are premultiplied alpha. This property is ignored if Alpha is false.
	PremultipliedAlpha bool

	// FailIfMajorPerformanceCaveat is whether a context will be created if
	// system performance is low.
	FailIfMajorPerformanceCaveat bool

	// PreserveDrawingBuffer is whether the drawing buffers are cleared once
	// content has been presented to the user.
	PreserveDrawingBuffer bool
}

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
