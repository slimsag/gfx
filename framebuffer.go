// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	Clearable

	// ReadPixelsUint8 reads RGBA 32/bpp pixel data into the given slice from a
	// rectangular area in the color buffer of this frame buffer.
	//
	// The x and y coordinates specify the frame buffer coordinates of the
	// first pixel that is read from the frame buffer. This location is the
	// lower left corner of the rectangular block of pixels.
	//
	// len(dst) must be >= width*height*4
	ReadPixelsUint8(x, y, width, height int, dst []uint8)

	// Status returns any framebuffer status error that might have occured. If
	// nil is returned, the framebuffer is ready for display.
	Status() error
}
