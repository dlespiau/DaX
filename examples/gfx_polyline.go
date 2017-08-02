package main

import (
	dax "github.com/dlespiau/dax/lib"
)

var window *dax.Window

type shapePolyline struct {
	dax.Scene

	poly           *dax.Polyline
	leftButtonDown bool
}

func (s *shapePolyline) Setup() {
	s.SetBackgroundColor(0, 0, 0, 1)
	s.poly = dax.NewPolyline()
}

func (s *shapePolyline) OnMouseMoved(x, y float32) {
	if s.leftButtonDown {
		s.poly.Add(x, y, 0)
	}
}

func (s *shapePolyline) OnMouseButtonPressed(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.leftButtonDown = true
		s.poly.Add(x, y, 0)
	case dax.MouseButtonRight:
		s.poly.Clear()
	}
}

func (s *shapePolyline) OnMouseButtonReleased(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.leftButtonDown = false
	}
}

func (s *shapePolyline) Update(time float64) {
	p := s.poly.Positions()
	for i := 0; i < len(p); i += 3 {
		p[i] += dax.Rand(-.5, .5)
		p[i+1] += dax.Rand(-.5, .5)
	}
}

func (s *shapePolyline) Draw(fb dax.Framebuffer) {
	fb.Draw(s.poly)
}

var gfxPolylineExample = Example{
	Category:    CategoryGraphics,
	Name:        "Polyline",
	Description: "Draw a polyline with your mouse",
	Scene:       &shapePolyline{},
}
