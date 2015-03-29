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

// Link implements the gfx.Program interface.
func (p *Program) Link(vert, frag gfx.Shader) bool {
	p.ctx.O.Call("attachShader", p.o, vert.Object().(*js.Object))
	p.ctx.O.Call("attachShader", p.o, frag.Object().(*js.Object))
	p.ctx.O.Call("linkProgram", p.o)
	return p.ctx.O.Call("getProgramParameter", p.o, p.ctx.LINK_STATUS).Bool()
}

// InfoLog implements the gfx.Program interface.
func (p *Program) InfoLog() string {
	return p.ctx.O.Call("getProgramInfoLog", p.o).String()
}

// AttribLocation implements the gfx.Program interface.
func (p *Program) AttribLocation(name string) gfx.AttribLocation {
	l := p.ctx.O.Call("getAttribLocation", p.o, name).Int()
	if l == -1 {
		return nil
	}
	return gfx.AttribLocation(l)
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
