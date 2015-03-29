// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
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
