// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
)

func convertTextureTarget(t gfx.TextureTarget) uint32 {
	switch t {
	case gfx.Texture2D:
		return gl.TEXTURE_2D
	case gfx.TextureCubeMapPositiveX:
		return gl.TEXTURE_CUBE_MAP_POSITIVE_X
	case gfx.TextureCubeMapNegativeX:
		return gl.TEXTURE_CUBE_MAP_NEGATIVE_X
	case gfx.TextureCubeMapPositiveY:
		return gl.TEXTURE_CUBE_MAP_POSITIVE_Y
	case gfx.TextureCubeMapNegativeY:
		return gl.TEXTURE_CUBE_MAP_NEGATIVE_Y
	case gfx.TextureCubeMapPositiveZ:
		return gl.TEXTURE_CUBE_MAP_POSITIVE_Z
	case gfx.TextureCubeMapNegativeZ:
		return gl.TEXTURE_CUBE_MAP_NEGATIVE_Z
	default:
		panic("invalid texture target parameter")
	}
}

// Texture implements the gfx.Framebuffer interface by wrapping a OpenGL
// texture object ID.
type Texture struct {
	// Object is literally the OpenGL texture object ID.
	Object uint32

	ctx *Context
}
