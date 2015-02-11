// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

// Program implements the gfx.Program interface by wrapping a OpenGL program
// object ID.
type Program struct {
	// Object is literally the OpenGL program object ID.
	Object uint32

	ctx *Context
}
