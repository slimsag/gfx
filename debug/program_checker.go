// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// programChecker is like the checker type, but for a gfx.Program. It
// implicitly invokes the Check method of the underlying context after each
// function call is made.
type programChecker struct {
	p   gfx.Program
	ctx gfx.Context
}

// Delete implements the gfx.Object interface.
func (p *programChecker) Delete() {
	p.p.Delete()
	p.ctx.Check()
}
