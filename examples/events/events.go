package main

import (
	"fmt"
	"github.com/dlespiau/dax"
)

type Scene struct {
	dax.Scene
}

func (s *Scene) OnMouseMoved(x, y float32) {
	fmt.Printf("Mouse moved at (%.0f, %.0f)\n", x, y)
}

func (s *Scene) OnMouseButtonPressed(b dax.MouseButton, x, y float32) {
	fmt.Printf("Button %v pressed at (%.0f, %.0f)\n", b, x, y)
}

func (s *Scene) OnMouseButtonReleased(b dax.MouseButton, x, y float32) {
	fmt.Printf("Button %v released at (%.0f, %.0f)\n", b, x, y)
}

func (s *Scene) OnRuneEntered(r rune) {
	fmt.Printf("Rune entered '%c'\n", r)
}

func main() {
	var app dax.Application

	window := app.CreateWindow("Events Example", 800, 600)
	window.SetScene(&Scene{})

	app.Run()
}
