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

// UniformLocation implements the gfx.Program interface.
func (p *Program) UniformLocation(name string) gfx.UniformLocation {
	l := p.ctx.O.Call("getUniformLocation", p.o, name)
	if l == nil {
		return nil
	}
	return gfx.UniformLocation(l)
}

// Uniform1fv implements the gfx.Program interface.
func (p *Program) Uniform1fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform1fv", l, data)
}

// Uniform1iv implements the gfx.Program interface.
func (p *Program) Uniform1iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform1iv", l, data)
}

// Uniform2fv implements the gfx.Program interface.
func (p *Program) Uniform2fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform2fv", l, data)
}

// Uniform2iv implements the gfx.Program interface.
func (p *Program) Uniform2iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform2iv", l, data)
}

// Uniform3fv implements the gfx.Program interface.
func (p *Program) Uniform3fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform3fv", l, data)
}

// Uniform3iv implements the gfx.Program interface.
func (p *Program) Uniform3iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform3iv", l, data)
}

// Uniform4fv implements the gfx.Program interface.
func (p *Program) Uniform4fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform4fv", l, data)
}

// Uniform4iv implements the gfx.Program interface.
func (p *Program) Uniform4iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniform4iv", l, data)
}

// UniformMatrix2fv implements the gfx.Program interface.
func (p *Program) UniformMatrix2fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniformMatrix2fv", l, transpose, data)
}

// UniformMatrix3fv implements the gfx.Program interface.
func (p *Program) UniformMatrix3fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniformMatrix3fv", l, transpose, data)
}

// UniformMatrix4fv implements the gfx.Program interface.
func (p *Program) UniformMatrix4fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.ctx.fastUseProgram(p.o)
	p.ctx.O.Call("uniformMatrix4fv", l, transpose, data)
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
