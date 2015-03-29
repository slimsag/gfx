// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// Buffer is a buffer object that contains data such as vertices or colors.
type Buffer interface {
	Object

	// DataSize prepares this buffer with size bytes of storage, initialized to
	// zero.
	//
	// The usage hint is only a performance hint, it has no effect on the
	// actual usage of the buffer.
	//
	// Calling this function may generate a OutOfMemory panic at Context.Check
	// time.
	DataSize(size int, usage BufferUsage)

	// Data* prepares this buffer with the given data.
	//
	// The usage hint is only a performance hint, it has no effect on the
	// actual usage of the buffer.
	//
	// Calling this function may generate a OutOfMemory panic at Context.Check
	// time.
	DataInt8(data []int8, usage BufferUsage)
	DataUint8(data []uint8, usage BufferUsage)
	DataInt16(data []int16, usage BufferUsage)
	DataUint16(data []uint16, usage BufferUsage)
	DataInt32(data []int32, usage BufferUsage)
	DataUint32(data []uint32, usage BufferUsage)
	DataFloat32(data []float32, usage BufferUsage)
	DataFloat64(data []float64, usage BufferUsage)

	// SubData updates a subarea of the data buffer with the given data,
	// starting at the offset in elements (not bytes).
	//
	// This function will generate an InvalidValue panic at Context.Check time
	// if the new data would write past the end of the buffer.
	//
	// Calling this function may generate a OutOfMemory panic at Context.Check
	// time.
	SubDataInt8(offset int, data []int8)
	SubDataUint8(offset int, data []uint8)
	SubDataInt16(offset int, data []int16)
	SubDataUint16(offset int, data []uint16)
	SubDataInt32(offset int, data []int32)
	SubDataUint32(offset int, data []uint32)
	SubDataFloat32(offset int, data []float32)
	SubDataFloat64(offset int, data []float64)

	// Draw draws the contents of this buffer as the given type of primitive
	// object (e.g. Triangles).
	//
	// The first parameter is the first element in the array to start drawing data
	// at, and count is the number of elements to draw (each measured in single
	// vertex units, e.g. a trangle is 3).
	Draw(p Primitive, first, count int)
}
