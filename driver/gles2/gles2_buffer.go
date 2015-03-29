// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"unsafe"

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

func (b *Buffer) data(size int, ptr unsafe.Pointer, usage gfx.BufferUsage) {
	typ := b.ctx.Enums[int(b.typ)]
	b.ctx.fastBindBuffer(typ, b.o)
	gl.BufferData(typ, size, ptr, b.ctx.Enums[int(usage)])
}

// DataInt8 implements the gfx.Buffer interface.
func (b *Buffer) DataInt8(data []int8, usage gfx.BufferUsage) {
	b.data(len(data), unsafe.Pointer(&data[0]), usage)
}

// DataUint8 implements the gfx.Buffer interface.
func (b *Buffer) DataUint8(data []uint8, usage gfx.BufferUsage) {
	b.data(len(data), unsafe.Pointer(&data[0]), usage)
}

// DataInt16 implements the gfx.Buffer interface.
func (b *Buffer) DataInt16(data []int16, usage gfx.BufferUsage) {
	b.data(len(data)*2, unsafe.Pointer(&data[0]), usage)
}

// DataUint16 implements the gfx.Buffer interface.
func (b *Buffer) DataUint16(data []uint16, usage gfx.BufferUsage) {
	b.data(len(data)*2, unsafe.Pointer(&data[0]), usage)
}

// DataInt32 implements the gfx.Buffer interface.
func (b *Buffer) DataInt32(data []int32, usage gfx.BufferUsage) {
	b.data(len(data)*4, unsafe.Pointer(&data[0]), usage)
}

// DataUint32 implements the gfx.Buffer interface.
func (b *Buffer) DataUint32(data []uint32, usage gfx.BufferUsage) {
	b.data(len(data)*4, unsafe.Pointer(&data[0]), usage)
}

// DataFloat32 implements the gfx.Buffer interface.
func (b *Buffer) DataFloat32(data []float32, usage gfx.BufferUsage) {
	b.data(len(data)*4, unsafe.Pointer(&data[0]), usage)
}

// DataFloat64 implements the gfx.Buffer interface.
func (b *Buffer) DataFloat64(data []float64, usage gfx.BufferUsage) {
	b.data(len(data)*8, unsafe.Pointer(&data[0]), usage)
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
