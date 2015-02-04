// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package debug implements graphics application debugging.
package debug

import "github.com/slimsag/gfx"

// Context wraps the given graphics context with a Checker, it is short-handed
// for:
//
//  c = debug.Checker(c)
//
func Context(c gfx.Context) gfx.Context {
	return Checker(c)
}
