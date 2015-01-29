// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64 386

package gl2

import (
	"github.com/go-gl/glow/gl/2.1/gl"
	"github.com/slimsag/gfx"
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