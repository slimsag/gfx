// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
)

// Renderbuffer implements the gfx.Renderbuffer interface by wrapping a OpenGL
// renderbuffer object ID.
type Renderbuffer struct {
	// Object is literally the OpenGL renderbuffer object ID.
	Object uint32

	ctx *Context
}

// useState binds the global OpenGL state for this local Renderbuffer object.
func (r *Renderbuffer) useState() {
	// Bind the renderbuffer now.
	r.ctx.fastBindRenderbuffer(r.Object)
}

// Storage implements the gfx.Renderbuffer interface.
func (r *Renderbuffer) Storage(internalFormat gfx.RenderbufferFormat, width, height int) {
	var f uint32
	switch internalFormat {
	case gfx.RGBA4:
		f = gl.RGBA4
	case gfx.RGB565:
		f = gl.RGB565
	case gfx.RGB5A1:
		f = gl.RGB5_A1
	case gfx.DepthComponent16:
		f = gl.DEPTH_COMPONENT16
	default:
		panic("Renderbuffer.Storage: invalid internalFormat parameter")
	}

	r.useState()
	gl.RenderbufferStorage(gl.RENDERBUFFER, f, int32(width), int32(height))
}

// Delete implements the gfx.Renderbuffer interface.
func (r *Renderbuffer) Delete() {
	gl.DeleteRenderbuffers(1, &r.Object)
}
