// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Renderbuffer implements the gfx.Renderbuffer interface by wrapping a
// WebGLRenderbuffer JavaScript object.
type Renderbuffer struct {
	// Object is literally the WebGLRenderbuffer object (or nil in the case of
	// the default renderbuffer).
	Object *js.Object

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
	r.ctx.Object.Call("renderbufferStorage", r.ctx.RENDERBUFFER, r.ctx.Enums[int(internalFormat)], width, height)
}

// Delete implements the gfx.Object interface.
func (r *Renderbuffer) Delete() {
	if r.Object == nil {
		return
	}
	r.ctx.Object.Call("deleteRenderbuffer", r.Object)
	r.Object = nil
}
