// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"

// Buffer implements the gfx.Buffer interface by wrapping a OpenGL buffer
// object ID.
type Buffer struct {
	// Object is literally the OpenGL buffer object ID.
	Object uint32

	ctx *Context
}

// Delete implements the gfx.Object interface.
func (b *Buffer) Delete() {
	if b.Object == 0 {
		return
	}
	gl.DeleteBuffers(1, &b.Object)
	b.Object = 0
}
