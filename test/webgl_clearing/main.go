// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build js

package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/debug"
	"github.com/slimsag/gfx/driver/webgl"
)

func main() {
	// Create the body and canvas.
	document := js.Global.Get("document")
	body := document.Call("createElement", "body")
	document.Set("body", body)
	canvas := document.Call("createElement", "canvas")
	body.Call("appendChild", canvas)

	// Set size of canvas to 640x480
	canvas.Set("width", 640)
	canvas.Set("height", 480)

	// Create a new WebGL context (could also webgl.Wrap an existing one).
	gl, err := webgl.New(canvas, nil)
	if err != nil {
		js.Global.Call("alert", err.Error())
	}

	// Wrap the context to get a debug context.
	gl = debug.Context(gl)

	fb := gl.Framebuffer()
	fb.Load(fb.NewState(
		fb.ClearColor(1, 0, 0, 1),
	))

	// Clear the color buffer to red.
	fb.Clear(gfx.ColorBuffer)
}
