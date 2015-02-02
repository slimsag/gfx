// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import "github.com/slimsag/gfx/internal/gl/2.0/gl"

// Renderbuffer implements the gfx.Renderbuffer interface by wrapping a OpenGL
// renderbuffer object ID.
type Renderbuffer struct {
	// Object is literally the OpenGL renderbuffer object ID.
	Object uint32

	ctx *Context
}

// Delete implements the gfx.Renderbuffer interface.
func (r *Renderbuffer) Delete() {
	gl.DeleteRenderbuffers(1, &r.Object)
}
