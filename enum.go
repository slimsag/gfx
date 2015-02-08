// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// TextureTarget represents a single texture target.
type TextureTarget int

// RenderbufferFormat represents a renderbuffer's storage format.
type RenderbufferFormat int

// FramebufferAttachment represents a framebuffer attachment point. It must be
// one of the predefined constants.
type FramebufferAttachment int

// BufferUsage is a hint that describes how a Buffer's data might be used. It
// must be one of the predefined constants StaticDraw, DynamicDraw, or
// StreamDraw.
type BufferUsage int

// Feature is a single OpenGL feature that can be enabled or disabled. It must
// be one of the predefined constants.
type Feature int

// Orientation represents a single orientation.
type Orientation int

// Facet represents a single facet.
type Facet int

const (
	// TextureTarget enumerations.
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

	// RenderbufferFormat enumerations.
	zeroRenderbufferFormat RenderbufferFormat = iota

	// RGBA4 is a 4-bits per channel renderbuffer storage format.
	RGBA4

	// RGB565 is a renderbuffer storage format with 5 bits red, 6 bits green,
	// and 5 bits blue, respectively.
	RGB565

	// RGB5A1 is a renderbuffer storage format with 5 bits for RGB and 1 bit
	// for alpha.
	RGB5A1

	// DepthComponent16 is a renderbuffer storage format with
	DepthComponent16

	// FramebufferAttachment enumerations.
	zeroFramebufferAttachment FramebufferAttachment = iota

	// ColorAttachment0 is a framebuffer attachment point for the color buffer.
	ColorAttachment0

	// ColorAttachment0 is a framebuffer attachment point for the depth buffer.
	DepthAttachment

	// ColorAttachment0 is a framebuffer attachment point for the stencil
	// buffer.
	StencilAttachment

	// DepthStencilAttachment is a framebuffer attachment point for the depth
	// and stencil buffer.
	DepthStencilAttachment

	// BufferUsage enumerations.
	zeroBufferUsage BufferUsage = iota

	// StaticDraw is a buffer usage hint where the data is static and generally
	// does not change.
	StaticDraw

	// DynamicDraw is a buffer usage hint where the data changes moderately
	// often.
	DynamicDraw

	// StreamDraw is a buffer usage hint where the data is likely to be used
	// just once and then changed immedietely thereafter.
	StreamDraw

	// Feature enumerations.
	zeroFeature Feature = iota

	// Blend is a feature that blends computed fragment color values with color
	// buffer values.
	Blend

	// DepthTest is a feature that enables or disables testing against depth
	// buffer values.
	DepthTest

	// CullFace is a feature that enables or disables the culling of polygon
	// faces.
	CullFace

	// PolygonOffsetFill adds an offset to the depth values of a polygon's
	// fragments.
	PolygonOffsetFill

	// ScissorTest abandons fragments outside the scissor rectangle.
	ScissorTest

	// Orientation enumerations.
	zeroOrientation Orientation = iota

	// CCW is a orientation for a counterclockwise winding. It is the initial
	// orientation.
	CCW

	// CW is a orientation for clockwise winding.
	CW

	// Facet enumerations.
	zeroFacet Facet = iota

	// Front is a facet for representing front-facing polygons.
	Front

	// Back is a facet for representing back-facing polygons.
	Back

	// FrontAndBack is a facet for representing both front and back facing
	// polygons.
	FrontAndBack

	// EnumMax is the maximum bound for enumerations. It may change in minor
	// releases and is the maximum value for any enumeration. I.e. enumerations
	// are integers in the range of [0 - EnumMax].
	EnumMax int = iota
)
