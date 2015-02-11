// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Shader implements the gfx.Shader interface by wrapping a WebGLShader
// JavaScript object.
type Shader struct {
	// Object is literally the WebGLShader object.
	Object js.Object

	ctx *Context
}

// Delete implements the gfx.Object interface.
func (s *Shader) Delete() {
	if s.Object == nil {
		return
	}
	s.ctx.Object.Call("deleteShader", s.Object)
	s.Object = nil
}
