// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// RenderbufferFormat represents a renderbuffer's storage format.
type RenderbufferFormat int

const (
	zeroRenderbufferFormat RenderbufferFormat = iota

	// RGBA4 is a 4-bits per channel renderbuffer storage format.
	RGBA4

	// RGB565 is a renderbuffer storage format with 5 bits red, 6 bits green,
	// and 5 bits blue, respectively.
	RGB565

	// RGB5A1 is a renderbuffer storage format with 5 bits for RGB and 1 bit
	// for alpha.
	RGB5A1

	// DepthComponent16 is a renderbuffer storage format with
	DepthComponent16
)

// RenderBuffer represents a buffer which can contain an image and act as the
// source or target of a render operation.
//
// It stores data that arepresents a single image. A Renderbuffer can be
// attached to a Framebuffer and used to store either stencil data or a
// combination of stencil and depth data.
type Renderbuffer interface {
	// Storage creates and initailizes this renderbuffer object's data store.
	Storage(internalFormat RenderbufferFormat, width, height int)

	// Delete deletes this renderbuffer object, it is unsafe to use this
	// renderbuffer after deletion.
	Delete()
}
