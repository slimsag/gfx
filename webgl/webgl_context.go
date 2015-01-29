// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"errors"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Context implements the gfx.Context interface.
type Context struct {
	// Object is literally the WebGLRenderingContext JavaScript object.
	js.Object

	// The default framebuffer implementation for the context.
	*Framebuffer

	// WebGL error codes (see the Check method).
	NO_ERROR                      int `js:"NO_ERROR"`
	OUT_OF_MEMORY                 int `js:"OUT_OF_MEMORY"`
	INVALID_ENUM                  int `js:"INVALID_ENUM"`
	INVALID_OPERATION             int `js:"INVALID_OPERATION"`
	INVALID_FRAMEBUFFER_OPERATION int `js:"INVALID_FRAMEBUFFER_OPERATION"`
	INVALID_VALUE                 int `js:"INVALID_VALUE"`
	CONTEXT_LOST_WEBGL            int `js:"CONTEXT_LOST_WEBGL"`

	FRAMEBUFFER        int `js:"FRAMEBUFFER"`
	DEPTH_BUFFER_BIT   int `js:"DEPTH_BUFFER_BIT"`
	STENCIL_BUFFER_BIT int `js:"STENCIL_BUFFER_BIT"`
	COLOR_BUFFER_BIT   int `js:"COLOR_BUFFER_BIT"`
}

func (c *Context) fastBindFramebuffer(framebuffer js.Object) {
	// TODO(slimsag): check that framebuffer is not already bound.
	c.Call("bindFramebuffer", c.FRAMEBUFFER, framebuffer)
}

func (c *Context) fastClearColor(v [4]float32) {
	// TODO(slimsag): check that clearColor is not already the given one.
	c.Call("clearColor", v[0], v[1], v[2], v[3])
}

func (c *Context) fastClearDepth(v float32) {
	// TODO(slimsag): check that clearDepth is not already the given one.
	c.Call("clearDepth", v)
}

func (c *Context) fastClearStencil(v int) {
	// TODO(slimsag): check that clearStencil is not already the given one.
	c.Call("clearStencil", v)
}

// Check implements the gfx.Context interface.
func (c *Context) Check() {
	e := c.Call("getError").Int()

	// Avoid the larger switch statement below, as no error is the most likely
	// case.
	if e == c.NO_ERROR {
		return
	}

	switch e {
	case c.OUT_OF_MEMORY:
		panic(gfx.OutOfMemory)
	case c.INVALID_ENUM:
		panic(gfx.InvalidEnum)
	case c.INVALID_OPERATION:
		panic(gfx.InvalidOperation)
	case c.INVALID_FRAMEBUFFER_OPERATION:
		panic(gfx.InvalidFramebufferOperation)
	case c.INVALID_VALUE:
		panic(gfx.InvalidValue)
	case c.CONTEXT_LOST_WEBGL:
		panic(gfx.ContextLost)
	default:
		panic(fmt.Sprintf("webgl: unhandled error 0x%X\n", e))
	}
}

// Wrap returns a new WebGL rendering context by wrapping the given JavaScript
// WebGLRenderingContext object.
func Wrap(o js.Object) gfx.Context {
	ctx := &Context{
		Object: o,
	}
	ctx.Framebuffer = &Framebuffer{
		Object: nil, // Default framebuffer object.
		ctx:    ctx,
	}
	return ctx
}

// ContextAttributes is a set of drawing surface attributes passed to New.
type ContextAttributes struct {
	// If the value is true, the drawing buffer has an alpha channel for the
	// purposes of performing OpenGL destination alpha operations and
	// compositing with the page. If the value is false, no alpha buffer is
	// available.
	Alpha bool

	// If the value is true, the drawing buffer has a depth buffer of at least
	// 16 bits. If the value is false, no depth buffer is available.
	Depth bool

	// If the value is true, the drawing buffer has a stencil buffer of at
	// least 8 bits. If the value is false, no stencil buffer is available.
	Stencil bool

	// If the value is true and the implementation supports antialiasing the
	// drawing buffer will perform antialiasing using its choice of technique
	// (multisample/supersample) and quality. If the value is false or the
	// implementation does not support antialiasing, no antialiasing is
	// performed.
	Antialias bool

	// If the value is true the page compositor will assume the drawing buffer
	// contains colors with premultiplied alpha. If the value is false the page
	// compositor will assume that colors in the drawing buffer are not
	// premultiplied. This flag is ignored if the alpha flag is false.
	PremultipliedAlpha bool

	// If false, once the drawing buffer is presented as described in the
	// "Drawing Buffer" section, the contents of the drawing buffer are cleared
	// to their default values. All elements of the drawing buffer (color,
	// depth and stencil) are cleared. If the value is true the buffers will
	// not be cleared and will preserve their values until cleared or
	// overwritten by the author.
	//
	// On some hardware setting the preserveDrawingBuffer flag to true can have
	// significant performance implications.
	PreserveDrawingBuffer bool

	// Provides a hint to the implementation suggesting that, if possible, it
	// creates a context that optimizes for power consumption over performance.
	// For example, on hardware that has more than one GPU, it may be the case
	// that one of them is less powerful but also uses less power. An
	// implementation may choose to, and may have to, ignore this hint.
	PreferLowPowerToHighPerformance bool

	// If the value is true, context creation will fail if the implementation
	// determines that the performance of the created WebGL context would be
	// dramatically lower than that of a native application making equivalent
	// OpenGL calls. This could happen for a number of reasons, including:
	//
	// An implementation might switch to a software rasterizer if the user's
	// GPU driver is known to be unstable.
	//
	// An implementation might require reading back the framebuffer from GPU
	// memory to system memory before compositing it with the rest of the page,
	// significantly reducing performance.
	//
	// Applications that don't require high performance should leave this
	// parameter at its default value of false. Applications that require high
	// performance may set this parameter to true, and if context creation
	// fails then the application may prefer to use a fallback rendering path
	// such as a 2D canvas context. Alternatively the application can retry
	// WebGL context creation with this parameter set to false, with the
	// knowledge that a reduced-fidelity rendering mode should be used to
	// improve performance.
	FailIfMajorPerformanceCaveat bool
}

// DefaultAttributes returns a copy of the default context attributes.
func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{
		Alpha:                           true,
		Depth:                           true,
		Stencil:                         false,
		Antialias:                       true,
		PremultipliedAlpha:              true,
		PreserveDrawingBuffer:           false,
		PreferLowPowerToHighPerformance: false,
		FailIfMajorPerformanceCaveat:    false,
	}
}

var (
	ErrNoWebGLSupport = errors.New("browser does not support WebGL")
)

// New takes an HTML5 canvas object and context attributes (nil for the default
// ones). If any error is returned, it will be of type ErrNoWebGLSupport
func New(canvas js.Object, ca *ContextAttributes) (gfx.Context, error) {
	if js.Global.Get("WebGLRenderingContext") == js.Undefined {
		return nil, ErrNoWebGLSupport
	}

	// Build the attribute dictionary needed by the JavaScript method.
	if ca == nil {
		ca = DefaultAttributes()
	}
	attrs := map[string]bool{
		"alpha":                           ca.Alpha,
		"depth":                           ca.Depth,
		"stencil":                         ca.Stencil,
		"antialias":                       ca.Antialias,
		"premultipliedAlpha":              ca.PremultipliedAlpha,
		"preserveDrawingBuffer":           ca.PreserveDrawingBuffer,
		"preferLowPowerToHighPerformance": ca.PreferLowPowerToHighPerformance,
		"failIfMajorPerformanceCaveat":    ca.FailIfMajorPerformanceCaveat,
	}

	// First try for standard "webgl" mode from Canvas.getContext.
	ctx := canvas.Call("getContext", "webgl", attrs)
	if ctx != nil {
		return Wrap(ctx), nil
	}

	// Next try for older "experimental-webgl" mode from Canvas.getContext.
	ctx = canvas.Call("getContext", "experimental-webgl", attrs)
	if ctx == nil {
		return nil, ErrNoWebGLSupport
	}
	return Wrap(ctx), nil
}
