// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package webgl

import "github.com/slimsag/gfx"

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
	csViewport
	csScissor
	csLineWidth
	csColorMask
	csCullFace
	csFrontFace
	csEnable
	csDisable
)

func (c *Context) glBlendColor(v interface{}) {
	x := v.([4]float32)
	c.Object.Call("blendColor", x[0], x[0], x[0], x[0])
}

// BlendColor implements the gfx.ContextStateProvider interface.
func (c *Context) BlendColor(r, g, b, a float32) gfx.ContextStateValue {
	return csv{
		value:        [4]float32{r, g, b, a},
		defaultValue: [4]float32{0, 0, 0, 0}, // TODO(slimsag): verify
		key:          csBlendColor,
		glCall:       c.glBlendColor,
	}
}

func (c *Context) glBlendEquation(v interface{}) {
	c.Object.Call("blendEquation", v.(int))
}

// BlendEquation implements the gfx.ContextStateProvider interface.
func (c *Context) BlendEquation(eq gfx.BlendEquation) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(eq)],
		defaultValue: c.Object.Get("FUNC_ADD").Int(), // TODO(slimsag): verify
		key:          csBlendEquation,
		glCall:       c.glBlendEquation,
	}
}

func (c *Context) glDepthMask(v interface{}) {
	c.Object.Call("depthMask", v.(bool))
}

// DepthMask implements the gfx.ContextStateProvider interface.
func (c *Context) DepthMask(m bool) gfx.ContextStateValue {
	return csv{
		value:        m,
		defaultValue: true, // TODO(slimsag): verify
		key:          csDepthMask,
		glCall:       c.glDepthMask,
	}
}

func (c *Context) glViewport(v interface{}) {
	x := v.([4]int)
	c.Object.Call("viewport", x[0], x[0], x[0], x[0])
}

// Viewport implements the gfx.ContextStateProvider interface.
func (c *Context) Viewport(x, y, width, height int) gfx.ContextStateValue {
	return csv{
		value:        [4]int{x, y, width, height},
		defaultValue: [4]int{0, 0, 0, 0}, // TODO(slimsag): track real default viewport values
		key:          csViewport,
		glCall:       c.glViewport,
	}
}

func (c *Context) glScissor(v interface{}) {
	x := v.([4]int)
	c.Object.Call("scissor", x[0], x[0], x[0], x[0])
}

// Scissor implements the gfx.ContextStateProvider interface.
func (c *Context) Scissor(x, y, width, height int) gfx.ContextStateValue {
	return csv{
		value:        [4]int{x, y, width, height},
		defaultValue: [4]int{0, 0, 0, 0}, // TODO(slimsag): track real default scissor values
		key:          csScissor,
		glCall:       c.glScissor,
	}
}

func (c *Context) glLineWidth(v interface{}) {
	c.Object.Call("lineWidth", v.(float64))
}

// LineWidth implements the gfx.ContextStateProvider interface.
func (c *Context) LineWidth(w float32) gfx.ContextStateValue {
	return csv{
		value:        w,
		defaultValue: 1.0,
		key:          csLineWidth,
		glCall:       c.glLineWidth,
	}
}

func (c *Context) glColorMask(v interface{}) {
	x := v.([4]bool)
	c.Object.Call("colorMask", x[0], x[1], x[2], x[3])
}

// ColorMask implements the gfx.ContextStateProvider interface.
func (c *Context) ColorMask(r, g, b, a bool) gfx.ContextStateValue {
	return csv{
		value:        [4]bool{r, g, b, a},
		defaultValue: [4]bool{true, true, true, true},
		key:          csColorMask,
		glCall:       c.glColorMask,
	}
}

func (c *Context) glCullFace(v interface{}) {
	c.Object.Call("cullFace", v.(int))
}

// CullFace implements the gfx.ContextStateProvider interface.
func (c *Context) CullFace(f gfx.Facet) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(f)],
		defaultValue: c.Object.Get("FRONT").Int(), // TODO(slimsag): verify
		key:          csCullFace,
		glCall:       c.glCullFace,
	}
}

func (c *Context) glFrontFace(v interface{}) {
	c.Object.Call("frontFace", v.(int))
}

// FrontFace implements the gfx.ContextStateProvider interface.
func (c *Context) FrontFace(o gfx.Orientation) gfx.ContextStateValue {
	return csv{
		value:        c.Enums[int(o)],
		defaultValue: c.Object.Get("CCW").Int(), // TODO(slimsag): verify
		key:          csFrontFace,
		glCall:       c.glFrontFace,
	}
}

type featureKey struct {
	csKey int
	f     gfx.Feature
}

func (c *Context) glEnable(v interface{}) {
	c.Object.Call("enable", v.(int))
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
		glCall: c.glEnable,
	}
}

func (c *Context) glDisable(v interface{}) {
	c.Object.Call("disable", v.(int))
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
		glCall: c.glDisable,
	}
}
