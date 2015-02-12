// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"fmt"

	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
)

// Context implements the gfx.Context interface.
type Context struct {
	// The default framebuffer implementation for the context.
	*Framebuffer

	// Enums maps a gfx enumeration to it's cooresponding OpenGL one.
	Enums *[gfx.EnumMax]uint32

	// Feature is a map of gfx feature enumerations to their current enabled or
	// disabled status.
	Features *[gfx.LastFeature - gfx.FirstFeature]bool

	LastBindFramebuffer  uint32
	LastBindRenderbuffer uint32
	LastBlendColor       [4]float32
	LastBlendEquation    gfx.BlendEquation
	LastDepthMask        bool
	LastClearColor       [4]float32
	LastClearDepth       float64
	LastClearStencil     int
	LastViewport         [4]int
	LastScissor          [4]int
	LastLineWidth        float32
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

	// BlendEquations.
	c.putEnum(int(gfx.FuncAdd), gl.FUNC_ADD)
	c.putEnum(int(gfx.FuncSubtract), gl.FUNC_SUBTRACT)
	c.putEnum(int(gfx.FuncReverseSubtract), gl.FUNC_REVERSE_SUBTRACT)
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
	gl.ClearDepth(v)
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
	fb := &Framebuffer{
		ctx: c,
	}
	gl.GenFramebuffers(1, &fb.Object)
	return fb
}

// NewRenderbuffer implements the gfx.Context interface.
func (c *Context) NewRenderbuffer() gfx.Renderbuffer {
	rb := &Renderbuffer{
		ctx: c,
	}
	gl.GenRenderbuffers(1, &rb.Object)
	return rb
}

// NewShader implements the gfx.Context interface.
func (c *Context) NewShader(t gfx.ShaderType) gfx.Shader {
	return &Shader{
		ctx:    c,
		Object: gl.CreateShader(c.Enums[int(t)]),
	}
}

// NewTexture implements the gfx.Context interface.
func (c *Context) NewTexture() gfx.Texture {
	t := &Texture{
		ctx: c,
	}
	gl.GenTextures(1, &t.Object)
	return t
}

// NewBuffer implements the gfx.Context interface.
func (c *Context) NewBuffer() gfx.Buffer {
	b := &Buffer{
		ctx: c,
	}
	gl.GenBuffers(1, &b.Object)
	return b
}

// NewProgram implements the gfx.Context interface.
func (c *Context) NewProgram() gfx.Program {
	return &Program{
		ctx:    c,
		Object: gl.CreateProgram(),
	}
}

// BlendColor implements the gfx.Context interface.
func (c *Context) BlendColor(r, g, b, a float32) {
	if c.LastBlendColor == [4]float32{r, g, b, a} {
		return
	}
	c.LastBlendColor = [4]float32{r, g, b, a}
	gl.BlendColor(r, g, b, a)
}

// BlendEquation implements the gfx.Context interface.
func (c *Context) BlendEquation(eq gfx.BlendEquation) {
	if c.LastBlendEquation == eq {
		return
	}
	c.LastBlendEquation = eq
	gl.BlendEquation(c.Enums[int(eq)])
}

// DepthMask implements the gfx.Context interface.
func (c *Context) DepthMask(m bool) {
	if c.LastDepthMask == m {
		return
	}
	c.LastDepthMask = m
	gl.DepthMask(m)
}

// Enable implements the gfx.Context interface.
func (c *Context) Enable(f gfx.Feature) {
	if c.Features[f-gfx.FirstFeature] {
		return
	}
	c.Features[f-gfx.FirstFeature] = true
	gl.Enable(c.Enums[int(f)])
}

// Disable implements the gfx.Context interface.
func (c *Context) Disable(f gfx.Feature) {
	if !c.Features[f-gfx.FirstFeature] {
		return
	}
	c.Features[f-gfx.FirstFeature] = false
	gl.Disable(c.Enums[int(f)])
}

// Viewport implements the gfx.Context interface.
func (c *Context) Viewport(x, y, width, height int) {
	if c.LastViewport == [4]int{x, y, width, height} {
		return
	}
	c.LastViewport = [4]int{x, y, width, height}
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

// Scissor implements the gfx.Context interface.
func (c *Context) Scissor(x, y, width, height int) {
	if c.LastScissor == [4]int{x, y, width, height} {
		return
	}
	c.LastScissor = [4]int{x, y, width, height}
	gl.Scissor(int32(x), int32(y), int32(width), int32(height))
}

// LineWidth implements the gfx.Context interface.
func (c *Context) LineWidth(w float32) {
	if c.LastLineWidth == w {
		return
	}
	c.LastLineWidth = w
	gl.LineWidth(w)
}

// ColorMask implements the gfx.Context interface.
func (c *Context) ColorMask(r, g, b, a bool) {
	if c.LastColorMask == [4]bool{r, g, b, a} {
		return
	}
	c.LastColorMask = [4]bool{r, g, b, a}
	gl.ColorMask(r, g, b, a)
}

// CullFace implements the gfx.Context interface.
func (c *Context) CullFace(f gfx.Facet) {
	if c.LastCullFace == f {
		return
	}
	c.LastCullFace = f
	gl.CullFace(c.Enums[int(f)])
}

// FrontFace implements the gfx.Context interface.
func (c *Context) FrontFace(o gfx.Orientation) {
	if c.LastFrontFace != o {
		return
	}
	c.LastFrontFace = o
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
		panic(fmt.Sprintf("gl2: unhandled error 0x%X\n", e))
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

// New returns a new OpenGL 2 graphics context. It must only be called under
// the presence of an active OpenGL context in the OS thread.
func New() (gfx.Context, error) {
	if err := gl.Init(); err != nil {
		return nil, err
	}

	ctx := &Context{
		Enums:    new([gfx.EnumMax]uint32),
		Features: new([gfx.LastFeature - gfx.FirstFeature]bool),
	}
	ctx.Framebuffer = &Framebuffer{
		Object: 0, // Default framebuffer object.
		ctx:    ctx,
	}
	ctx.loadEnums()
	return ctx, nil
}
