// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

// Renderbuffer implements the gfx.Renderbuffer interface by wrapping a OpenGL
// renderbuffer object ID.
type Renderbuffer struct {
	// Object is literally the OpenGL renderbuffer object ID.
	Object uint32

	ctx *Context
}
