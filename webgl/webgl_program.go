// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Program implements the gfx.Program interface by wrapping a WebGLProgram
// JavaScript object.
type Program struct {
	// Object is literally the WebGLProgram object.
	Object js.Object

	ctx *Context
}

// Delete implements the gfx.Object interface.
func (p *Program) Delete() {
	if p.Object == nil {
		return
	}
	p.ctx.Object.Call("deleteProgram", p.Object)
	p.Object = nil
}
