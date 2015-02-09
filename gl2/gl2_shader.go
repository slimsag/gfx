// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

// Shader implements the gfx.Shader interface by wrapping a OpenGL shader
// object ID.
type Shader struct {
	// Object is literally the OpenGL shader object ID.
	Object uint32

	ctx *Context
}
