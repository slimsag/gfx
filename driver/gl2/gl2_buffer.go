// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
)

// Buffer implements the gfx.Buffer interface by wrapping a OpenGL buffer
// object ID.
type Buffer struct {
	// o is literally the OpenGL buffer object ID.
	o uint32

	ctx *Context
	typ gfx.BufferType
}

// Delete implements the gfx.Object interface.
func (b *Buffer) Delete() {
	if b.o == 0 {
		return
	}
	gl.DeleteBuffers(1, &b.o)
	b.o = 0
}

// Object implements the gfx.Object interface.
func (b *Buffer) Object() interface{} {
	return b.o
}
