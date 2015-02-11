// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// UniformLocation corresponds to the location of a uniform variable in a
// shader program.
//
// TODO(slimsag): consider simplification to e.g. integer
type UniformLocation interface {
}

// Shader represents which content and how that content is drawn to the render
// target.
type Shader interface {
	Object
}
