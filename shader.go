// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// Shader represents which content and how that content is drawn to the render
// target.
type Shader interface {
	Object

	// Compile compiles the given GLSL shader source code into a binary that will
	// be linked via a Program object. It returns a boolean representing whether
	// or not compilation was successful.
	Compile(src string) bool

	// InfoLog returns the compiler information log of this shader.
	InfoLog() string
}
