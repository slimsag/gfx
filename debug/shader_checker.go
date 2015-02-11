// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// shaderChecker is like the checker type, but for a gfx.Shader. It implicitly
// invokes the Check method of the underlying context after each function call
// is made.
type shaderChecker struct {
	s   gfx.Shader
	ctx gfx.Context
}