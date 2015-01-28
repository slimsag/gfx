package debug

import "github.com/slimsag/gfx"

// Checker is a gfx.Context that implicitly invokes the Check method of the
// underlying context after each function call made. Thus, if any error should
// occur you will receive a nice stack trace where that error occured.
//
// It should not be used in production builds, as it has a performance
// overhead.
type Checker struct {
	gfx.Context
}

// Check implements the gfx.Context interface.
func (c Checker) Check() {
	// We don't want caller to accidently grab the error, so we stub out the
	// call here.
	return
}

// Context wraps the given graphics context with a Checker, it is short-handed
// for:
//
//  c = debug.Checker{
//      Context: c,
//  }
//
func Context(c gfx.Context) gfx.Context {
	return Checker{
		Context: c,
	}
}
