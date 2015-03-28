// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Texture implements the gfx.Texture interface by wrapping a WebGLTexture
// JavaScript object.
type Texture struct {
	// o is literally the WebGLTexture object.
	o *js.Object

	ctx *Context
	typ gfx.TextureType
}

// Delete implements the gfx.Object interface.
func (t *Texture) Delete() {
	if t.o == nil {
		return
	}
	t.ctx.O.Call("deleteTexture", t.o)
	t.o = nil
}

// Object implements the gfx.Object interface.
func (t *Texture) Object() interface{} {
	return t.o
}
