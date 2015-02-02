// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// TextureTarget represents a single texture target.
type TextureTarget int

const (
	zeroTextureTarget TextureTarget = iota

	// Texture2D is a 2D image.
	Texture2D

	// TextureCubeMapPositiveX is a image for the positive X face of the cube.
	TextureCubeMapPositiveX

	// TextureCubeMapNegativeX is a image for the negative X face of the cube.
	TextureCubeMapNegativeX

	// TextureCubeMapPositiveY is a image for the positive Y face of the cube.
	TextureCubeMapPositiveY

	// TextureCubeMapNegativeY is a image for the negative Y face of the cube.
	TextureCubeMapNegativeY

	// TextureCubeMapPositiveZ is a image for the positive Z face of the cube.
	TextureCubeMapPositiveZ

	// TextureCubeMapNegativeZ is a image for the negative Z face of the cube.
	TextureCubeMapNegativeZ
)

// Texture represents a single texture. It is used for images and cube maps
// when rendering shapes.
type Texture interface {
}
