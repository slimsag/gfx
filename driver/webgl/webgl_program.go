// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Program implements the gfx.Program interface by wrapping a WebGLProgram
// JavaScript object.
type Program struct {
	// o is literally the WebGLProgram object.
	o *js.Object

	ctx *Context
}

// AttachShader implements the gfx.Program interface.
func (p *Program) AttachShader(s gfx.Shader) {
	p.ctx.O.Call("attachShader", p.o, s.Object().(*js.Object))
}

// Link implements the gfx.Program interface.
func (p *Program) Link() {
	p.ctx.O.Call("linkProgram", p.o)
}

// Delete implements the gfx.Object interface.
func (p *Program) Delete() {
	if p.o == nil {
		return
	}
	p.ctx.O.Call("deleteProgram", p.o)
	p.o = nil
}

// Object implements the gfx.Object interface.
func (p *Program) Object() interface{} {
	return p.o
}