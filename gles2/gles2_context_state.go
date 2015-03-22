// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build arm gles2

package gles2

import (
	"github.com/slimsag/gfx"
	gl "github.com/slimsag/gfx/internal/gles2/2.0/gles2"
)

type csv struct {
	key                 interface{}
	value, defaultValue interface{}
	glCall              func(value interface{})
}

type contextState []gfx.ContextStateValue

func (c contextState) find(k interface{}) (index int, pair csv) {
	var i interface{}
	for index, i = range c {
		pair = i.(csv)
		if pair.key != k {
			// Non-equal keys.
			continue
		}
		return index, pair
	}
	return -1, csv{}
}

func (c *Context) NewState(values ...gfx.ContextStateValue) gfx.ContextState {
	return contextState(values)
}

func (c *Context) Load(s gfx.ContextState) {
	var st contextState
	if s != nil {
		st = s.(contextState)
	}

	// For any state not explicitly mentioned in the current state, revert it
	// to the default state.
	for _, curI := range c.current {
		cur := curI.(csv)
		if index, _ := st.find(cur.key); index != -1 {
			continue
		}

		// Revert to the default state.
		if cur.value == cur.defaultValue {
			// Already using this value! Do nothing.
			continue
		}
		cur.glCall(cur.defaultValue)
	}

	// For each state explicitly mentioned in the destination state, apply it
	// if needed.
	for _, dstI := range st {
		dst := dstI.(csv)
		found, cur := c.current.find(dst.key)
		if found != -1 && cur.value == dst.value {
			// Already using this value! Do nothing.
			continue
		}

		// Did not find a matching previous value.
		dst.glCall(dst.value)
	}
	c.current = st
}

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
)

func glBlendColor(v interface{}) {
	x := v.([4]float32)
	gl.BlendColor(x[0], x[0], x[0], x[0])
}

// BlendColor implements the gfx.ContextStateProvider interface.
func (c *Context) BlendColor(r, g, b, a float32) gfx.ContextStateValue {
	return csv{
		value:        [4]float32{r, g, b, a},
		defaultValue: [4]float32{0, 0, 0, 0}, // TODO(slimsag): verify
		key:          csBlendColor,
		glCall:       glBlendColor,
	}
}

func glBlendEquation(v interface{}) {
	gl.BlendEquation(v.(uint32))
}

// BlendEquation implements the gfx.ContextStateProvider interface.
func (c *Context) BlendEquation(eq gfx.BlendEquation) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(eq)],
		defaultValue: uint32(gl.FUNC_ADD), // TODO(slimsag): verify
		key:          csBlendEquation,
		glCall:       glBlendEquation,
	}
}

func glDepthMask(v interface{}) {
	gl.DepthMask(v.(bool))
}

// DepthMask implements the gfx.ContextStateProvider interface.
func (c *Context) DepthMask(m bool) gfx.ContextStateValue {
	return csv{
		value:        m,
		defaultValue: true, // TODO(slimsag): verify
		key:          csDepthMask,
		glCall:       glDepthMask,
	}
}

func glUseProgram(v interface{}) {
	gl.UseProgram(v.(uint32))
}

// UseProgram implements the gfx.ContextStateProvider interface.
func (c *Context) UseProgram(p gfx.Program) gfx.ContextStateValue {
	return csv{
		value:        p,
		defaultValue: nil, // TODO(slimsag): verify
		key:          csUseProgram,
		glCall:       glUseProgram,
	}
}

func glViewport(v interface{}) {
	x := v.([4]int32)
	gl.Viewport(x[0], x[0], x[0], x[0])
}

// Viewport implements the gfx.ContextStateProvider interface.
func (c *Context) Viewport(x, y, width, height int) gfx.ContextStateValue {
	return csv{
		value:        [4]int32{int32(x), int32(y), int32(width), int32(height)},
		defaultValue: [4]int32{0, 0, 0, 0}, // TODO(slimsag): track real default viewport values
		key:          csViewport,
		glCall:       glViewport,
	}
}

func glScissor(v interface{}) {
	x := v.([4]int32)
	gl.Scissor(x[0], x[0], x[0], x[0])
}

// Scissor implements the gfx.ContextStateProvider interface.
func (c *Context) Scissor(x, y, width, height int) gfx.ContextStateValue {
	return csv{
		value:        [4]int32{int32(x), int32(y), int32(width), int32(height)},
		defaultValue: [4]int32{0, 0, 0, 0}, // TODO(slimsag): track real default scissor values
		key:          csScissor,
		glCall:       glScissor,
	}
}

func glLineWidth(v interface{}) {
	gl.LineWidth(v.(float32))
}

// LineWidth implements the gfx.ContextStateProvider interface.
func (c *Context) LineWidth(w float32) gfx.ContextStateValue {
	return csv{
		value:        w,
		defaultValue: 1.0,
		key:          csLineWidth,
		glCall:       glLineWidth,
	}
}

func glColorMask(v interface{}) {
	x := v.([4]bool)
	gl.ColorMask(x[0], x[1], x[2], x[3])
}

// ColorMask implements the gfx.ContextStateProvider interface.
func (c *Context) ColorMask(r, g, b, a bool) gfx.ContextStateValue {
	return csv{
		value:        [4]bool{r, g, b, a},
		defaultValue: [4]bool{true, true, true, true},
		key:          csColorMask,
		glCall:       glColorMask,
	}
}

func glCullFace(v interface{}) {
	gl.CullFace(v.(uint32))
}

// CullFace implements the gfx.ContextStateProvider interface.
func (c *Context) CullFace(f gfx.Facet) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(f)],
		defaultValue: uint32(gl.FRONT), // TODO(slimsag): verify
		key:          csCullFace,
		glCall:       glCullFace,
	}
}

func glFrontFace(v interface{}) {
	gl.FrontFace(v.(uint32))
}

// FrontFace implements the gfx.ContextStateProvider interface.
func (c *Context) FrontFace(o gfx.Orientation) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(o)],
		defaultValue: uint32(gl.CCW), // TODO(slimsag): verify
		key:          csFrontFace,
		glCall:       glFrontFace,
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
	return csv{
		value:        c.Enums[int(f)],
		defaultValue: 0, // TODO(slimsag): default feature values!
		key: featureKey{
			csKey: csEnable,
			f:     f,
		},
		glCall: glEnable,
	}
}

func glDisable(v interface{}) {
	gl.Disable(v.(uint32))
}

// Disable implements the gfx.ContextStateProvider interface.
func (c *Context) Disable(f gfx.Feature) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(f)],
		defaultValue: 0, // TODO(slimsag): default feature values!
		key: featureKey{
			csKey: csDisable,
			f:     f,
		},
		glCall: glDisable,
	}
}
