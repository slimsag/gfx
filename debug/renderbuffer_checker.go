// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// rbChecker is like the checker type, but for a gfx.Renderbuffer. It
// implicitly invokes the Check method of the underlying context after each
// function call is made.
type rbChecker struct {
	rb  gfx.Renderbuffer
	ctx gfx.Context
}

// Storage implements the gfx.Renderbuffer interface.
func (r *rbChecker) Storage(internalFormat gfx.RenderbufferFormat, width, height int) {
	r.rb.Storage(internalFormat, width, height)
	r.ctx.Check()
}

// Delete implements the gfx.Object interface.
func (r *rbChecker) Delete() {
	r.rb.Delete()
	r.ctx.Check()
}

// Object implements the gfx.Object interface.
func (r *rbChecker) Object() interface{} {
	return r.rb.Object()
}
