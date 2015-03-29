// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// checker is a gfx.Context that implicitly invokes the Check method of the
// underlying context after each function call is made. Thus, if any error
// should occur you will receive a nice stack trace where that error occured.
type checker struct {
	fb  *fbChecker
	ctx gfx.Context
}

// Framebuffer implements the gfx.Context interface.
func (c *checker) Framebuffer() gfx.Framebuffer {
	return c.fb
}

// NewFramebuffer implements the gfx.Context interface.
func (c *checker) NewFramebuffer() gfx.Framebuffer {
	v := c.ctx.NewFramebuffer()
	c.ctx.Check()
	return &fbChecker{
		fb:  v,
		ctx: c,
	}
}

// NewRenderbuffer implements the gfx.Context interface.
func (c *checker) NewRenderbuffer() gfx.Renderbuffer {
	v := c.ctx.NewRenderbuffer()
	c.ctx.Check()
	return &rbChecker{
		rb:  v,
		ctx: c,
	}
}

// NewShader implements the gfx.Context interface.
func (c *checker) NewShader(t gfx.ShaderType) gfx.Shader {
	s := c.ctx.NewShader(t)
	c.ctx.Check()
	return &shaderChecker{
		s:   s,
		ctx: c,
	}
}

// NewTexture implements the gfx.Context interface.
func (c *checker) NewTexture(t gfx.TextureType) gfx.Texture {
	tex := c.ctx.NewTexture(t)
	c.ctx.Check()
	return &textureChecker{
		t:   tex,
		ctx: c,
	}
}

// NewBuffer implements the gfx.Context interface.
func (c *checker) NewBuffer() gfx.Buffer {
	b := c.ctx.NewBuffer()
	c.ctx.Check()
	return &bufferChecker{
		b:   b,
		ctx: c,
	}
}

// NewProgram implements the gfx.Context interface.
func (c *checker) NewProgram() gfx.Program {
	p := c.ctx.NewProgram()
	c.ctx.Check()
	return &programChecker{
		p:   p,
		ctx: c,
	}
}

// NewState implements the gfx.Context interface.
func (c *checker) NewState(values ...gfx.ContextStateValue) gfx.ContextState {
	return c.ctx.NewState(values...)
}

// Load implements the gfx.Context interface.
func (c *checker) Load(s gfx.ContextState) {
	c.ctx.Load(s)
	c.ctx.Check()
}

// BlendColor implements the gfx.Context interface.
func (c *checker) BlendColor(r, g, b, a float32) gfx.ContextStateValue {
	return c.ctx.BlendColor(r, g, b, a)
}

// BlendEquation implements the gfx.Context interface.
func (c *checker) BlendEquation(eq gfx.BlendEquation) gfx.ContextStateValue {
	return c.ctx.BlendEquation(eq)
}

// DepthMask implements the gfx.Context interface.
func (c *checker) DepthMask(m bool) gfx.ContextStateValue {
	return c.ctx.DepthMask(m)
}

// Enable implements the gfx.Context interface.
func (c *checker) Enable(f gfx.Feature) gfx.ContextStateValue {
	return c.ctx.Enable(f)
}

// Disable implements the gfx.Context interface.
func (c *checker) Disable(f gfx.Feature) gfx.ContextStateValue {
	return c.ctx.Disable(f)
}

// UseProgram implements the gfx.Context interface.
func (c *checker) UseProgram(p gfx.Program) gfx.ContextStateValue {
	return c.ctx.UseProgram(p)
}

// Viewport implements the gfx.Context interface.
func (c *checker) Viewport(x, y, width, height int) gfx.ContextStateValue {
	return c.ctx.Viewport(x, y, width, height)
}

// Scissor implements the gfx.Context interface.
func (c *checker) Scissor(x, y, width, height int) gfx.ContextStateValue {
	return c.ctx.Scissor(x, y, width, height)
}

// LineWidth implements the gfx.Context interface.
func (c *checker) LineWidth(w float32) gfx.ContextStateValue {
	return c.ctx.LineWidth(w)
}

// ColorMask implements the gfx.Context interface.
func (c *checker) ColorMask(r, g, b, a bool) gfx.ContextStateValue {
	return c.ctx.ColorMask(r, g, b, a)
}

// CullFace implements the gfx.Context interface.
func (c *checker) CullFace(f gfx.Facet) gfx.ContextStateValue {
	return c.ctx.CullFace(f)
}

// FrontFace implements the gfx.Context interface.
func (c *checker) FrontFace(o gfx.Orientation) gfx.ContextStateValue {
	return c.ctx.FrontFace(o)
}

// Check implements the gfx.Context interface.
func (c *checker) Check() {
	// We don't want caller to accidently grab the error, so we stub out the
	// call here.
	return
}

// Flush implements the gfx.Context interface.
func (c *checker) Flush() {
	c.ctx.Flush()
	c.ctx.Check()
}

// Finish implements the gfx.Context interface.
func (c *checker) Finish() {
	c.ctx.Finish()
	c.ctx.Check()
}

// Checker wraps the given graphics context such that each function call to the
// context (or any object gotten from it, e.g. a Framebuffer) has an implicit
// Check() call after it.
//
// This ensures that, should any error occur in the context, you will receive
// a nice Go stack trace with the exact function where the error was made.
//
// Additionally, it will generate panics for any Framebuffer operations whose
// Status is != nil.
func Checker(c gfx.Context) gfx.Context {
	return &checker{
		fb: &fbChecker{
			fb:  c.Framebuffer(),
			ctx: c,
		},
		ctx: c,
	}
}
