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
	s "github.com/slimsag/gfx/internal/state"
)

// Framebuffer implements the gfx.Framebuffer interface by wrapping a
// WebGLFramebuffer JavaScript object.
type Framebuffer struct {
	s.Framebuffer

	// o is literally the OpenGL framebuffer object ID (or zero in the case of the
	// default framebuffer).
	o uint32

	ctx *Context
}

// useState binds the global OpenGL state for this local Framebuffer object.
func (f *Framebuffer) useState() {
	// Bind the framebuffer now.
	if f.ctx.fastBindFramebuffer(f.o) {
		f.GLCall(nil)
	}
	f.GLCall(f.Loaded)
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

// Texture2D implements the gfx.Framebuffer interface.
func (f *Framebuffer) Texture2D(attachment gfx.FramebufferAttachment, target gfx.TextureTarget, tex gfx.Texture) {
	f.useState()
	gl.FramebufferTexture2D(
		gl.FRAMEBUFFER,
		f.ctx.Enums[int(attachment)],
		f.ctx.Enums[int(target)],
		tex.Object().(uint32),
		0,
	)
}

// Renderbuffer implements the gfx.Framebuffer interface.
func (f *Framebuffer) Renderbuffer(attachment gfx.FramebufferAttachment, buf gfx.Renderbuffer) {
	f.useState()
	gl.FramebufferTexture2D(
		gl.FRAMEBUFFER,
		f.ctx.Enums[int(attachment)],
		gl.RENDERBUFFER,
		buf.Object().(uint32),
		0,
	)
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

// Delete implements the gfx.Object interface.
func (f *Framebuffer) Delete() {
	if f.o == 0 {
		return
	}
	gl.DeleteFramebuffers(1, &f.o)
	f.o = 0
}

// Object implements the gfx.Object interface.
func (f *Framebuffer) Object() interface{} {
	return f.o
}

const (
	csClearColor = iota
	csClearDepth
	csClearStencil
)

func glClearColor(v interface{}) {
	x := v.([4]float32)
	gl.ClearColor(x[0], x[1], x[2], x[3])
}

// ClearColor implements the gfx.ContextStateProvider interface.
func (f *Framebuffer) ClearColor(r, g, b, a float32) gfx.FramebufferStateValue {
	return s.CSV{
		Value:        [4]float32{r, g, b, a},
		DefaultValue: [4]float32{0, 0, 0, 0}, // TODO(slimsag): verify
		Key:          csClearColor,
		GLCall:       glClearColor,
	}
}

func glClearDepth(v interface{}) {
	gl.ClearDepthf(v.(float32))
}

// ClearDepth implements the gfx.ContextStateProvider interface.
func (f *Framebuffer) ClearDepth(depth float64) gfx.FramebufferStateValue {
	return s.CSV{
		Value:        float32(depth),
		DefaultValue: float32(0), // TODO(slimsag): verify
		Key:          csClearDepth,
		GLCall:       glClearDepth,
	}
}

func glClearStencil(v interface{}) {
	gl.ClearStencil(v.(int32))
}

// ClearStencil implements the gfx.ContextStateProvider interface.
func (f *Framebuffer) ClearStencil(stencil int) gfx.FramebufferStateValue {
	return s.CSV{
		Value:        int32(stencil),
		DefaultValue: int32(0), // TODO(slimsag): verify
		Key:          csClearStencil,
		GLCall:       glClearStencil,
	}
}
