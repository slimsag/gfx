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

// ClearMask represents a bitmask to choose which buffers to clear during a
// framebuffer clearing operation. It must be one of the predefined constants.
type ClearMask int

// Clearing masks to select the color, depth, and stencil buffers. They can be
// bitwise OR'd together to select multiple:
//
//  colorAndDepth := ColorBufer|DepthBuffer
//
const (
	ColorBuffer ClearMask = 1 << iota
	DepthBuffer
	StencilBuffer
)

// Framebuffer is a collection of buffers that serve as a rendering
// destination.
type Framebuffer interface {
	Object
	FramebufferStateProvider

	// Clear clears the buffers selected by the bitmask to their respective
	// clear values. Multiple bitmasks can be OR'd together to select multiple
	// buffers to clear at once:
	//
	//  // Clear both(!) the color and depth buffers in one call.
	//  Clear(ColorBuffer|DepthBuffer)
	//
	Clear(m ClearMask)

	// ReadPixelsUint8 reads RGBA 32/bpp pixel data into the given slice from a
	// rectangular area in the color buffer of this frame buffer.
	//
	// The x and y coordinates specify the frame buffer coordinates of the
	// first pixel that is read from the frame buffer. This location is the
	// lower left corner of the rectangular block of pixels.
	//
	// len(dst) must be >= width*height*4
	ReadPixelsUint8(x, y, width, height int, dst []uint8)

	// Texture2D attaches a 2D texture to this framebuffer object.
	Texture2D(attachment FramebufferAttachment, target TextureTarget, tex Texture)

	// Renderbuffer attaches a renderbuffer to this framebuffer object.
	Renderbuffer(attachment FramebufferAttachment, buf Renderbuffer)

	// Status returns any framebuffer status error that might have occured. If
	// nil is returned, the framebuffer is ready for display.
	//
	// Primarily you should expect to handle ErrFramebufferUnsupported, which
	// is returned when the framebuffer attachment combination is not supported
	// by the hardware.
	Status() error
}

// FramebufferStateValue represents a single value as part of a framebuffer's
// state, for example the clear color.
//
// The underlying type is platform-specific, do not access it directly or make
// assumptions about it.
type FramebufferStateValue interface{}

// FramebufferState solely represents a framebuffer's unique state. Any values
// not explicitly specified are assumed to be their defaults.
//
// The underlying type is platform-specific, do not access it directly or make
// assumptions about it.
type FramebufferState interface{}

// FramebufferStateProvider provides access to a framebuffer's state.
type FramebufferStateProvider interface {
	// NewFramebufferState returns a new framebuffer state for the given values.
	NewFramebufferState(values ...FramebufferStateValue) FramebufferState

	// LoadFramebufferState loads the given framebuffer state, replacing the
	// previous one. If s == nil then the default state is loaded.
	LoadFramebufferState(s FramebufferState)

	// ClearColor sets the color to clear the color buffer to upon a call to
	// the Clear method.
	ClearColor(r, g, b, a float32) FramebufferStateValue

	// ClearDepth sets the value to clear the depth buffer to upon a depth
	// buffer clearing operation (a call to Clear with the DepthBuffer clear
	// mask)
	ClearDepth(depth float64) FramebufferStateValue

	// ClearStencil sets the value to clear the stencil buffer to upon a
	// stencil buffer clearing operation (a call to Clear with the
	// StencilBuffer clear mask)
	ClearStencil(stencil int) FramebufferStateValue
}
