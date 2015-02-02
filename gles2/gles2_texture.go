// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

// Texture implements the gfx.Framebuffer interface by wrapping a OpenGL
// texture object ID.
type Texture struct {
	// Object is literally the OpenGL texture object ID.
	Object uint32

	ctx *Context
}
