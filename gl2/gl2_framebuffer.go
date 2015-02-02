// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"fmt"
	"unsafe"

	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
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

// Texture2D implements the gfx.Framebuffer interface.
func (f *Framebuffer) Texture2D(attachment gfx.FramebufferAttachment, target gfx.TextureTarget, tex gfx.Texture) {
	// Convert the attachment parameter.
	var a uint32
	switch attachment {
	case gfx.ColorAttachment0:
		a = gl.COLOR_ATTACHMENT0
	case gfx.DepthAttachment:
		a = gl.DEPTH_ATTACHMENT
	case gfx.StencilAttachment:
		a = gl.STENCIL_ATTACHMENT
	case gfx.DepthStencilAttachment:
		a = gl.DEPTH_STENCIL_ATTACHMENT
	default:
		panic("Framebuffer.Texture2D: invalid framebuffer attachment parameter")
	}

	// Convert the target parameter.
	var t uint32
	switch target {
	case gfx.Texture2D:
		t = gl.TEXTURE_2D
	case gfx.TextureCubeMapPositiveX:
		t = gl.TEXTURE_CUBE_MAP_POSITIVE_X
	case gfx.TextureCubeMapNegativeX:
		t = gl.TEXTURE_CUBE_MAP_NEGATIVE_X
	case gfx.TextureCubeMapPositiveY:
		t = gl.TEXTURE_CUBE_MAP_POSITIVE_Y
	case gfx.TextureCubeMapNegativeY:
		t = gl.TEXTURE_CUBE_MAP_NEGATIVE_Y
	case gfx.TextureCubeMapPositiveZ:
		t = gl.TEXTURE_CUBE_MAP_POSITIVE_Z
	case gfx.TextureCubeMapNegativeZ:
		t = gl.TEXTURE_CUBE_MAP_NEGATIVE_Z
	default:
		panic("Framebuffer.Texture2D: invalid texture target parameter")
	}
	f.useState()
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, a, t, tex.(*Texture).Object, 0)
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
	//case gl.FRAMEBUFFER_INCOMPLETE_DIMENSIONS:
	//	return gfx.ErrFramebufferIncompleteDimensions
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
