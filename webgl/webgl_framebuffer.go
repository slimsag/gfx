// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Framebuffer implements the gfx.Framebuffer interface by wrapping a
// WebGLFramebuffer JavaScript object.
type Framebuffer struct {
	// Object is literally the WebGLFramebuffer object (or nil in the case of
	// the default framebuffer).
	Object js.Object

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
	f.ctx.Object.Call("clear", mask)
}

// ReadPixelsUint8 implements the gfx.Framebuffer interface.
func (f *Framebuffer) ReadPixelsUint8(x, y, width, height int, dst []uint8) {
	f.useState()
	f.ctx.Object.Call("readPixels", x, y, width, height, f.ctx.RGBA, f.ctx.UNSIGNED_BYTE, dst)
}

// Texture2D implements the gfx.Framebuffer interface.
func (f *Framebuffer) Texture2D(attachment gfx.FramebufferAttachment, target gfx.TextureTarget, tex gfx.Texture) {
	// Convert the attachment parameter.
	var a int
	switch attachment {
	case gfx.ColorAttachment0:
		a = f.ctx.COLOR_ATTACHMENT0
	case gfx.DepthAttachment:
		a = f.ctx.DEPTH_ATTACHMENT
	case gfx.StencilAttachment:
		a = f.ctx.STENCIL_ATTACHMENT
	case gfx.DepthStencilAttachment:
		a = f.ctx.DEPTH_STENCIL_ATTACHMENT
	default:
		panic("Framebuffer.Texture2D: invalid framebuffer attachment parameter")
	}

	// Convert the target parameter.
	var t int
	switch target {
	case gfx.Texture2D:
		t = f.ctx.TEXTURE_2D
	case gfx.TextureCubeMapPositiveX:
		t = f.ctx.TEXTURE_CUBE_MAP_POSITIVE_X
	case gfx.TextureCubeMapNegativeX:
		t = f.ctx.TEXTURE_CUBE_MAP_NEGATIVE_X
	case gfx.TextureCubeMapPositiveY:
		t = f.ctx.TEXTURE_CUBE_MAP_POSITIVE_Y
	case gfx.TextureCubeMapNegativeY:
		t = f.ctx.TEXTURE_CUBE_MAP_NEGATIVE_Y
	case gfx.TextureCubeMapPositiveZ:
		t = f.ctx.TEXTURE_CUBE_MAP_POSITIVE_Z
	case gfx.TextureCubeMapNegativeZ:
		t = f.ctx.TEXTURE_CUBE_MAP_NEGATIVE_Z
	default:
		panic("Framebuffer.Texture2D: invalid texture target parameter")
	}
	f.useState()
	f.ctx.Object.Call("framebufferTexture2D", f.ctx.FRAMEBUFFER, a, t, tex.(*Texture).Object, 0)
}

// Status implements the gfx.Framebuffer interface.
func (f *Framebuffer) Status() error {
	f.useState()
	e := f.ctx.Object.Call("checkFramebufferStatus", f.ctx.FRAMEBUFFER).Int()

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

// Delete implements the gfx.Framebuffer interface.
func (f *Framebuffer) Delete() {
	f.ctx.Object.Call("deleteFramebuffer", f.Object)
}
