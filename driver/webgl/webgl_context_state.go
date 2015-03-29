// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import (
	"github.com/slimsag/gfx"
	s "github.com/slimsag/gfx/internal/state"
)

const (
	csBlendColor = iota
	csBlendEquation
	csDepthMask
	csUseProgram
	csViewport
	csScissor
	csLineWidth
	csColorMask
	csCullFace
	csFrontFace
	csEnable
	csDisable
	csEnableVertexAttribArray
)

func (c *Context) glBlendColor(v interface{}) {
	x := v.([4]float32)
	c.O.Call("blendColor", x[0], x[1], x[2], x[3])
}

// BlendColor implements the gfx.ContextStateProvider interface.
func (c *Context) BlendColor(r, g, b, a float32) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]float32{r, g, b, a},
		DefaultValue: [4]float32{0, 0, 0, 0}, // TODO(slimsag): verify
		Key:          csBlendColor,
		GLCall:       c.glBlendColor,
	}
}

func (c *Context) glBlendEquation(v interface{}) {
	c.O.Call("blendEquation", v.(int))
}

// BlendEquation implements the gfx.ContextStateProvider interface.
func (c *Context) BlendEquation(eq gfx.BlendEquation) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(eq)],
		DefaultValue: c.O.Get("FUNC_ADD").Int(), // TODO(slimsag): verify
		Key:          csBlendEquation,
		GLCall:       c.glBlendEquation,
	}
}

func (c *Context) glDepthMask(v interface{}) {
	c.O.Call("depthMask", v.(bool))
}

// DepthMask implements the gfx.ContextStateProvider interface.
func (c *Context) DepthMask(m bool) gfx.ContextStateValue {
	return s.CSV{
		Value:        m,
		DefaultValue: true, // TODO(slimsag): verify
		Key:          csDepthMask,
		GLCall:       c.glDepthMask,
	}
}

func (c *Context) glUseProgram(v interface{}) {
	c.O.Call("useProgram", v)
}

// UseProgram implements the gfx.ContextStateProvider interface.
func (c *Context) UseProgram(p gfx.Program) gfx.ContextStateValue {
	return s.CSV{
		Value:        p,
		DefaultValue: nil, // TODO(slimsag): verify
		Key:          csUseProgram,
		GLCall:       c.glUseProgram,
	}
}

func (c *Context) glViewport(v interface{}) {
	x := v.([4]int)
	c.O.Call("viewport", x[0], x[1], x[2], x[3])
}

// Viewport implements the gfx.ContextStateProvider interface.
func (c *Context) Viewport(x, y, width, height int) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]int{x, y, width, height},
		DefaultValue: [4]int{0, 0, 0, 0}, // TODO(slimsag): track real default viewport values
		Key:          csViewport,
		GLCall:       c.glViewport,
	}
}

func (c *Context) glScissor(v interface{}) {
	x := v.([4]int)
	c.O.Call("scissor", x[0], x[1], x[2], x[3])
}

// Scissor implements the gfx.ContextStateProvider interface.
func (c *Context) Scissor(x, y, width, height int) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]int{x, y, width, height},
		DefaultValue: [4]int{0, 0, 0, 0}, // TODO(slimsag): track real default scissor values
		Key:          csScissor,
		GLCall:       c.glScissor,
	}
}

func (c *Context) glLineWidth(v interface{}) {
	c.O.Call("lineWidth", v.(float64))
}

// LineWidth implements the gfx.ContextStateProvider interface.
func (c *Context) LineWidth(w float32) gfx.ContextStateValue {
	return s.CSV{
		Value:        w,
		DefaultValue: 1.0,
		Key:          csLineWidth,
		GLCall:       c.glLineWidth,
	}
}

func (c *Context) glColorMask(v interface{}) {
	x := v.([4]bool)
	c.O.Call("colorMask", x[0], x[1], x[2], x[3])
}

// ColorMask implements the gfx.ContextStateProvider interface.
func (c *Context) ColorMask(r, g, b, a bool) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]bool{r, g, b, a},
		DefaultValue: [4]bool{true, true, true, true},
		Key:          csColorMask,
		GLCall:       c.glColorMask,
	}
}

func (c *Context) glCullFace(v interface{}) {
	c.O.Call("cullFace", v.(int))
}

// CullFace implements the gfx.ContextStateProvider interface.
func (c *Context) CullFace(f gfx.Facet) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(f)],
		DefaultValue: c.O.Get("FRONT").Int(), // TODO(slimsag): verify
		Key:          csCullFace,
		GLCall:       c.glCullFace,
	}
}

func (c *Context) glFrontFace(v interface{}) {
	c.O.Call("frontFace", v.(int))
}

// FrontFace implements the gfx.ContextStateProvider interface.
func (c *Context) FrontFace(o gfx.Orientation) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(o)],
		DefaultValue: c.O.Get("CCW").Int(), // TODO(slimsag): verify
		Key:          csFrontFace,
		GLCall:       c.glFrontFace,
	}
}

type featureKey struct {
	csKey int
	f     gfx.Feature
}

func (c *Context) glEnable(v interface{}) {
	c.O.Call("enable", v.(int))
}

// Enable implements the gfx.ContextStateProvider interface.
func (c *Context) Enable(f gfx.Feature) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(f)],
		DefaultValue: 0, // TODO(slimsag): default feature values!
		Key: featureKey{
			csKey: csEnable,
			f:     f,
		},
		GLCall: c.glEnable,
	}
}

func (c *Context) glDisable(v interface{}) {
	c.O.Call("disable", v.(int))
}

// Disable implements the gfx.ContextStateProvider interface.
func (c *Context) Disable(f gfx.Feature) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(f)],
		DefaultValue: 0, // TODO(slimsag): default feature values!
		Key: featureKey{
			csKey: csDisable,
			f:     f,
		},
		GLCall: c.glDisable,
	}
}

func (c *Context) glEnableVertexAttribArray(v interface{}) {
	c.O.Call("enableVertexAttribArray", v.(int))
}

// EnableVertexAttribArray implements the gfx.ContextStateProvider interface.
func (c *Context) EnableVertexAttribArray(l gfx.AttribLocation) gfx.ContextStateValue {
	return s.CSV{
		Value:        l,
		DefaultValue: 0, // TODO(slimsag): default enable VAA values!
		Key:          csEnableVertexAttribArray,
		GLCall:       c.glEnableVertexAttribArray,
	}
}
