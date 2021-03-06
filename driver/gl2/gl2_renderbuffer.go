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
	// o is literally the OpenGL renderbuffer object ID.
	o uint32

	ctx *Context
}

// useState binds the global OpenGL state for this local Renderbuffer object.
func (r *Renderbuffer) useState() {
	// Bind the renderbuffer now.
	r.ctx.fastBindRenderbuffer(r.o)
}

// Storage implements the gfx.Renderbuffer interface.
func (r *Renderbuffer) Storage(internalFormat gfx.RenderbufferFormat, width, height int) {
	r.useState()
	gl.RenderbufferStorage(gl.RENDERBUFFER, r.ctx.Enums[int(internalFormat)], int32(width), int32(height))
}

// Delete implements the gfx.Object interface.
func (r *Renderbuffer) Delete() {
	if r.o == 0 {
		return
	}
	gl.DeleteRenderbuffers(1, &r.o)
	r.o = 0
}

// Object implements the gfx.Object interface.
func (r *Renderbuffer) Object() interface{} {
	return r.o
}
