package gfx

// ClearColor is used by a Clearable framebuffer object to clear the color
// buffer to a specific [R, G, B, A] color.
//
//  // Clear color buffer to red:
//  Clear(gfx.ClearColor{1, 0, 0, 1})
//
//  // Equivilent OpenGL calls:
//  glClearColor(1, 0, 0, 1)
//  glClear(GL_COLOR_BUFFER_BIT)
//
type ClearColor [4]float32

// ClearDepth is used by a Clearable framebuffer object to clear the depth
// buffer to a specific stencil value.
//
//  // Clear depth buffer to zero:
//  Clear(gfx.ClearDepth(0))
//
//  // Equivilent OpenGL calls:
//  glClearDepth(0)
//  glClear(GL_DEPTH_BUFFER_BIT)
//
type ClearDepth float32

// ClearStencil is used by a Clearable framebuffer object to clear the stencil
// buffer to a specific stencil value.
//
//  // Clear stencil buffer to zero:
//  Clear(gfx.ClearStencil(0))
//
//  // Equivilent OpenGL calls:
//  glClearStencil(0)
//  glClear(GL_STENCIL_BUFFER_BIT)
//
type ClearStencil int
