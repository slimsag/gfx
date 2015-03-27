// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Buffer implements the gfx.Buffer interface by wrapping a WebGLBuffer
// JavaScript object.
type Buffer struct {
	// o is literally the WebGLBuffer object.
	o *js.Object

	ctx *Context
}

// Delete implements the gfx.Object interface.
func (b *Buffer) Delete() {
	if b.o == nil {
		return
	}
	b.ctx.O.Call("deleteBuffer", b.o)
	b.o = nil
}

// Object implements the gfx.Object interface.
func (b *Buffer) Object() interface{} {
	return b.o
}
