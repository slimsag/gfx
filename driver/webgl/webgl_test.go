package webgl

import "github.com/slimsag/gfx"

// Just for ensuring we meet the interface requirements.
func init() {
	_ = gfx.Buffer(&Buffer{})
	_ = gfx.Context(&Context{})
	_ = gfx.Framebuffer(Framebuffer{})
	_ = gfx.Program(&Program{})
	_ = gfx.Renderbuffer(&Renderbuffer{})
	_ = gfx.Shader(&Shader{})
	_ = gfx.Texture(&Texture{})
}
