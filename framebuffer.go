package dax

import (
	"github.com/go-gl/gl/v3.3-core/gl"

	m "github.com/dlespiau/dax/math"
)

type Framebuffer struct {
	renderer   *renderer
	projection m.Mat4
}

func NewFramebuffer() *Framebuffer {
	fb := new(Framebuffer)
	fb.renderer = newRenderer()
	return fb
}

type Drawable interface {
	draw(fb *Framebuffer)
}

func (fb *Framebuffer) Draw(d Drawable) {
	d.draw(fb)
}

func (fb *Framebuffer) SetViewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}
