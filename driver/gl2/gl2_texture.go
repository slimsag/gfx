// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
)

// Texture implements the gfx.Texture interface by wrapping a OpenGL
// texture object ID.
type Texture struct {
	// o is literally the OpenGL texture object ID.
	o uint32

	ctx *Context
	typ gfx.TextureType
}

// Type implements the gfx.Texture interface.
func (t *Texture) Type() gfx.TextureType {
	return t.typ
}

// Delete implements the gfx.Object interface.
func (t *Texture) Delete() {
	if t.o == 0 {
		return
	}
	gl.DeleteTextures(1, &t.o)
	t.o = 0
}

// Object implements the gfx.Object interface.
func (t *Texture) Object() interface{} {
	return t.o
}
