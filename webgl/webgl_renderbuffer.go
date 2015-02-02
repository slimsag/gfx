// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Renderbuffer implements the gfx.Renderbuffer interface by wrapping a
// WebGLRenderbuffer JavaScript object.
type Renderbuffer struct {
	// Object is literally the WebGLRenderbuffer object (or nil in the case of
	// the default renderbuffer).
	Object js.Object

	ctx *Context
}

// Delete implements the gfx.Renderbuffer interface.
func (r *Renderbuffer) Delete() {
	r.ctx.Object.Call("deleteRenderbuffer", r.Object)
}
