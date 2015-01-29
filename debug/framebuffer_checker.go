// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// fbChecker is like the checker type, but for a gfx.Framebuffer. It implicitly
// invokes the Check method of the underlying context after each function call
// is made.
type fbChecker struct {
	fb  gfx.Framebuffer
	ctx gfx.Context
}

// ClearColor implements the gfx.Clearable interface.
func (f *fbChecker) ClearColor(r, g, b, a float32) {
	f.fb.ClearColor(r, g, b, a)
	f.ctx.Check()
}

// ClearDepth implements the gfx.Clearable interface.
func (f *fbChecker) ClearDepth(depth float64) {
	f.fb.ClearDepth(depth)
	f.ctx.Check()
}

// ClearStencil implements the gfx.Clearable interface.
func (f *fbChecker) ClearStencil(stencil int) {
	f.fb.ClearStencil(stencil)
	f.ctx.Check()
}

// Clear implements the gfx.Clearable interface.
func (f *fbChecker) Clear(m gfx.ClearMask) {
	f.fb.Clear(m)
	f.ctx.Check()
}
