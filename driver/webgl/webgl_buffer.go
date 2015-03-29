// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Buffer implements the gfx.Buffer interface by wrapping a WebGLBuffer
// JavaScript object.
type Buffer struct {
	// o is literally the WebGLBuffer object.
	o *js.Object

	ctx *Context
	typ gfx.BufferType
}

// DataSize implements the gfx.Buffer interface.
func (b *Buffer) DataSize(size int, usage gfx.BufferUsage) {
	typ := b.ctx.Enums[int(b.typ)]
	b.ctx.fastBindBuffer(typ, b.o)
	b.ctx.O.Call("bufferData", typ, size, b.ctx.Enums[int(usage)])
}

func (b *Buffer) data(x interface{}, usage gfx.BufferUsage) {
	typ := b.ctx.Enums[int(b.typ)]
	b.ctx.fastBindBuffer(typ, b.o)
	b.ctx.O.Call("bufferData", typ, x, b.ctx.Enums[int(usage)])
}

// DataInt8 implements the gfx.Buffer interface.
func (b *Buffer) DataInt8(data []int8, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataUint8 implements the gfx.Buffer interface.
func (b *Buffer) DataUint8(data []uint8, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataInt16 implements the gfx.Buffer interface.
func (b *Buffer) DataInt16(data []int16, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataUint16 implements the gfx.Buffer interface.
func (b *Buffer) DataUint16(data []uint16, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataInt32 implements the gfx.Buffer interface.
func (b *Buffer) DataInt32(data []int32, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataUint32 implements the gfx.Buffer interface.
func (b *Buffer) DataUint32(data []uint32, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataFloat32 implements the gfx.Buffer interface.
func (b *Buffer) DataFloat32(data []float32, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// DataFloat64 implements the gfx.Buffer interface.
func (b *Buffer) DataFloat64(data []float64, usage gfx.BufferUsage) {
	b.data(data, usage)
}

// Delete implements the gfx.Object interface.
func (b *Buffer) Delete() {
	if b.o == nil {
		return
	}
	b.ctx.O.Call("deleteBuffer", b.o)
	b.o = nil
}

// Object implements the gfx.Object interface.
func (b *Buffer) Object() interface{} {
	return b.o
}
