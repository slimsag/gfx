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

func (b *Buffer) subData(offset, size int, ptr unsafe.Pointer) {
	typ := b.ctx.Enums[int(b.typ)]
	b.ctx.fastBindBuffer(typ, b.o)
	gl.BufferSubData(typ, offset, size, ptr)
}

// SubDataInt8 implements the gfx.Buffer interface.
func (b *Buffer) SubDataInt8(offset int, data []int8) {
	b.subData(offset, len(data), unsafe.Pointer(&data[0]))
}

// SubDataUint8 implements the gfx.Buffer interface.
func (b *Buffer) SubDataUint8(offset int, data []uint8) {
	b.subData(offset, len(data), unsafe.Pointer(&data[0]))
}

// SubDataInt16 implements the gfx.Buffer interface.
func (b *Buffer) SubDataInt16(offset int, data []int16) {
	b.subData(offset*2, len(data)*2, unsafe.Pointer(&data[0]))
}

// SubDataUint16 implements the gfx.Buffer interface.
func (b *Buffer) SubDataUint16(offset int, data []uint16) {
	b.subData(offset*2, len(data)*2, unsafe.Pointer(&data[0]))
}

// SubDataInt32 implements the gfx.Buffer interface.
func (b *Buffer) SubDataInt32(offset int, data []int32) {
	b.subData(offset*4, len(data)*4, unsafe.Pointer(&data[0]))
}

// SubDataUint32 implements the gfx.Buffer interface.
func (b *Buffer) SubDataUint32(offset int, data []uint32) {
	b.subData(offset*4, len(data)*4, unsafe.Pointer(&data[0]))
}

// SubDataFloat32 implements the gfx.Buffer interface.
func (b *Buffer) SubDataFloat32(offset int, data []float32) {
	b.subData(offset*4, len(data)*4, unsafe.Pointer(&data[0]))
}

// SubDataFloat64 implements the gfx.Buffer interface.
func (b *Buffer) SubDataFloat64(offset int, data []float64) {
	b.subData(offset*8, len(data)*8, unsafe.Pointer(&data[0]))
}

// Draw implements the gfx.Buffer interface.
func (b *Buffer) Draw(p gfx.Primitive, first, count int) {
	b.ctx.fastBindBuffer(b.ctx.Enums[int(b.typ)], b.o)
	if b.typ == gfx.ArrayBuffer {
		gl.DrawArrays(b.ctx.Enums[int(p)], int32(first), int32(count))
	} else {
		offset := uint32(first * 4)
		gl.DrawElements(b.ctx.Enums[int(p)], int32(count), gl.UNSIGNED_SHORT, unsafe.Pointer(uintptr(offset)))
	}
}

// VertexAttribPointer implements the gfx.Buffer interface.
func (b *Buffer) VertexAttribPointer(l gfx.AttribLocation, size int, normalized bool, stride, offset int) {
	b.ctx.fastBindBuffer(b.ctx.Enums[int(b.typ)], b.o)
	gl.VertexAttribPointer(uint32(l.(int32)), int32(size), gl.FLOAT, normalized, int32(stride), unsafe.Pointer(uintptr(offset)))
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
