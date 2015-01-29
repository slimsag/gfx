package main

import (
	"runtime"

	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/debug"
	"github.com/slimsag/gfx/gl2"
	glfw "github.com/slimsag/glfw3"
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
