// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// ContextStateValue represents a single value as part of a context's state,
// for example a boolean representing whether multisampling is enabled or not.
//
// The underlying type is platform-specific, do not access it directly or make
// assumptions about it.
type ContextStateValue interface{}

// ContextState solely represents a context's unique state. Any values not
// explicitly specified are assumed to be their defaults.
//
// The underlying type is platform-specific, do not access it directly or make
// assumptions about it.
type ContextState interface{}

// ContextStateProvider provides access to a graphics context's state.
type ContextStateProvider interface {
	// NewState returns a new context state for the given values.
	NewState(values ...ContextStateValue) ContextState

	// Load loads the given context state, replacing the previous one. If
	// s == nil then the default state is loaded.
	Load(s ContextState)

	// BlendColor specifies the blend color used to calculate source and
	// destination blending.
	BlendColor(r, g, b, a float32) ContextStateValue

	// BlendEquation sets the equation used to blend RGB and Alpha values of an
	// incoming source fragment with a destination values as stored in the
	// fragment's frame buffer.
	BlendEquation(eq BlendEquation) ContextStateValue

	// DepthMask sets whether or not you can write to the depth buffer.
	DepthMask(m bool) ContextStateValue

	// Enable enables the given feature.
	Enable(f Feature) ContextStateValue

	// Disable disables the given feature.
	Disable(f Feature) ContextStateValue

	// Viewport sets the rectangular viewable area that contains the rendering
	// results of the drawing buffer.
	Viewport(x, y, width, height int) ContextStateValue

	// Scissor sets the dimensions of the scissor box.
	Scissor(x, y, width, height int) ContextStateValue

	// LineWidth specifies the width of rasterized lines. The initial value is 1.
	//
	// The actual width is determined by rounding the supplied width to the
	// nearest integer. (If the rounding results in the value 0, it is as if
	// the line width were 1.) If ∣Δx∣>=∣Δy∣, i pixels are filled in each
	// column that is rasterized, where i is the rounded value of width.
	// Otherwise, i pixels are filled in each row that is rasterized.
	//
	// There is a range of supported line widths. Only width 1 is guaranteed to
	// be supported; others depend on the implementation. To query the range of
	// supported widths, call Get with argument AliasedLineWidthRange.
	LineWidth(w float32) ContextStateValue

	// ColorMask lets you set whether individual colors can be written when
	// drawing or rendering to a framebuffer.
	//
	// The default value is true, all colors can be written to the framebuffer.
	// false on any parameter disables that color from being written.
	ColorMask(r, g, b, a bool) ContextStateValue

	// CullFace sets which facets are candidates for culling.
	CullFace(f Facet) ContextStateValue

	// FrontFace sets the orientation of front-facing polygons.
	FrontFace(o Orientation) ContextStateValue
}
