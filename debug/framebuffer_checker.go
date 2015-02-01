// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// fbChecker is like the checker type, but for a gfx.Framebuffer. It implicitly
// invokes the Check method of the underlying context after each function call
// is made.
type fbChecker struct {
	fb  gfx.Framebuffer
	ctx gfx.Context
}

// ClearColor implements the gfx.Clearable interface.
func (f *fbChecker) ClearColor(r, g, b, a float32) {
	f.fb.ClearColor(r, g, b, a)
	f.ctx.Check()
}

// ClearDepth implements the gfx.Clearable interface.
func (f *fbChecker) ClearDepth(depth float64) {
	f.fb.ClearDepth(depth)
	f.ctx.Check()
}

// ClearStencil implements the gfx.Clearable interface.
func (f *fbChecker) ClearStencil(stencil int) {
	f.fb.ClearStencil(stencil)
	f.ctx.Check()
}

// Clear implements the gfx.Clearable interface.
func (f *fbChecker) Clear(m gfx.ClearMask) {
	// Verify bitmask argument.
	if m == 0 {
		panic("Framebuffer.Clear: invalid clear mask argument (0)")
	}

	// Clearing all possible bits should yield zero.
	tmp := m
	tmp &^= gfx.ColorBuffer
	tmp &^= gfx.DepthBuffer
	tmp &^= gfx.StencilBuffer
	if tmp != 0 {
		panic("Framebuffer.Clear: invalid clear mask argument")
	}

	f.fb.Clear(m)
	f.ctx.Check()
}

// ReadPixelsUint8 implements the gfx.Framebuffer interface.
func (f *fbChecker) ReadPixelsUint8(x, y, width, height int, format gfx.PixelFormat, dataType gfx.PixelDataType, dst []uint8) {
	// Verify format argument.
	if format != gfx.RGBA {
		panic("Framebuffer.ReadPixelsUint8: invalid format (expect gfx.RGBA)")
	}

	// Verify dataType argument.
	if dataType != gfx.UnsignedByte {
		panic("Framebuffer.ReadPixelsUint8: invalid data type (expect gfx.UnsignedByte)")
	}

	// Verify destination buffer size.
	if len(dst) < width*height*4 {
		panic("Framebuffer.ReadPixelsUint8: dst buffer is not large enough")
	}

	f.fb.ReadPixelsUint8(x, y, width, height, format, dataType, dst)
	f.ctx.Check()
}
