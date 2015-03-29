// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// AttribLocation is an opaque type representing the location of a attribute
// variable in a GLSL program.
type AttribLocation interface{}

// UniformLocation is an opaque type representing the location of a uniform
// variable in a GLSL program.
type UniformLocation interface{}

// Program represents the programmable OpenGL pipeline and associated shader
// programs.
type Program interface {
	Object

	// Link links the given vertex and fragment shaders into a program so that it
	// can be used by the GPU. It returns whether or not linking the shaders into
	// a program was successful or not.
	Link(vert, frag Shader) bool

	// InfoLog returns the linker information log of this program.
	InfoLog() string

	// AttribLocation returns the location in this program of the named attribute
	// variable, or nil if there is no such variable.
	AttribLocation(name string) AttribLocation

	// UniformLocation returns the location in this program of the named uniform
	// variable, or nil if there is no such variable.
	UniformLocation(name string) UniformLocation

	// Uniform{1,2,3,4}{f,i}v sets values for a N component floating-point or
	// integer vector into a uniform location as a vector or vector array.
	Uniform1fv(l UniformLocation, data []float32)
	Uniform1iv(l UniformLocation, data []int32)
	Uniform2fv(l UniformLocation, data []float32)
	Uniform2iv(l UniformLocation, data []int32)
	Uniform3fv(l UniformLocation, data []float32)
	Uniform3iv(l UniformLocation, data []int32)
	Uniform4fv(l UniformLocation, data []float32)
	Uniform4iv(l UniformLocation, data []int32)

	// UniformMatrix{2,3,4}{f,i}v sets values for a 4x4 floating point or integer
	// vector matrix into a uniform location as a matrix or a matrix array.
	UniformMatrix2fv(l UniformLocation, transpose bool, data []float32)
	UniformMatrix3fv(l UniformLocation, transpose bool, data []float32)
	UniformMatrix4fv(l UniformLocation, transpose bool, data []float32)
}
