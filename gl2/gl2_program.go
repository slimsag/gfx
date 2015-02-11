// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import "github.com/slimsag/gfx/internal/gl/2.0/gl"

// Program implements the gfx.Program interface by wrapping a OpenGL program
// object ID.
type Program struct {
	// Object is literally the OpenGL program object ID.
	Object uint32

	ctx *Context
}

// Delete implements the gfx.Object interface.
func (p *Program) Delete() {
	if p.Object == 0 {
		return
	}
	gl.DeleteProgram(p.Object)
	p.Object = 0
}
