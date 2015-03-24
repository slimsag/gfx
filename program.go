// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// Program represents the programmable OpenGL pipeline and associated shader
// programs.
type Program interface {
	Object

	// AttachShader attaches the given shader object to this program.
	//
	// If the shader is already attached to a program, or another shader of the
	// same type is already attached to this program, a InvalidOperation panic
	// will be generated at Context.Check time.
	AttachShader(s Shader)
}
