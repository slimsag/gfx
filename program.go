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
}
