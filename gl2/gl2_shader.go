// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import "github.com/slimsag/gfx/internal/gl/2.0/gl"

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
	sources := gl.Str(src)
	gl.ShaderSource(s.o, 1, &sources, &lengths)
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
