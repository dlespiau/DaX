package dax

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"

	m "github.com/dlespiau/dax/math"
)

type Framebuffer interface {
	Size() (width, height int)
	SetSize(width, height int)

	SetViewport(x, y, width, height int)
	Projection() *m.Mat4
	SetProjection(projection *m.Mat4)

	Draw(d Drawable)

	Screenshot() *image.RGBA

	// private
	render() *renderer
}

type Drawable interface {
	draw(fb Framebuffer)
}

type OnScreen struct {
	renderer      *renderer
	width, height int
	projection    m.Mat4
}

func newOnScreen(width, height int) *OnScreen {
	fb := new(OnScreen)
	fb.renderer = newRenderer()
	return fb
}

func (fb *OnScreen) Size() (width, height int) {
	return fb.width, fb.height
}

func (fb *OnScreen) Projection() *m.Mat4 {
	return &fb.projection
}

func (fb *OnScreen) SetProjection(projection *m.Mat4) {
	fb.projection = *projection
}

func (fb *OnScreen) render() *renderer {
	return fb.renderer
}

func (fb *OnScreen) Draw(d Drawable) {
	d.draw(fb)
}

func (fb *OnScreen) SetSize(width, height int) {
	fb.width = width
	fb.height = height
}

func (fb *OnScreen) SetViewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (fb *OnScreen) Screenshot() *image.RGBA {
	pixels := make([]byte, fb.width*fb.height*4)

	gl.ReadPixels(0, 0, int32(fb.width), int32(fb.height), gl.RGBA,
		gl.UNSIGNED_BYTE, unsafe.Pointer(&pixels[0]))

	return &image.RGBA{pixels, fb.width * 4,
		image.Rect(0, 0, fb.width, fb.height)}
}
