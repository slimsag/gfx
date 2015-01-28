package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/slimsag/gfx"
	"github.com/slimsag/gfx/debug"
	"github.com/slimsag/gfx/webgl"
)

func main() {
	// Create a canvas element.
	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	document.Get("body").Call("appendChild", canvas)

	// Set size of canvas to 640x480 
	canvas.Set("width", 640)
	canvas.Set("height", 480)

	// Create a new WebGL context (could also webgl.Wrap an existing one).
	ctx, err := webgl.New(canvas, nil)
	if err != nil {
		js.Global.Call("alert", err.Error())
	}

	// Wrap the context to get a debug context.
	ctx = debug.Context(ctx)

	// Clear the color buffer to red.
	ctx.Clear(gfx.ClearColor{1, 0, 0, 1})
}
