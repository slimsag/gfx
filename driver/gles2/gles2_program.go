// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
)

// Program implements the gfx.Program interface by wrapping a OpenGL program
// object ID.
type Program struct {
	// o is literally the OpenGL program object ID.
	o uint32

	ctx *Context
}

// Link implements the gfx.Program interface.
func (p *Program) Link(vert, frag gfx.Shader) bool {
	gl.AttachShader(p.o, vert.Object().(uint32))
	gl.AttachShader(p.o, frag.Object().(uint32))
	gl.LinkProgram(p.o)

	var status int32
	gl.GetProgramiv(p.o, gl.LINK_STATUS, &status)
	return status == 1
}

// InfoLog implements the gfx.Program interface.
func (p *Program) InfoLog() string {
	var length int32
	gl.GetShaderiv(p.o, gl.INFO_LOG_LENGTH, &length)
	log := make([]byte, length)
	gl.GetProgramInfoLog(p.o, length, nil, &log[0])
	return string(log)
}

// AttribLocation implements the gfx.Program interface.
func (p *Program) AttribLocation(name string) gfx.AttribLocation {
	l := gl.GetAttribLocation(p.o, gl.Str(name+"\x00"))
	if l == -1 {
		return nil
	}
	return gfx.AttribLocation(l)
}

// UniformLocation implements the gfx.Program interface.
func (p *Program) UniformLocation(name string) gfx.UniformLocation {
	l := gl.GetUniformLocation(p.o, gl.Str(name+"\x00"))
	if l == -1 {
		return nil
	}
	return gfx.UniformLocation(l)
}

// Uniform1fv implements the gfx.Program interface.
func (p *Program) Uniform1fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform1fv(l.(int32), int32(len(data)), &data[0])
}

// Uniform1iv implements the gfx.Program interface.
func (p *Program) Uniform1iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform1iv(l.(int32), int32(len(data)), &data[0])
}

// Uniform2fv implements the gfx.Program interface.
func (p *Program) Uniform2fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform2fv(l.(int32), int32(len(data)/2), &data[0])
}

// Uniform2iv implements the gfx.Program interface.
func (p *Program) Uniform2iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform2iv(l.(int32), int32(len(data)/2), &data[0])
}

// Uniform3fv implements the gfx.Program interface.
func (p *Program) Uniform3fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform3fv(l.(int32), int32(len(data)/3), &data[0])
}

// Uniform3iv implements the gfx.Program interface.
func (p *Program) Uniform3iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform3iv(l.(int32), int32(len(data)/3), &data[0])
}

// Uniform4fv implements the gfx.Program interface.
func (p *Program) Uniform4fv(l gfx.UniformLocation, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform4fv(l.(int32), int32(len(data)/4), &data[0])
}

// Uniform4iv implements the gfx.Program interface.
func (p *Program) Uniform4iv(l gfx.UniformLocation, data []int32) {
	p.ctx.fastUseProgram(p.o)
	gl.Uniform4iv(l.(int32), int32(len(data)/4), &data[0])
}

// UniformMatrix2fv implements the gfx.Program interface.
func (p *Program) UniformMatrix2fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.UniformMatrix2fv(l.(int32), int32(len(data)/(2*2)), transpose, &data[0])
}

// UniformMatrix3fv implements the gfx.Program interface.
func (p *Program) UniformMatrix3fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.UniformMatrix3fv(l.(int32), int32(len(data)/(3*3)), transpose, &data[0])
}

// UniformMatrix4fv implements the gfx.Program interface.
func (p *Program) UniformMatrix4fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.ctx.fastUseProgram(p.o)
	gl.UniformMatrix4fv(l.(int32), int32(len(data)/(4*4)), transpose, &data[0])
}

// Delete implements the gfx.Object interface.
func (p *Program) Delete() {
	if p.o == 0 {
		return
	}
	gl.DeleteProgram(p.o)
	p.o = 0
}

// Object implements the gfx.Object interface.
func (p *Program) Object() interface{} {
	return p.o
}
