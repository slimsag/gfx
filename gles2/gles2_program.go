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

// AttachShader implements the gfx.Program interface.
func (p *Program) AttachShader(s gfx.Shader) {
	gl.AttachShader(p.o, s.Object().(uint32))
}

// Link implements the gfx.Program interface.
func (p *Program) Link() {
	gl.LinkProgram(p.o)
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
