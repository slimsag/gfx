package gfx

import "errors"

var (
	ErrFramebufferIncompleteAttachment        = errors.New("framebuffer: attachment types are mismatched")
	ErrFramebufferIncompleteMissingAttachment = errors.New("framebuffer: missing attachment")
	ErrFramebufferIncompleteDimensions        = errors.New("framebuffer: the width and height of the attachments are not the same")
	ErrFramebufferUnsupported                 = errors.New("framebuffer: the attachments aren't supported")
)

// Framebuffer is a collection of buffers that serve as a rendering
// destination.
type Framebuffer interface {
	// Status returns any framebuffer status error that might have occured. If
	// nil is returned, the framebuffer is ready for display.
	Status() error
}
