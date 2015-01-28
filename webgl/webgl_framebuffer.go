// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
)

// Framebuffer implements the gfx.Framebuffer interface by wrapping a
// WebGLFramebuffer JavaScript object.
type Framebuffer struct {
	// Object is literally the WebGLFramebuffer object (or nil in the case of
	// the default framebuffer).
	js.Object

	ctx *Context
}

// Clear implements the gfx.Framebuffer interface.
func (f Framebuffer) Clear(ops ...interface{}) {
	// Check for an invalid number of arguments.
	if len(ops) == 0 || len(ops) > 3 {
		panic("Framebuffer.Clear: invalid number of arguments")
	}

	// Build the proper bitmask and store the needed clear values.
	var mask int
	for _, op := range ops {
		switch v := op.(type) {
		case gfx.ClearColor:
			mask |= f.ctx.COLOR_BUFFER_BIT
			f.ctx.fastClearColor(v)
		case gfx.ClearDepth:
			mask |= f.ctx.DEPTH_BUFFER_BIT
			f.ctx.fastClearDepth(v)
		case gfx.ClearStencil:
			mask |= f.ctx.STENCIL_BUFFER_BIT
			f.ctx.fastClearStencil(v)
		default:
			panic("Framebuffer.Clear: invalid operation")
		}
	}

	// Bind the framebuffer, clear with the bitmask.
	f.ctx.fastBindFramebuffer(f.Object)
	f.ctx.Call("clear", mask)
}
