// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import "github.com/slimsag/gfx"

// programChecker is like the checker type, but for a gfx.Program. It
// implicitly invokes the Check method of the underlying context after each
// function call is made.
type programChecker struct {
	p   gfx.Program
	ctx gfx.Context
}

// Link implements the gfx.Program interface.
func (p *programChecker) Link(vert, frag gfx.Shader) bool {
	success := p.p.Link(vert, frag)
	p.ctx.Check()
	return success
}

// InfoLog implements the gfx.Program interface.
func (p *programChecker) InfoLog() string {
	infoLog := p.p.InfoLog()
	p.ctx.Check()
	return infoLog
}

// AttribLocation implements the gfx.Program interface.
func (p *programChecker) AttribLocation(name string) gfx.AttribLocation {
	l := p.p.AttribLocation(name)
	p.ctx.Check()
	return l
}

// UniformLocation implements the gfx.Program interface.
func (p *programChecker) UniformLocation(name string) gfx.UniformLocation {
	l := p.p.UniformLocation(name)
	p.ctx.Check()
	return l
}

// Uniform1fv implements the gfx.Program interface.
func (p *programChecker) Uniform1fv(l gfx.UniformLocation, data []float32) {
	p.p.Uniform1fv(l, data)
	p.ctx.Check()
}

// Uniform1iv implements the gfx.Program interface.
func (p *programChecker) Uniform1iv(l gfx.UniformLocation, data []int32) {
	p.p.Uniform1iv(l, data)
	p.ctx.Check()
}

// Uniform2fv implements the gfx.Program interface.
func (p *programChecker) Uniform2fv(l gfx.UniformLocation, data []float32) {
	p.p.Uniform2fv(l, data)
	p.ctx.Check()
}

// Uniform2iv implements the gfx.Program interface.
func (p *programChecker) Uniform2iv(l gfx.UniformLocation, data []int32) {
	p.p.Uniform2iv(l, data)
	p.ctx.Check()
}

// Uniform3fv implements the gfx.Program interface.
func (p *programChecker) Uniform3fv(l gfx.UniformLocation, data []float32) {
	p.p.Uniform3fv(l, data)
	p.ctx.Check()
}

// Uniform3iv implements the gfx.Program interface.
func (p *programChecker) Uniform3iv(l gfx.UniformLocation, data []int32) {
	p.p.Uniform3iv(l, data)
	p.ctx.Check()
}

// Uniform4fv implements the gfx.Program interface.
func (p *programChecker) Uniform4fv(l gfx.UniformLocation, data []float32) {
	p.p.Uniform4fv(l, data)
	p.ctx.Check()
}

// Uniform4iv implements the gfx.Program interface.
func (p *programChecker) Uniform4iv(l gfx.UniformLocation, data []int32) {
	p.p.Uniform4iv(l, data)
	p.ctx.Check()
}

// UniformMatrix2fv implements the gfx.Program interface.
func (p *programChecker) UniformMatrix2fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.p.UniformMatrix2fv(l, transpose, data)
	p.ctx.Check()
}

// UniformMatrix3fv implements the gfx.Program interface.
func (p *programChecker) UniformMatrix3fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.p.UniformMatrix3fv(l, transpose, data)
	p.ctx.Check()
}

// UniformMatrix4fv implements the gfx.Program interface.
func (p *programChecker) UniformMatrix4fv(l gfx.UniformLocation, transpose bool, data []float32) {
	p.p.UniformMatrix4fv(l, transpose, data)
	p.ctx.Check()
}

// Delete implements the gfx.Object interface.
func (p *programChecker) Delete() {
	p.p.Delete()
	p.ctx.Check()
}

// Object implements the gfx.Object interface.
func (p *programChecker) Object() interface{} {
	return p.p.Object()
}
