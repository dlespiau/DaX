package dax

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Framebuffer interface {
	Size() (width, height int)
	SetSize(width, height int)

	GetCamera() Camera
	SetCamera(camera Camera)
	SetViewport(x, y, width, height int)

	Draw(d Drawer)

	Screenshot() *image.RGBA

	// private
	render() *renderer
}

type onScreen struct {
	renderer      *renderer
	width, height int
	camera        Camera
}

func newOnScreen(width, height int) *onScreen {
	fb := new(onScreen)
	fb.renderer = newRenderer()
	return fb
}

func (fb *onScreen) Size() (width, height int) {
	return fb.width, fb.height
}

func (fb *onScreen) GetCamera() Camera {
	return fb.camera
}

func (fb *onScreen) SetCamera(camera Camera) {
	fb.camera = camera
}

func (fb *onScreen) render() *renderer {
	return fb.renderer
}

func (fb *onScreen) Draw(d Drawer) {
	d.Draw(fb)
}

func (fb *onScreen) SetSize(width, height int) {
	fb.width = width
	fb.height = height
}

func (fb *onScreen) SetViewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (fb *onScreen) Screenshot() *image.RGBA {
	pixels := make([]byte, fb.width*fb.height*4)

	gl.ReadPixels(0, 0, int32(fb.width), int32(fb.height), gl.RGBA,
		gl.UNSIGNED_BYTE, unsafe.Pointer(&pixels[0]))

	return &image.RGBA{pixels, fb.width * 4,
		image.Rect(0, 0, fb.width, fb.height)}
}
