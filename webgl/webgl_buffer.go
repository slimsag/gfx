// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Buffer implements the gfx.Buffer interface by wrapping a WebGLBuffer
// JavaScript object.
type Buffer struct {
	// Object is literally the WebGLBuffer object.
	Object js.Object

	ctx *Context
}
