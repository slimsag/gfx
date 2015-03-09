// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"testing"
	//"fmt"
)

type csv struct {
	key                 string
	value, defaultValue interface{}
}

func shadows(on bool) ContextStateValue {
	return csv{"shadows", on, true}
}

func lighting(on bool) ContextStateValue {
	return csv{"lighting", on, true}
}

func physics(on bool) ContextStateValue {
	return csv{"physics", on, true}
}

type contextState []ContextStateValue

func (c contextState) find(k string) (index int, pair csv) {
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

type context struct {
	current contextState
}

func (c *context) NewState(values ...ContextStateValue) ContextState {
	return contextState(values)
}

func (c *context) Load(s ContextState) {
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
		c.glCall(cur.key, cur.defaultValue)
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
		c.glCall(dst.key, dst.value)
	}
	c.current = st
}

func (c *context) glCall(name string, newValue interface{}) {
	//fmt.Printf("glCall(%v, new=%v)\n", name, newValue)
}

func TestContextState(t *testing.T) {
	ctx := Context(&context{
		current: contextState{
			lighting(true),
			shadows(true),
			physics(true),
		},
	})

	// shadows: true (default).
	// lighting: false (explicit).
	// physics: true (default).
	ctx.Load(ctx.NewState(
		lighting(false),
	))

	// shadows: false (explicit).
	// lighting: explicit (false).
	// physics: true (default).
	ctx.Load(ctx.NewState(
		shadows(false),
		lighting(false),
	))

	// shadows: true (default).
	// lighting: true (default).
	// physics: true (default).
	ctx.Load(nil)
}

func BenchmarkContextState(b *testing.B) {
	ctx := Context(&context{
		current: contextState{
			lighting(true),
			shadows(true),
			physics(true),
		},
	})

	// shadows: true (default).
	// lighting: false (explicit).
	// physics: true (default).
	lightsOff := ctx.NewState(
		lighting(false),
	)

	// shadows: false (explicit).
	// lighting: explicit (false).
	// physics: true (default).
	lightsAndShadowsOff := ctx.NewState(
		shadows(false),
		lighting(false),
	)

	for i := 0; i < b.N; i++ {
		switch i % 3 {
		case 0:
			ctx.Load(lightsOff)
		case 1:
			ctx.Load(lightsAndShadowsOff)
		case 2:
			// shadows: true (default).
			// lighting: true (default).
			// physics: true (default).
			ctx.Load(nil)
		}
	}
}
