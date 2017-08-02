package main

import (
	"github.com/dlespiau/dax"
)

type miscProperties struct {
	dax.Scene

	poly           *dax.Polyline
	leftButtonDown bool
	dx             float32 `dax:"property"`
	dy             float32 `dax:"property"`
}

func (s *miscProperties) Setup() {
	s.poly = dax.NewPolyline()
	s.dx = .5
	s.dy = .5
}

func (s *miscProperties) OnMouseMoved(x, y float32) {
	if s.leftButtonDown {
		s.poly.Add(x, y, 0)
	}
}

func (s *miscProperties) OnMouseButtonPressed(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.leftButtonDown = true
		s.poly.Add(x, y, 0)
	case dax.MouseButtonRight:
		s.poly.Clear()
	}
}

func (s *miscProperties) OnMouseButtonReleased(b dax.MouseButton, x, y float32) {
	switch b {
	case dax.MouseButtonLeft:
		s.leftButtonDown = false
	}
}

func (s *miscProperties) Update(time float64) {
	p := s.poly.Positions()
	for i := 0; i < len(p); i += 3 {
		p[i] += dax.Rand(-s.dx, s.dx)
		p[i+1] += dax.Rand(-s.dy, s.dy)
	}
}

func (s *miscProperties) Draw(fb dax.Framebuffer) {
	fb.Draw(s.poly)
}

var miscPropertiesExample = Example{
	Name:        "misc properties",
	Description: "Draw a polyline with your mouse",
	Scene:       &miscProperties{},
}
