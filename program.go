package gfx

// Program represents the programmable OpenGL pipeline and associated shader
// programs.
type Program interface {
	// AttachShader attaches the given shader object to this program.
	//
	// If the shader is already attached to a program, or another shader of the
	// same type is already attached to this program, ErrInvalidOperation is
	// generated (see Context.Error).
	AttachShader(s Shader) error
}
