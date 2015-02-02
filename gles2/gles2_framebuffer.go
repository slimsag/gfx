// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"fmt"
	"unsafe"

	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
)

// Framebuffer implements the gfx.Framebuffer interface by wrapping a
// WebGLFramebuffer JavaScript object.
type Framebuffer struct {
	// Object is literally the OpenGL framebuffer object ID (or zero in the
	// case of the default framebuffer).
	Object uint32

	ctx *Context

	// State tied to this framebuffer object.
	clearColor   [4]float32
	clearDepth   float64
	clearStencil int
}

// useState binds the global OpenGL state for this local Framebuffer object.
func (f *Framebuffer) useState() {
	// Global OpenGL state.
	f.ctx.fastClearColor(f.clearColor)
	f.ctx.fastClearDepth(f.clearDepth)
	f.ctx.fastClearStencil(f.clearStencil)

	// Bind the framebuffer now.
	f.ctx.fastBindFramebuffer(f.Object)
}

// ClearColor implements the gfx.Clearable interface.
func (f *Framebuffer) ClearColor(r, g, b, a float32) {
	f.clearColor = [4]float32{r, g, b, a}
}

// ClearDepth implements the gfx.Clearable interface.
func (f *Framebuffer) ClearDepth(depth float64) {
	f.clearDepth = depth
}

// ClearStencil implements the gfx.Clearable interface.
func (f *Framebuffer) ClearStencil(stencil int) {
	f.clearStencil = stencil
}

// Clear implements the gfx.Framebuffer interface.
func (f *Framebuffer) Clear(m gfx.ClearMask) {
	var mask uint32
	if m&gfx.ColorBuffer != 0 {
		mask |= gl.COLOR_BUFFER_BIT
	}
	if m&gfx.DepthBuffer != 0 {
		mask |= gl.DEPTH_BUFFER_BIT
	}
	if m&gfx.StencilBuffer != 0 {
		mask |= gl.STENCIL_BUFFER_BIT
	}

	// Use this framebuffer's state, and perform the clear operation.
	f.useState()
	gl.Clear(mask)
}

// ReadPixelsUint8 implements the gfx.Framebuffer interface.
func (f *Framebuffer) ReadPixelsUint8(x, y, width, height int, dst []uint8) {
	f.useState()
	dstPtr := unsafe.Pointer(&dst[0])
	gl.ReadPixels(int32(x), int32(y), int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, dstPtr)
}

// Status implements the gfx.Framebuffer interface.
func (f *Framebuffer) Status() error {
	f.useState()
	e := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)

	// Avoid the larger switch statement below, as no error is the most likely
	// case.
	if e == gl.FRAMEBUFFER_COMPLETE {
		return nil
	}

	switch e {
	case gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		return gfx.ErrFramebufferIncompleteAttachment
	case gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		return gfx.ErrFramebufferIncompleteMissingAttachment
	case gl.FRAMEBUFFER_INCOMPLETE_DIMENSIONS:
		return gfx.ErrFramebufferIncompleteDimensions
	case gl.FRAMEBUFFER_UNSUPPORTED:
		return gfx.ErrFramebufferIncompleteDimensions
	default:
		panic(fmt.Sprintf("gl2: unhandled framebuffer status 0x%X\n", e))
	}
}

// Delete implements the gfx.Framebuffer interface.
func (f *Framebuffer) Delete() {
	gl.DeleteFramebuffers(1, &f.Object)
}
