// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import "errors"

// Panics that Context.Check may generate.
var (
	OutOfMemory                 = errors.New("out of memory")
	InvalidEnum                 = errors.New("invalid enum")
	InvalidOperation            = errors.New("invalid operation")
	InvalidFramebufferOperation = errors.New("invalid framebuffer operation")
	InvalidValue                = errors.New("invalid value")
	StackOverflow               = errors.New("stack overflow")
	StackUnderflow              = errors.New("stack underflow")
	ContextLost                 = errors.New("context lost")
)

// Context is a graphics context. Unlike traditional OpenGL contexts, it is not
// tied to a specific OS thread (but still must be accessed from only one
// goroutine/thread at a time).
type Context interface {
	// Framebuffer is the default framebuffer of the context, i.e. the one that
	// when drawn to, appears on the window.
	Framebuffer

	// NewFramebuffer returns a new Framebuffer object used for storing
	// complete frames.
	NewFramebuffer() Framebuffer

	// NewRenderbuffer returns a new Renderbuffer object used for storing
	// complete frames.
	NewRenderbuffer() Renderbuffer

	// NewShader returns a new Shader object which represents the content and
	// how exactly it is drawn to a render target.
	NewShader(t ShaderType) Shader

	// NewTexture returns a new Texture object which is used for images and
	// cube maps when rendering shapes.
	NewTexture() Texture

	// NewBuffer returns a new buffer object that contains data such as
	// vertices or colors.
	NewBuffer() Buffer

	// NewProgram returns a new program object which represents the
	// programmable OpenGL pipeline and associated shader programs.
	NewProgram() Program

	// BlendColor specifies the blend color used to calculate source and
	// destination blending.
	BlendColor(r, g, b, a float32)

	// BlendEquation sets the equation used to blend RGB and Alpha values of an
	// incoming source fragment with a destination values as stored in the
	// fragment's frame buffer.
	BlendEquation(eq BlendEquation)

	// DepthMask sets whether or not you can write to the depth buffer.
	DepthMask(m bool)

	// Enable enables the given feature.
	Enable(f Feature)

	// Disable disables the given feature.
	Disable(f Feature)

	// IsEnabled tells if the given feature is enabled or not.
	IsEnabled(f Feature) bool

	// Viewport sets the rectangular viewable area that contains the rendering
	// results of the drawing buffer.
	Viewport(x, y, width, height int)

	// Scissor sets the dimensions of the scissor box.
	Scissor(x, y, width, height int)

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
	//
	// TODO(slimsag): implement the Context.Get method.
	LineWidth(w float32)

	// ColorMask lets you set whether individual colors can be written when
	// drawing or rendering to a framebuffer.
	//
	// The default value is true, all colors can be written to the framebuffer.
	// false on any parameter disables that color from being written.
	ColorMask(r, g, b, a bool)

	// CullFace sets which facets are candidates for culling.
	CullFace(f Facet)

	// FrontFace sets the orientation of front-facing polygons.
	FrontFace(o Orientation)

	// Check checks that no errors have occured in the context. If an error
	// occurs, it is either a programmer error (passing an invalid value, etc)
	// or a serious device error (running out of memory, losing the context).
	Check()

	// Flush flushes any buffered commands out to the graphics hardware as
	// quickly as possible. Execution may not be completed in any particular
	// time period, but does complete in finite time.
	//
	// Drivers may buffer commands and queue them for sending to hardware in
	// large chunks. For this reason, all programs should call Flush whenever
	// they count on having all of the previously issued commands completed.
	Flush()

	// Finish blocks (does not return) until the effects of all previously
	// called commands are complete. Such effects include all changes to the
	// graphics state, all changes to connection state, and all changes to the
	// framebuffer contents.
	//
	// In most cases, you shouldn't ever use Finish, but rather Flush.
	Finish()
}
