package dax

import (
	m "github.com/dlespiau/dax/math"
)

type Polyline struct {
	vertices []float32
}

func NewPolyline() *Polyline {
	return NewPolylineWithSize(9)
}

func NewPolylineWithSize(n_vertices int) *Polyline {
	p := new(Polyline)
	p.vertices = make([]float32, 0, n_vertices*3)

	return p
}

func (p *Polyline) Size() int {
	return len(p.vertices) / 3
}

func (p *Polyline) Clear() {
	p.vertices = p.vertices[:0]
}

func (p *Polyline) Add(x, y, z float32) {
	p.vertices = append(p.vertices, x, y, z)
}

func (p *Polyline) AddVertex(v *m.Vec3) {
	p.vertices = append(p.vertices, v[0], v[1], v[2])
}

func (p *Polyline) AddPoint(point *m.Point) {
	p.vertices = append(p.vertices, point[0], point[1], 0)
}

func (p *Polyline) Positions() []float32 {
	return p.vertices
}
func (p *Polyline) draw(fb *Framebuffer) {
	fb.renderer.DrawPolyline(fb, p)
}
