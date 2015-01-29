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

	LastBindFramebuffer uint32
	LastClearColor      [4]float32
	LastClearDepth      float64
	LastClearStencil    int
}

func (c *Context) fastBindFramebuffer(framebuffer uint32) {
	if c.LastBindFramebuffer == framebuffer {
		return
	}
	c.LastBindFramebuffer = framebuffer
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
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
	ctx.Framebuffer = &Framebuffer{
		Object: 0, // Default framebuffer object.
		ctx:    ctx,
	}
	return ctx, nil
}
