// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// bufferChecker is like the checker type, but for a gfx.Buffer. It implicitly
// invokes the Check method of the underlying context after each function call
// is made.
type bufferChecker struct {
	b   gfx.Buffer
	ctx gfx.Context
}

// Delete implements the gfx.Object interface.
func (b *bufferChecker) Delete() {
	b.b.Delete()
	b.ctx.Check()
}

// Object implements the gfx.Object interface.
func (b *bufferChecker) Object() interface{} {
	return b.b.Object()
}
