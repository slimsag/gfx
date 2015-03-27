// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Shader implements the gfx.Shader interface by wrapping a WebGLShader
// JavaScript object.
type Shader struct {
	// o is literally the WebGLShader object.
	o *js.Object

	ctx *Context
}

// Source implements the gfx.Shader interface.
func (s *Shader) Source(src string) {
	s.ctx.O.Call("shaderSource", s.o, src)
}

// Compile implements the gfx.Shader interface.
func (s *Shader) Compile() {
	s.ctx.O.Call("compileShader", s.o)
}

// Delete implements the gfx.Object interface.
func (s *Shader) Delete() {
	if s.o == nil {
		return
	}
	s.ctx.O.Call("deleteShader", s.o)
	s.o = nil
}

// Object implements the gfx.Object interface.
func (s *Shader) Object() interface{} {
	return s.o
}