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

// Compile implements the gfx.Shader interface.
func (s *Shader) Compile(src string) bool {
	lengths := int32(len(src))
	sources := gl.Str(src + "\x00")
	gl.ShaderSource(s.o, 1, &sources, &lengths)

	gl.CompileShader(s.o)
	var success int32
	gl.GetShaderiv(s.o, gl.COMPILE_STATUS, &success)
	return success == 1
}

// InfoLog implements the gfx.Shader interface.
func (s *Shader) InfoLog() string {
	var length int32
	gl.GetShaderiv(s.o, gl.INFO_LOG_LENGTH, &length)
	log := make([]byte, length)
	gl.GetShaderInfoLog(s.o, length, nil, &log[0])
	return string(log)
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
