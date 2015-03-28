// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// Texture represents a single texture. It is used for images and cube maps
// when rendering shapes.
type Texture interface {
	Object

	// Type returns the type of this texture, either TextureType2D or
	// TextureTypeCubeMap.
	Type() TextureType
}
