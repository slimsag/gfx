// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
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

func glBlendColor(v interface{}) {
	x := v.([4]float32)
	gl.BlendColor(x[0], x[1], x[2], x[3])
}

// BlendColor implements the gfx.ContextStateProvider interface.
func (c *Context) BlendColor(r, g, b, a float32) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]float32{r, g, b, a},
		DefaultValue: [4]float32{0, 0, 0, 0}, // TODO(slimsag): verify
		Key:          csBlendColor,
		GLCall:       glBlendColor,
	}
}

func glBlendEquation(v interface{}) {
	gl.BlendEquation(v.(uint32))
}

// BlendEquation implements the gfx.ContextStateProvider interface.
func (c *Context) BlendEquation(eq gfx.BlendEquation) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(eq)],
		DefaultValue: uint32(gl.FUNC_ADD), // TODO(slimsag): verify
		Key:          csBlendEquation,
		GLCall:       glBlendEquation,
	}
}

func glDepthMask(v interface{}) {
	gl.DepthMask(v.(bool))
}

// DepthMask implements the gfx.ContextStateProvider interface.
func (c *Context) DepthMask(m bool) gfx.ContextStateValue {
	return s.CSV{
		Value:        m,
		DefaultValue: true, // TODO(slimsag): verify
		Key:          csDepthMask,
		GLCall:       glDepthMask,
	}
}

func glUseProgram(v interface{}) {
	gl.UseProgram(v.(uint32))
}

// UseProgram implements the gfx.ContextStateProvider interface.
func (c *Context) UseProgram(p gfx.Program) gfx.ContextStateValue {
	return s.CSV{
		Value:        p,
		DefaultValue: nil, // TODO(slimsag): verify
		Key:          csUseProgram,
		GLCall:       glUseProgram,
	}
}

func glViewport(v interface{}) {
	x := v.([4]int32)
	gl.Viewport(x[0], x[1], x[2], x[3])
}

// Viewport implements the gfx.ContextStateProvider interface.
func (c *Context) Viewport(x, y, width, height int) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]int32{int32(x), int32(y), int32(width), int32(height)},
		DefaultValue: [4]int32{0, 0, 0, 0}, // TODO(slimsag): track real default viewport values
		Key:          csViewport,
		GLCall:       glViewport,
	}
}

func glScissor(v interface{}) {
	x := v.([4]int32)
	gl.Scissor(x[0], x[1], x[2], x[3])
}

// Scissor implements the gfx.ContextStateProvider interface.
func (c *Context) Scissor(x, y, width, height int) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]int32{int32(x), int32(y), int32(width), int32(height)},
		DefaultValue: [4]int32{0, 0, 0, 0}, // TODO(slimsag): track real default scissor values
		Key:          csScissor,
		GLCall:       glScissor,
	}
}

func glLineWidth(v interface{}) {
	gl.LineWidth(v.(float32))
}

// LineWidth implements the gfx.ContextStateProvider interface.
func (c *Context) LineWidth(w float32) gfx.ContextStateValue {
	return s.CSV{
		Value:        w,
		DefaultValue: 1.0,
		Key:          csLineWidth,
		GLCall:       glLineWidth,
	}
}

func glColorMask(v interface{}) {
	x := v.([4]bool)
	gl.ColorMask(x[0], x[1], x[2], x[3])
}

// ColorMask implements the gfx.ContextStateProvider interface.
func (c *Context) ColorMask(r, g, b, a bool) gfx.ContextStateValue {
	return s.CSV{
		Value:        [4]bool{r, g, b, a},
		DefaultValue: [4]bool{true, true, true, true},
		Key:          csColorMask,
		GLCall:       glColorMask,
	}
}

func glCullFace(v interface{}) {
	gl.CullFace(v.(uint32))
}

// CullFace implements the gfx.ContextStateProvider interface.
func (c *Context) CullFace(f gfx.Facet) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(f)],
		DefaultValue: uint32(gl.FRONT), // TODO(slimsag): verify
		Key:          csCullFace,
		GLCall:       glCullFace,
	}
}

func glFrontFace(v interface{}) {
	gl.FrontFace(v.(uint32))
}

// FrontFace implements the gfx.ContextStateProvider interface.
func (c *Context) FrontFace(o gfx.Orientation) gfx.ContextStateValue {
	return s.CSV{
		Value:        c.Enums[int(o)],
		DefaultValue: uint32(gl.CCW), // TODO(slimsag): verify
		Key:          csFrontFace,
		GLCall:       glFrontFace,
	}
}

type featureKey struct {
	csKey int
	f     gfx.Feature
}

func glEnable(v interface{}) {
	gl.Enable(v.(uint32))
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
		GLCall: glEnable,
	}
}

func glDisable(v interface{}) {
	gl.Disable(v.(uint32))
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
		GLCall: glDisable,
	}
}

func glEnableVertexAttribArray(v interface{}) {
	gl.EnableVertexAttribArray(uint32(v.(int32)))
}

// EnableVertexAttribArray implements the gfx.ContextStateProvider interface.
func (c *Context) EnableVertexAttribArray(l gfx.AttribLocation) gfx.ContextStateValue {
	return s.CSV{
		Value:        l.(int32),
		DefaultValue: 0, // TODO(slimsag): default enable VAA values!
		Key:          csEnableVertexAttribArray,
		GLCall:       glEnableVertexAttribArray,
	}
}
