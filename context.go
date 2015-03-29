// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import "errors"

// Panics that Context.Check may generate.
var (
	OutOfMemory                 = errors.New("out of memory")
	InvalidEnum                 = errors.New("invalid enum")
	InvalidOperation            = errors.New("invalid operation")
	InvalidFramebufferOperation = errors.New("invalid framebuffer operation")
	InvalidValue                = errors.New("invalid value")
	StackOverflow               = errors.New("stack overflow")
	StackUnderflow              = errors.New("stack underflow")
	ContextLost                 = errors.New("context lost")
)

// Context is a graphics context. Unlike traditional OpenGL contexts, it is not
// tied to a specific OS thread (but still must be accessed from only one
// goroutine/thread at a time).
type Context interface {
	ContextStateProvider

	// Framebuffer returns the default framebuffer of the context (i.e. the one
	// that when drawn to appears on the window).
	Framebuffer() Framebuffer

	// NewFramebuffer returns a new Framebuffer object used for storing
	// complete frames.
	NewFramebuffer() Framebuffer

	// NewRenderbuffer returns a new Renderbuffer object used for storing
	// complete frames.
	NewRenderbuffer() Renderbuffer

	// NewShader returns a new Shader object which represents the content and
	// how exactly it is drawn to a render target.
	NewShader(t ShaderType) Shader

	// NewTexture returns a new Texture object which is used for images and
	// cube maps when rendering shapes.
	NewTexture(t TextureType) Texture

	// NewBuffer returns a new buffer object that contains data such as
	// vertices or colors.
	NewBuffer(t BufferType) Buffer

	// NewProgram returns a new program object which represents the
	// programmable OpenGL pipeline and associated shader programs.
	NewProgram() Program

	// Check checks that no errors have occured in the context. If an error
	// occurs, it is either a programmer error (passing an invalid value, etc)
	// or a serious device error (running out of memory, losing the context).
	Check()

	// Flush flushes any buffered commands out to the graphics hardware as
	// quickly as possible. Execution may not be completed in any particular
	// time period, but does complete in finite time.
	//
	// Drivers may buffer commands and queue them for sending to hardware in
	// large chunks. For this reason, all programs should call Flush whenever
	// they count on having all of the previously issued commands completed.
	Flush()

	// Finish blocks (does not return) until the effects of all previously
	// called commands are complete. Such effects include all changes to the
	// graphics state, all changes to connection state, and all changes to the
	// framebuffer contents.
	//
	// In most cases, you shouldn't ever use Finish, but rather Flush.
	Finish()
}
