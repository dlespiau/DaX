package main

import (
	dax "github.com/dlespiau/dax/lib"
)

var window *dax.Window

type Scene struct {
	dax.Scene

	poly           *dax.Polyline
	leftButtonDown bool
}

func (s *Scene) Setup() {
	s.SetBackgroundColor(0, 0, 0, 1)
	s.poly = dax.NewPolyline()
}

func (s *Scene) OnMouseMoved(x, y float32) {
	if s.leftButtonDown {
		s.poly.Add(x, y, 0)
	}
}

func (s *Scene) OnMouseButtonPressed(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.leftButtonDown = true
		s.poly.Add(x, y, 0)
	case dax.MouseButtonRight:
		s.poly.Clear()
	}
}

func (s *Scene) OnMouseButtonReleased(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.leftButtonDown = false
	}
}

func (s *Scene) Update(time float64) {
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
	app := dax.NewApplication("Polyline Example")

	window := app.CreateWindow(app.Name, 800, 600)
	window.SetScene(&Scene{})

	app.Run()
}
