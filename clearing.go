package gfx

// ClearColor is used by a Clearable framebuffer object to clear the color
// buffer to a specific [R, G, B, A] color.
//
//  // Clear color buffer to it's initial color:
//  Clear(gfx.ClearColor{0, 0, 0, 0})
//
//  // Equivilent OpenGL calls:
//  glClearColor(0, 0, 0, 0)
//  glClear(GL_COLOR_BUFFER_BIT)
//
// All values are clamped to a range of [0, 1].
type ClearColor [4]float32

// ClearDepth is used by a Clearable framebuffer object to clear the depth
// buffer to a specific stencil value.
//
//  // Clear depth buffer to it's initial value:
//  Clear(gfx.ClearDepth(1))
//
//  // Equivilent OpenGL calls:
//  glClearDepth(1)
//  glClear(GL_DEPTH_BUFFER_BIT)
//
// All values are clamped to a range of [0, 1].
type ClearDepth float32

// ClearStencil is used by a Clearable framebuffer object to clear the stencil
// buffer to a specific stencil value.
//
//  // Clear stencil buffer to it's initial value:
//  Clear(gfx.ClearStencil(0))
//
//  // Equivilent OpenGL calls:
//  glClearStencil(0)
//  glClear(GL_STENCIL_BUFFER_BIT)
//
// The value is masked with 2m-1, where m is the number of bits in the stencil
// buffer.
type ClearStencil int
