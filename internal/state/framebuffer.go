package state

import "github.com/slimsag/gfx"

type FramebufferState []gfx.FramebufferStateValue

func (c FramebufferState) Find(k interface{}) (index int, pair CSV) {
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

type Framebuffer struct {
	current FramebufferState
	Loaded  gfx.FramebufferState
}

func (f *Framebuffer) NewState(values ...gfx.FramebufferStateValue) gfx.FramebufferState {
	return FramebufferState(values)
}

func (f *Framebuffer) Load(s gfx.FramebufferState) {
	f.Loaded = s
}

func (f *Framebuffer) GLCall(ld gfx.FramebufferState) {
	s := ld
	var st FramebufferState
	if s != nil {
		st = s.(FramebufferState)
	}

	// For any state not explicitly mentioned in the current state, revert it
	// to the default state.
	for _, curI := range f.current {
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
		found, cur := f.current.Find(dst.Key)
		if found != -1 && cur.Value == dst.Value {
			// Already using this value! Do nothing.
			continue
		}

		// Did not find a matching previous value.
		dst.GLCall(dst.Value)
	}
	f.current = st
}
