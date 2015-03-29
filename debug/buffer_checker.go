// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// bufferChecker is like the checker type, but for a gfx.Buffer. It implicitly
// invokes the Check method of the underlying context after each function call
// is made.
type bufferChecker struct {
	b   gfx.Buffer
	ctx gfx.Context
}

// DataSize implements the gfx.Buffer interface.
func (b *bufferChecker) DataSize(size int, usage gfx.BufferUsage) {
	b.b.DataSize(size, usage)
	b.ctx.Check()
}

// DataInt8 implements the gfx.Buffer interface.
func (b *bufferChecker) DataInt8(data []int8, usage gfx.BufferUsage) {
	b.b.DataInt8(data, usage)
	b.ctx.Check()
}

// DataUint8 implements the gfx.Buffer interface.
func (b *bufferChecker) DataUint8(data []uint8, usage gfx.BufferUsage) {
	b.b.DataUint8(data, usage)
	b.ctx.Check()
}

// DataInt16 implements the gfx.Buffer interface.
func (b *bufferChecker) DataInt16(data []int16, usage gfx.BufferUsage) {
	b.b.DataInt16(data, usage)
	b.ctx.Check()
}

// DataUint16 implements the gfx.Buffer interface.
func (b *bufferChecker) DataUint16(data []uint16, usage gfx.BufferUsage) {
	b.b.DataUint16(data, usage)
	b.ctx.Check()
}

// DataInt32 implements the gfx.Buffer interface.
func (b *bufferChecker) DataInt32(data []int32, usage gfx.BufferUsage) {
	b.b.DataInt32(data, usage)
	b.ctx.Check()
}

// DataUint32 implements the gfx.Buffer interface.
func (b *bufferChecker) DataUint32(data []uint32, usage gfx.BufferUsage) {
	b.b.DataUint32(data, usage)
	b.ctx.Check()
}

// DataFloat32 implements the gfx.Buffer interface.
func (b *bufferChecker) DataFloat32(data []float32, usage gfx.BufferUsage) {
	b.b.DataFloat32(data, usage)
	b.ctx.Check()
}

// DataFloat64 implements the gfx.Buffer interface.
func (b *bufferChecker) DataFloat64(data []float64, usage gfx.BufferUsage) {
	b.b.DataFloat64(data, usage)
	b.ctx.Check()
}

// SubDataInt8 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataInt8(offset int, data []int8) {
	b.b.SubDataInt8(offset, data)
	b.ctx.Check()
}

// SubDataUint8 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataUint8(offset int, data []uint8) {
	b.b.SubDataUint8(offset, data)
	b.ctx.Check()
}

// SubDataInt16 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataInt16(offset int, data []int16) {
	b.b.SubDataInt16(offset, data)
	b.ctx.Check()
}

// SubDataUint16 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataUint16(offset int, data []uint16) {
	b.b.SubDataUint16(offset, data)
	b.ctx.Check()
}

// SubDataInt32 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataInt32(offset int, data []int32) {
	b.b.SubDataInt32(offset, data)
	b.ctx.Check()
}

// SubDataUint32 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataUint32(offset int, data []uint32) {
	b.b.SubDataUint32(offset, data)
	b.ctx.Check()
}

// SubDataFloat32 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataFloat32(offset int, data []float32) {
	b.b.SubDataFloat32(offset, data)
	b.ctx.Check()
}

// SubDataFloat64 implements the gfx.Buffer interface.
func (b *bufferChecker) SubDataFloat64(offset int, data []float64) {
	b.b.SubDataFloat64(offset, data)
	b.ctx.Check()
}

// Draw implements the gfx.Buffer interface.
func (b *bufferChecker) Draw(p gfx.Primitive, first, count int) {
	b.b.Draw(p, first, count)
	b.ctx.Check()
}

// Delete implements the gfx.Object interface.
func (b *bufferChecker) Delete() {
	b.b.Delete()
	b.ctx.Check()
}

// Object implements the gfx.Object interface.
func (b *bufferChecker) Object() interface{} {
	return b.b.Object()
}
