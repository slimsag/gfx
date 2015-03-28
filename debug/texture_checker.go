// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// textureChecker is like the checker type, but for a gfx.Texture. It
// implicitly invokes the Check method of the underlying context after each
// function call is made.
type textureChecker struct {
	t   gfx.Texture
	ctx gfx.Context
}

// Type implements the gfx.Texture interface.
func (t *textureChecker) Type() gfx.TextureType {
	typ := t.t.Type()
	t.ctx.Check()
	return typ
}

// Delete implements the gfx.Object interface.
func (t *textureChecker) Delete() {
	t.t.Delete()
	t.ctx.Check()
}

// Object implements the gfx.Object interface.
func (t *textureChecker) Object() interface{} {
	return t.t.Object()
}
