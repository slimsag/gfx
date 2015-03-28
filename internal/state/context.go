package state

import "github.com/slimsag/gfx"

type CSV struct {
	Key                 interface{}
	Value, DefaultValue interface{}
	GLCall              func(value interface{})
}

type ContextState []gfx.ContextStateValue

func (c ContextState) Find(k interface{}) (index int, pair CSV) {
	var i interface{}
	for index, i = range c {
		pair = i.(CSV)
		if pair.Key != k {
			// Non-equal keys.
			continue
		}
		return index, pair
	}
	return -1, CSV{}
}

type Context struct {
	current ContextState
}

func (c *Context) NewState(values ...gfx.ContextStateValue) gfx.ContextState {
	return ContextState(values)
}

func (c *Context) Load(s gfx.ContextState) {
	var st ContextState
	if s != nil {
		st = s.(ContextState)
	}

	// For any state not explicitly mentioned in the current state, revert it
	// to the default state.
	for _, curI := range c.current {
		cur := curI.(CSV)
		if index, _ := st.Find(cur.Key); index != -1 {
			continue
		}

		// Revert to the default state.
		if cur.Value == cur.DefaultValue {
			// Already using this value! Do nothing.
			continue
		}
		cur.GLCall(cur.DefaultValue)
	}

	// For each state explicitly mentioned in the destination state, apply it
	// if needed.
	for _, dstI := range st {
		dst := dstI.(CSV)
		found, cur := c.current.Find(dst.Key)
		if found != -1 && cur.Value == dst.Value {
			// Already using this value! Do nothing.
			continue
		}

		// Did not find a matching previous value.
		dst.GLCall(dst.Value)
	}
	c.current = st
}
