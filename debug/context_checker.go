// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// checker is a gfx.Context that implicitly invokes the Check method of the
// underlying context after each function call is made. Thus, if any error
// should occur you will receive a nice stack trace where that error occured.
type checker struct {
	*fbChecker
	ctx gfx.Context
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
		fbChecker: &fbChecker{
			fb:  c.(gfx.Framebuffer),
			ctx: c,
		},
		ctx: c,
	}
}
