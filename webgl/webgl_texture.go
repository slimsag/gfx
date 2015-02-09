// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/gopherjs/gopherjs/js"

// Texture implements the gfx.Texture interface by wrapping a WebGLTexture
// JavaScript object.
type Texture struct {
	// Object is literally the WebGLTexture object.
	Object js.Object

	ctx *Context
}
