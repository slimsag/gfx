// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/debug"
	"github.com/slimsag/gfx/driver/gl2"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	// Initialize OpenGL 2 graphics context.
	window.MakeContextCurrent()
	gl, err := gl2.New()
	if err != nil {
		panic(err)
	}

	// Wrap the context to get a debug context.
	gl = debug.Context(gl)

	for !window.ShouldClose() {
		// Clear the color buffer to red.
		gl.ClearColor(1, 0, 0, 1)
		gl.Clear(gfx.ColorBuffer)

		// Swap the front and back buffers, poll for events.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
