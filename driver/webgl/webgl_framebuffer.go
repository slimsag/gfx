// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
	s "github.com/slimsag/gfx/internal/state"
)

// Framebuffer implements the gfx.Framebuffer interface by wrapping a
// WebGLFramebuffer JavaScript object.
type Framebuffer struct {
	s.Framebuffer

	// o is literally the WebGLFramebuffer object (or nil in the case of the
	// default framebuffer).
	o *js.Object

	ctx *Context

	// State tied to this framebuffer object.
	clearColor   [4]float32
	clearDepth   float64
	clearStencil int
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
	var mask int
	if m&gfx.ColorBuffer != 0 {
		mask |= f.ctx.COLOR_BUFFER_BIT
	}
	if m&gfx.DepthBuffer != 0 {
		mask |= f.ctx.DEPTH_BUFFER_BIT
	}
	if m&gfx.StencilBuffer != 0 {
		mask |= f.ctx.STENCIL_BUFFER_BIT
	}

	// Use this framebuffer's state, and perform the clear operation.
	f.useState()
	f.ctx.O.Call("clear", mask)
}

// ReadPixelsUint8 implements the gfx.Framebuffer interface.
func (f *Framebuffer) ReadPixelsUint8(x, y, width, height int, dst []uint8) {
	f.useState()
	f.ctx.O.Call("readPixels", x, y, width, height, f.ctx.RGBA, f.ctx.UNSIGNED_BYTE, dst)
}

// Texture2D implements the gfx.Framebuffer interface.
func (f *Framebuffer) Texture2D(attachment gfx.FramebufferAttachment, target gfx.TextureTarget, tex gfx.Texture) {
	f.useState()
	f.ctx.O.Call(
		"framebufferTexture2D",
		f.ctx.FRAMEBUFFER,
		f.ctx.Enums[int(attachment)],
		f.ctx.Enums[int(target)],
		tex.Object().(*js.Object),
		0,
	)
}

// Renderbuffer implements the gfx.Framebuffer interface.
func (f *Framebuffer) Renderbuffer(attachment gfx.FramebufferAttachment, buf gfx.Renderbuffer) {
	f.useState()
	f.ctx.O.Call(
		"framebufferTexture2D",
		f.ctx.FRAMEBUFFER,
		f.ctx.Enums[int(attachment)],
		f.ctx.RENDERBUFFER,
		buf.Object().(*js.Object),
		0,
	)
}

// Status implements the gfx.Framebuffer interface.
func (f *Framebuffer) Status() error {
	f.useState()
	e := f.ctx.O.Call("checkFramebufferStatus", f.ctx.FRAMEBUFFER).Int()

	// Avoid the larger switch statement below, as no error is the most likely
	// case.
	if e == f.ctx.FRAMEBUFFER_COMPLETE {
		return nil
	}

	switch e {
	case f.ctx.FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		return gfx.ErrFramebufferIncompleteAttachment
	case f.ctx.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		return gfx.ErrFramebufferIncompleteMissingAttachment
	case f.ctx.FRAMEBUFFER_INCOMPLETE_DIMENSIONS:
		return gfx.ErrFramebufferIncompleteDimensions
	case f.ctx.FRAMEBUFFER_UNSUPPORTED:
		return gfx.ErrFramebufferIncompleteDimensions
	default:
		panic(fmt.Sprintf("webgl: unhandled framebuffer status 0x%X\n", e))
	}
}

// Delete implements the gfx.Object interface.
func (f *Framebuffer) Delete() {
	if f.o == nil {
		return
	}
	f.ctx.O.Call("deleteFramebuffer", f.o)
	f.o = nil
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

func (c *Context) glClearColor(v interface{}) {
	x := v.([4]float32)
	c.O.Call("clearColor", x[0], x[1], x[2], x[3])
}

// ClearColor implements the gfx.ContextStateProvider interface.
func (f *Framebuffer) ClearColor(r, g, b, a float32) gfx.FramebufferStateValue {
	return s.CSV{
		Value:        [4]float32{r, g, b, a},
		DefaultValue: [4]float32{0, 0, 0, 0}, // TODO(slimsag): verify
		Key:          csClearColor,
		GLCall:       f.ctx.glClearColor,
	}
}

func (c *Context) glClearDepth(v interface{}) {
	c.O.Call("clearDepth", v.(float64))
}

// ClearDepth implements the gfx.ContextStateProvider interface.
func (f *Framebuffer) ClearDepth(depth float64) gfx.FramebufferStateValue {
	return s.CSV{
		Value:        depth,
		DefaultValue: 0, // TODO(slimsag): verify
		Key:          csClearDepth,
		GLCall:       f.ctx.glClearDepth,
	}
}

func (c *Context) glClearStencil(v interface{}) {
	c.O.Call("clearStencil", v.(int))
}

// ClearStencil implements the gfx.ContextStateProvider interface.
func (f *Framebuffer) ClearStencil(stencil int) gfx.FramebufferStateValue {
	return s.CSV{
		Value:        stencil,
		DefaultValue: 0, // TODO(slimsag): verify
		Key:          csClearStencil,
		GLCall:       f.ctx.glClearStencil,
	}
}
