// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build amd64,!gles2 386,!gles2

package gl2

import (
	"fmt"

	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/internal/gl/2.0/gl"
	"github.com/slimsag/gfx/internal/state"
)

// Context implements the gfx.Context interface.
type Context struct {
	state.Context

	// Enums maps a gfx enumeration to it's cooresponding OpenGL one.
	Enums [gfx.EnumMax]uint32

	LastBindFramebuffer  uint32
	LastBindRenderbuffer uint32
	LastBindBuffer       uint32
	LastUseProgram       uint32

	// The default framebuffer implementation for the context.
	fb Framebuffer

	puts int
}

func (c *Context) putEnum(gfxEnum int, glEnum uint32) {
	c.puts++
	if gfxEnum != int(gfx.Points) && glEnum == 0 {
		fmt.Println("gfxEnum:", gfxEnum)
		fmt.Println("glEnum:", glEnum)
		panic("putEnum: got invalid enum")
	}
	c.Enums[gfxEnum] = glEnum
}

func (c *Context) loadEnums() {
	// Texture targets.
	c.putEnum(int(gfx.Texture2D), gl.TEXTURE_2D)
	c.putEnum(int(gfx.TextureCubeMapPositiveX), gl.TEXTURE_CUBE_MAP_POSITIVE_X)
	c.putEnum(int(gfx.TextureCubeMapNegativeX), gl.TEXTURE_CUBE_MAP_NEGATIVE_X)
	c.putEnum(int(gfx.TextureCubeMapPositiveY), gl.TEXTURE_CUBE_MAP_POSITIVE_Y)
	c.putEnum(int(gfx.TextureCubeMapNegativeY), gl.TEXTURE_CUBE_MAP_NEGATIVE_Y)
	c.putEnum(int(gfx.TextureCubeMapPositiveZ), gl.TEXTURE_CUBE_MAP_POSITIVE_Z)
	c.putEnum(int(gfx.TextureCubeMapNegativeZ), gl.TEXTURE_CUBE_MAP_NEGATIVE_Z)

	// Texture types.
	c.putEnum(int(gfx.TextureType2D), gl.TEXTURE_2D)
	c.putEnum(int(gfx.TextureTypeCubeMap), gl.TEXTURE_CUBE_MAP)

	// Renderbuffer storage formats.
	c.putEnum(int(gfx.RGBA4), gl.RGBA4)
	c.putEnum(int(gfx.RGB565), gl.RGB565)
	c.putEnum(int(gfx.RGB5A1), gl.RGB5_A1)
	c.putEnum(int(gfx.DepthComponent16), gl.DEPTH_COMPONENT16)

	// Framebuffer attachment points.
	c.putEnum(int(gfx.ColorAttachment0), gl.COLOR_ATTACHMENT0)
	c.putEnum(int(gfx.DepthAttachment), gl.DEPTH_ATTACHMENT)
	c.putEnum(int(gfx.StencilAttachment), gl.STENCIL_ATTACHMENT)
	c.putEnum(int(gfx.DepthStencilAttachment), gl.DEPTH_STENCIL_ATTACHMENT)

	// Buffer usage hints.
	c.putEnum(int(gfx.StaticDraw), gl.STATIC_DRAW)
	c.putEnum(int(gfx.DynamicDraw), gl.DYNAMIC_DRAW)
	c.putEnum(int(gfx.StreamDraw), gl.STREAM_DRAW)

	// Buffer types.
	c.putEnum(int(gfx.ArrayBuffer), gl.ARRAY_BUFFER)
	c.putEnum(int(gfx.ElementArrayBuffer), gl.ELEMENT_ARRAY_BUFFER)

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

	// Shader types.
	c.putEnum(int(gfx.VertexShader), gl.VERTEX_SHADER)
	c.putEnum(int(gfx.FragmentShader), gl.FRAGMENT_SHADER)

	// Blend equations.
	c.putEnum(int(gfx.FuncAdd), gl.FUNC_ADD)
	c.putEnum(int(gfx.FuncSubtract), gl.FUNC_SUBTRACT)
	c.putEnum(int(gfx.FuncReverseSubtract), gl.FUNC_REVERSE_SUBTRACT)

	// Primitive types.
	c.putEnum(int(gfx.Points), gl.POINTS)
	c.putEnum(int(gfx.Lines), gl.LINES)
	c.putEnum(int(gfx.LineStrip), gl.LINE_STRIP)
	c.putEnum(int(gfx.LineLoop), gl.LINE_LOOP)
	c.putEnum(int(gfx.Triangles), gl.TRIANGLES)
	c.putEnum(int(gfx.TriangleStrip), gl.TRIANGLE_STRIP)
	c.putEnum(int(gfx.TriangleFan), gl.TRIANGLE_FAN)

	// Verify that we put all enums into the array.
	if c.puts != len(c.Enums) {
		for k, e := range c.Enums {
			if e == 0 {
				fmt.Println("Missing:", k)
			}
		}
		panic("Did not put all enums (see above)")
	}
}

func (c *Context) fastBindFramebuffer(framebuffer uint32) bool {
	if c.LastBindFramebuffer == framebuffer {
		return false
	}
	c.LastBindFramebuffer = framebuffer
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
	return true
}

func (c *Context) fastBindRenderbuffer(renderbuffer uint32) bool {
	if c.LastBindRenderbuffer == renderbuffer {
		return false
	}
	c.LastBindRenderbuffer = renderbuffer
	gl.BindRenderbuffer(gl.RENDERBUFFER, renderbuffer)
	return true
}

func (c *Context) fastBindBuffer(typ, buffer uint32) bool {
	if c.LastBindBuffer == buffer {
		return false
	}
	c.LastBindBuffer = buffer
	gl.BindBuffer(typ, buffer)
	return true
}

func (c *Context) fastUseProgram(program uint32) bool {
	if c.LastUseProgram == program {
		return false
	}
	c.LastUseProgram = program
	gl.UseProgram(program)
	return true
}

// Framebuffer implements the gfx.Context interface.
func (c *Context) Framebuffer() gfx.Framebuffer {
	return &c.fb
}

// NewFramebuffer implements the gfx.Context interface.
func (c *Context) NewFramebuffer() gfx.Framebuffer {
	fb := &Framebuffer{
		ctx: c,
	}
	gl.GenFramebuffers(1, &fb.o)
	return fb
}

// NewRenderbuffer implements the gfx.Context interface.
func (c *Context) NewRenderbuffer() gfx.Renderbuffer {
	rb := &Renderbuffer{
		ctx: c,
	}
	gl.GenRenderbuffers(1, &rb.o)
	return rb
}

// NewShader implements the gfx.Context interface.
func (c *Context) NewShader(t gfx.ShaderType) gfx.Shader {
	return &Shader{
		ctx: c,
		o:   gl.CreateShader(c.Enums[int(t)]),
	}
}

// NewTexture implements the gfx.Context interface.
func (c *Context) NewTexture(t gfx.TextureType) gfx.Texture {
	tex := &Texture{
		ctx: c,
		typ: t,
	}
	gl.GenTextures(1, &tex.o)
	return tex
}

// NewBuffer implements the gfx.Context interface.
func (c *Context) NewBuffer(t gfx.BufferType) gfx.Buffer {
	b := &Buffer{
		ctx: c,
		typ: t,
	}
	gl.GenBuffers(1, &b.o)
	return b
}

// NewProgram implements the gfx.Context interface.
func (c *Context) NewProgram() gfx.Program {
	return &Program{
		ctx: c,
		o:   gl.CreateProgram(),
	}
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

	ctx := &Context{}
	ctx.fb.o = 0 // Default framebuffer object.
	ctx.fb.ctx = ctx
	ctx.loadEnums()
	return ctx, nil
}
