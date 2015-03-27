// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"

// Shader implements the gfx.Shader interface by wrapping a OpenGL shader
// object ID.
type Shader struct {
	// o is literally the OpenGL shader object ID.
	o uint32

	ctx *Context
}

// Source implements the gfx.Shader interface.
func (s *Shader) Source(src string) {
	lengths := int32(len(src))
	sources := gl.Str(src + "\x00")
	gl.ShaderSource(s.o, 1, &sources, &lengths)
}

// Compile implements the gfx.Shader interface.
func (s *Shader) Compile() {
	gl.CompileShader(s.o)
}

// Delete implements the gfx.Object interface.
func (s *Shader) Delete() {
	if s.o == 0 {
		return
	}
	gl.DeleteShader(s.o)
	s.o = 0
}

// Object implements the gfx.Object interface.
func (s *Shader) Object() interface{} {
	return s.o
}
