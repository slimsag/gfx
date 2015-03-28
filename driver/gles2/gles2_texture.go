// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
)

// Texture implements the gfx.Texture interface by wrapping a OpenGL
// texture object ID.
type Texture struct {
	// o is literally the OpenGL texture object ID.
	o uint32

	ctx *Context
	typ gfx.TextureType
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
