package main

import (
	dax "github.com/dlespiau/dax/lib"
)

var window *dax.Window

type Scene struct {
	dax.Scene

	poly             *dax.Polyline
	left_button_down bool
}

func (s *Scene) Setup() {
	s.SetBackgroundColor(0, 0, 0, 1)
	s.poly = dax.NewPolyline()
}

func (s *Scene) OnMouseMoved(x, y float32) {
	if s.left_button_down {
		s.poly.Add(x, y, 0)
	}
}

func (s *Scene) OnMouseButtonPressed(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.left_button_down = true
		s.poly.Add(x, y, 0)
	case dax.MouseButtonRight:
		s.poly.Clear()
	}
}

func (s *Scene) OnMouseButtonReleased(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.left_button_down = false
	}
}

func (s *Scene) Update() {
	p := s.poly.Positions()
	for i := 0; i < len(p); i += 3 {
		p[i] += dax.Rand(-.5, .5)
		p[i+1] += dax.Rand(-.5, .5)
	}
}

func (s *Scene) Draw(fb dax.Framebuffer) {
	fb.Draw(s.poly)
}

func main() {
	var app dax.Application

	window = app.CreateWindow("Polyline Example", 800, 600)
	window.SetScene(&Scene{})

	app.Run()
}
