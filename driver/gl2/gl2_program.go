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
