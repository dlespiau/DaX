package dax

import (
	"github.com/go-gl/gl/v3.3-core/gl"

	m "github.com/dlespiau/dax/math"
)

type Framebuffer struct {
	renderer      *renderer
	width, height int
	projection    m.Mat4
}

func NewFramebuffer(width, height int) *Framebuffer {
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

func (fb *Framebuffer) setSize(width, height int) {
	fb.width = width
	fb.height = height
}

func (fb *Framebuffer) SetViewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}
