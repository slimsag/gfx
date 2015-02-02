// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

func (c *Context) convertTextureTarget(t gfx.TextureTarget) int {
	// Convert the target parameter.
	switch t {
	case gfx.Texture2D:
		return c.TEXTURE_2D
	case gfx.TextureCubeMapPositiveX:
		return c.TEXTURE_CUBE_MAP_POSITIVE_X
	case gfx.TextureCubeMapNegativeX:
		return c.TEXTURE_CUBE_MAP_NEGATIVE_X
	case gfx.TextureCubeMapPositiveY:
		return c.TEXTURE_CUBE_MAP_POSITIVE_Y
	case gfx.TextureCubeMapNegativeY:
		return c.TEXTURE_CUBE_MAP_NEGATIVE_Y
	case gfx.TextureCubeMapPositiveZ:
		return c.TEXTURE_CUBE_MAP_POSITIVE_Z
	case gfx.TextureCubeMapNegativeZ:
		return c.TEXTURE_CUBE_MAP_NEGATIVE_Z
	default:
		panic("invalid texture target parameter")
	}
}

// Texture implements the gfx.Texture interface by wrapping a WebGLTexture
// JavaScript object.
type Texture struct {
	Object js.Object
}
