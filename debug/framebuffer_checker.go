// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// fbChecker is like the checker type, but for a gfx.Framebuffer. It implicitly
// invokes the Check method of the underlying context after each function call
// is made.
type fbChecker struct {
	fb  gfx.Framebuffer
	ctx gfx.Context
}

// NewState implements the gfx.FramebufferStateProvider interface.
func (f *fbChecker) NewState(values ...gfx.FramebufferStateValue) gfx.FramebufferState {
	return f.fb.NewState(values...)
}

// Load implements the gfx.FramebufferStateProvider interface.
func (f *fbChecker) Load(s gfx.FramebufferState) {
	f.fb.Load(s)
	f.check()
}

// ClearColor implements the gfx.FramebufferStateProvider interface.
func (f *fbChecker) ClearColor(r, g, b, a float32) gfx.FramebufferStateValue {
	return f.fb.ClearColor(r, g, b, a)
}

// ClearDepth implements the gfx.FramebufferStateProvider interface.
func (f *fbChecker) ClearDepth(depth float64) gfx.FramebufferStateValue {
	return f.fb.ClearDepth(depth)
}

// ClearStencil implements the gfx.FramebufferStateProvider interface.
func (f *fbChecker) ClearStencil(stencil int) gfx.FramebufferStateValue {
	return f.fb.ClearStencil(stencil)
}

// Clear implements the gfx.Framebuffer interface.
func (f *fbChecker) Clear(m gfx.ClearMask) {
	// Verify bitmask argument.
	if m == 0 {
		panic("Framebuffer.Clear: invalid clear mask argument (0)")
	}

	// Clearing all possible bits should yield zero.
	tmp := m
	tmp &^= gfx.ColorBuffer
	tmp &^= gfx.DepthBuffer
	tmp &^= gfx.StencilBuffer
	if tmp != 0 {
		panic("Framebuffer.Clear: invalid clear mask argument")
	}

	f.fb.Clear(m)
	f.check()
}

// ReadPixelsUint8 implements the gfx.Framebuffer interface.
func (f *fbChecker) ReadPixelsUint8(x, y, width, height int, dst []uint8) {
	// Verify destination buffer size.
	if len(dst) < width*height*4 {
		panic("Framebuffer.ReadPixelsUint8: dst buffer is not large enough")
	}

	f.fb.ReadPixelsUint8(x, y, width, height, dst)
	f.check()
}

// Texture2D implements the gfx.Framebuffer interface.
func (f *fbChecker) Texture2D(attachment gfx.FramebufferAttachment, target gfx.TextureTarget, tex gfx.Texture) {
	f.fb.Texture2D(attachment, target, tex)
	f.check()
}

// Renderbuffer implements the gfx.Framebuffer interface.
func (f *fbChecker) Renderbuffer(attachment gfx.FramebufferAttachment, buf gfx.Renderbuffer) {
	f.fb.Renderbuffer(attachment, buf)
	f.check()
}

// Status implements the gfx.Framebuffer interface.
func (f *fbChecker) Status() error {
	status := f.fb.Status()
	f.ctx.Check()
	return status
}

// Delete implements the gfx.Object interface.
func (f *fbChecker) Delete() {
	f.fb.Delete()
	f.ctx.Check() // not f.check() because it uses the deleted framebuffer.
}

// Object implements the gfx.Object interface.
func (f *fbChecker) Object() interface{} {
	return f.fb.Object()
}

func (f *fbChecker) check() {
	if status := f.Status(); status != nil {
		panic(status)
	}
}
