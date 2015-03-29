// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
)

// Buffer implements the gfx.Buffer interface by wrapping a OpenGL buffer
// object ID.
type Buffer struct {
	// o is literally the OpenGL buffer object ID.
	o uint32

	ctx *Context
	typ gfx.BufferType
}

// DataSize implements the gfx.Buffer interface.
func (b *Buffer) DataSize(size int, usage gfx.BufferUsage) {
	typ := b.ctx.Enums[int(b.typ)]
	b.ctx.fastBindBuffer(typ, b.o)
	gl.BufferData(typ, size, nil, b.ctx.Enums[int(usage)])
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
