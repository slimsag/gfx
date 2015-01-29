// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// BufferUsage is a hint that describes how a Buffer's data might be used. It
// must be one of the predefined constants StaticDraw, DynamicDraw, or
// StreamDraw.
type BufferUsage int

const (
	// StaticDraw is a buffer usage hint where the data is static and generally
	// does not change.
	StaticDraw BufferUsage = iota

	// DynamicDraw is a buffer usage hint where the data changes moderately
	// often.
	DynamicDraw

	// StreamDraw is a buffer usage hint where the data is likely to be used
	// just once and then changed immedietely thereafter.
	StreamDraw
)

// Buffer is a buffer object that contains data such as vertices or colors.
type Buffer interface {
	// DataSize prepares this buffer with size bytes of storage, initialized to
	// zero.
	//
	// The usage hint is only a performance hint, it has no effect on the
	// actual usage of the buffer.
	//
	// This function may generate an ErrOutOfMemory error, see Context.Error
	// for more details.
	DataSize(size int, usage BufferUsage)

	// Data prepares this buffer with the given data.
	//
	// The usage hint is only a performance hint, it has no effect on the
	// actual usage of the buffer.
	//
	// This function may generate an ErrOutOfMemory error, see Context.Error
	// for more details.
	//
	// TODO(slimsag): typeof(data) == ArrayBuffer
	Data(data interface{}, usage BufferUsage)

	// SubData updates a subarea of the data buffer with the given data,
	// starting at the offset.
	//
	// This function may generate an ErrInvalidValue error if the new data
	// would write past the end of the buffer.
	//
	// TODO(slimsag): typeof(data) == ArrayBuffer
	SubData(offset uint, data interface{})
}
