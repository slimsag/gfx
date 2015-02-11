// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
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
	r.useState()
	gl.RenderbufferStorage(gl.RENDERBUFFER, r.ctx.Enums[int(internalFormat)], int32(width), int32(height))
}

// Delete implements the gfx.Object interface.
func (r *Renderbuffer) Delete() {
	if r.Object == 0 {
		return
	}
	gl.DeleteRenderbuffers(1, &r.Object)
	r.Object = 0
}
