// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// checker is a gfx.Context that implicitly invokes the Check method of the
// underlying context after each function call is made. Thus, if any error
// should occur you will receive a nice stack trace where that error occured.
type checker struct {
	*fbChecker
	ctx gfx.Context
}

// Check implements the gfx.Context interface.
func (c *checker) Check() {
	// We don't want caller to accidently grab the error, so we stub out the
	// call here.
	return
}

// Checker wraps the given graphics context such that each function call to the
// context (or any object gotten from it, e.g. a Framebuffer) has an implicit
// Check() call after it.
//
// This ensures that, should any error occur in the context, you will receive
// a nice Go stack trace with the exact function where the error was made.
func Checker(c gfx.Context) gfx.Context {
	return &checker{
		fbChecker: &fbChecker{
			fb:  c.(gfx.Framebuffer),
			ctx: c,
		},
		ctx: c,
	}
}
