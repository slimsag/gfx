package gfx

// ClearMask represents a bitmask to choose which buffers to clear during a
// framebuffer clearing operation. It must be one of the predefined constants.
type ClearMask int

// Clearing masks to select the color, depth, and stencil buffers. They can be
// bitwise OR'd together to select multiple:
//
//  colorAndDepth := ColorBufer|DepthBuffer
//
const (
	ColorBuffer ClearMask = 1 << iota
	DepthBuffer
	StencilBuffer
)

// Clearable represents the clearing state and API of a framebuffer object.
type Clearable interface {
	// ClearColor sets the color to clear the color buffer to upon a call to
	// the Clear method.
	ClearColor(r, g, b, a float32)

	// ClearDepth sets the value to clear the depth buffer to upon a depth
	// buffer clearing operation (a call to Clear with the DepthBuffer clear
	// mask)
	ClearDepth(depth float32)

	// ClearStencil sets the value to clear the stencil buffer to upon a
	// stencil buffer clearing operation (a call to Clear with the
	// StencilBuffer clear mask)
	ClearStencil(stencil int)

	// Clear clears the buffers selected by the bitmask to their respective
	// clear values. Multiple bitmasks can be OR'd together to select multiple
	// buffers to clear at once:
	//
	//  // Clear both(!) the color and depth buffers in one call.
	//  Clear(ColorBuffer|DepthBuffer)
	//
	Clear(m ClearMask)
}
