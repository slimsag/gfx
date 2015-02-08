// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"fmt"

	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
)

// Context implements the gfx.Context interface.
type Context struct {
	// The default framebuffer implementation for the context.
	*Framebuffer

	// Enums maps a gfx enumeration to it's cooresponding OpenGL one.
	Enums *[gfx.EnumMax]uint32

	LastBindFramebuffer  uint32
	LastBindRenderbuffer uint32
	LastClearColor       [4]float32
	LastClearDepth       float64
	LastClearStencil     int
	LastColorMask        [4]bool
	LastCullFace         gfx.Facet
	LastFrontFace        gfx.Orientation
}

func (c *Context) putEnum(gfxEnum int, glEnum uint32) {
	if glEnum == 0 {
		fmt.Println("gfxEnum:", gfxEnum)
		fmt.Println("glEnum:", glEnum)
		panic("putEnum: got invalid enum")
	}
	c.Enums[gfxEnum] = glEnum
}

func (c *Context) loadEnums() {
	c.Enums = new([gfx.EnumMax]uint32)

	// Framebuffer attachment points.
	c.putEnum(int(gfx.ColorAttachment0), gl.COLOR_ATTACHMENT0)
	c.putEnum(int(gfx.DepthAttachment), gl.DEPTH_ATTACHMENT)
	c.putEnum(int(gfx.StencilAttachment), gl.STENCIL_ATTACHMENT)
	c.putEnum(int(gfx.DepthStencilAttachment), gl.DEPTH_STENCIL_ATTACHMENT)

	// Texture targets.
	c.putEnum(int(gfx.Texture2D), gl.TEXTURE_2D)
	c.putEnum(int(gfx.TextureCubeMapPositiveX), gl.TEXTURE_CUBE_MAP_POSITIVE_X)
	c.putEnum(int(gfx.TextureCubeMapNegativeX), gl.TEXTURE_CUBE_MAP_NEGATIVE_X)
	c.putEnum(int(gfx.TextureCubeMapPositiveY), gl.TEXTURE_CUBE_MAP_POSITIVE_Y)
	c.putEnum(int(gfx.TextureCubeMapNegativeY), gl.TEXTURE_CUBE_MAP_NEGATIVE_Y)
	c.putEnum(int(gfx.TextureCubeMapPositiveZ), gl.TEXTURE_CUBE_MAP_POSITIVE_Z)
	c.putEnum(int(gfx.TextureCubeMapNegativeZ), gl.TEXTURE_CUBE_MAP_NEGATIVE_Z)

	// Renderbuffer storage formats.
	c.putEnum(int(gfx.RGBA4), gl.RGBA4)
	c.putEnum(int(gfx.RGB565), gl.RGB565)
	c.putEnum(int(gfx.RGB5A1), gl.RGB5_A1)
	c.putEnum(int(gfx.DepthComponent16), gl.DEPTH_COMPONENT16)

	// Features.
	c.putEnum(int(gfx.Blend), gl.BLEND)
	c.putEnum(int(gfx.DepthTest), gl.DEPTH_TEST)
	c.putEnum(int(gfx.CullFace), gl.CULL_FACE)
	c.putEnum(int(gfx.PolygonOffsetFill), gl.POLYGON_OFFSET_FILL)
	c.putEnum(int(gfx.ScissorTest), gl.SCISSOR_TEST)

	// Orientations.
	c.putEnum(int(gfx.CCW), gl.CCW)
	c.putEnum(int(gfx.CW), gl.CW)

	// Facets.
	c.putEnum(int(gfx.Front), gl.FRONT)
	c.putEnum(int(gfx.Back), gl.BACK)
	c.putEnum(int(gfx.FrontAndBack), gl.FRONT_AND_BACK)
}

func (c *Context) fastBindFramebuffer(framebuffer uint32) {
	if c.LastBindFramebuffer == framebuffer {
		return
	}
	c.LastBindFramebuffer = framebuffer
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
}

func (c *Context) fastBindRenderbuffer(renderbuffer uint32) {
	if c.LastBindRenderbuffer == renderbuffer {
		return
	}
	c.LastBindRenderbuffer = renderbuffer
	gl.BindRenderbuffer(gl.RENDERBUFFER, renderbuffer)
}

func (c *Context) fastClearColor(v [4]float32) {
	if c.LastClearColor == v {
		return
	}
	c.LastClearColor = v
	gl.ClearColor(v[0], v[1], v[2], v[3])
}

func (c *Context) fastClearDepth(v float64) {
	if c.LastClearDepth == v {
		return
	}
	c.LastClearDepth = v
	gl.ClearDepthf(float32(v))
}

func (c *Context) fastClearStencil(v int) {
	if c.LastClearStencil == v {
		return
	}
	c.LastClearStencil = v
	gl.ClearStencil(int32(v))
}

// NewFramebuffer implements the gfx.Context interface.
func (c *Context) NewFramebuffer() gfx.Framebuffer {
	fb := new(Framebuffer)
	gl.GenFramebuffers(1, &fb.Object)
	return fb
}

// NewRenderbuffer implements the gfx.Context interface.
func (c *Context) NewRenderbuffer() gfx.Renderbuffer {
	rb := new(Renderbuffer)
	gl.GenRenderbuffers(1, &rb.Object)
	return rb
}

// Enable implements the gfx.Context interface.
func (c *Context) Enable(f gfx.Feature) {
	gl.Enable(c.Enums[int(f)])
}

// Disable implements the gfx.Context interface.
func (c *Context) Disable(f gfx.Feature) {
	gl.Disable(c.Enums[int(f)])
}

// ColorMask implements the gfx.Context interface.
func (c *Context) ColorMask(r, g, b, a bool) {
	if c.LastColorMask == [4]bool{r, g, b, a} {
		return
	}
	gl.ColorMask(r, g, b, a)
}

// CullFace implements the gfx.Context interface.
func (c *Context) CullFace(f gfx.Facet) {
	if c.LastCullFace == f {
		return
	}
	gl.CullFace(c.Enums[int(f)])
}

// FrontFace implements the gfx.Context interface.
func (c *Context) FrontFace(o gfx.Orientation) {
	if c.LastFrontFace == o {
		return
	}
	gl.FrontFace(c.Enums[int(o)])
}

// Check implements the gfx.Context interface.
func (c *Context) Check() {
	e := gl.GetError()

	// Avoid the larger switch statement below, as no error is the most likely
	// case.
	if e == gl.NO_ERROR {
		return
	}

	switch e {
	case gl.OUT_OF_MEMORY:
		panic(gfx.OutOfMemory)
	case gl.INVALID_ENUM:
		panic(gfx.InvalidEnum)
	case gl.INVALID_OPERATION:
		panic(gfx.InvalidOperation)
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		panic(gfx.InvalidFramebufferOperation)
	case gl.INVALID_VALUE:
		panic(gfx.InvalidValue)
	case gl.STACK_OVERFLOW:
		panic(gfx.StackOverflow)
	case gl.STACK_UNDERFLOW:
		panic(gfx.StackUnderflow)
	case gl.CONTEXT_LOST:
		panic(gfx.ContextLost)
	default:
		panic(fmt.Sprintf("gles2: unhandled error 0x%X\n", e))
	}
}

// Flush implements the gfx.Context interface.
func (c *Context) Flush() {
	gl.Flush()
}

// Finish implements the gfx.Context interface.
func (c *Context) Finish() {
	gl.Finish()
}

// New returns a new OpenGL ES 2 graphics context. It must only be called under
// the presence of an active OpenGL context in the OS thread.
func New() (gfx.Context, error) {
	if err := gl.Init(); err != nil {
		return nil, err
	}

	ctx := &Context{}
	ctx.Framebuffer = &Framebuffer{
		Object: 0, // Default framebuffer object.
		ctx:    ctx,
	}
	ctx.loadEnums()
	return ctx, nil
}
