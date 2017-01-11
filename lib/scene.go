package dax

import (
	"fmt"
	"reflect"
)

type Scener interface {
	Setup()
	TearDown()

	// XXX: Shouldn't probably be part of the interface, but needed to
	// generically clear the framebuffer
	BackgroundColor() *Color

	Update(time float64)
	Draw(fb Framebuffer)

	// events
	OnResize(fb Framebuffer, width, height int)
	OnKeyPressed()
	OnKeyReleased()
	OnMouseMoved(x, y float32)
	OnMouseButtonPressed(button MouseButton, x, y float32)
	OnMouseButtonReleased(button MouseButton, x, y float32)
	OnRuneEntered(r rune)
}

type Scene struct {
	camera          Camera
	name            string
	backgroundColor Color
}

func (s *Scene) Setup() {
}

func toScene(s Scener) *Scene {
	if scene, ok := s.(*Scene); ok {
		return scene
	}

	v := reflect.ValueOf(s).Elem().FieldByName("Scene")
	return v.Addr().Interface().(*Scene)
}

func sceneSetup(s Scener, fb Framebuffer) {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tag := f.Tag.Get("dax")
		if tag == "" {
			continue
		}

		if tag == "property" {
			fmt.Println(f.Name)
		}

	}

	s.Setup()

	if scene := toScene(s); scene != nil && scene.camera == nil {
		// If we don't have a camera by that point, we default to an
		// orthographic one placing (0, 0) at the top left corner and
		// making (width - 1, height - 1) the bottom right corner
		width, height := fb.Size()
		scene.camera = NewScreenSpaceCamera(width, height, -1, 1)
	}
}

func (s *Scene) TearDown() {
}

func (s *Scene) BackgroundColor() *Color {
	return &s.backgroundColor
}

func (s *Scene) SetBackgroundColor(r, g, b, a float32) {
	s.backgroundColor.R = r
	s.backgroundColor.G = g
	s.backgroundColor.B = b
	s.backgroundColor.A = a
}

func (s *Scene) SetCamera(camera Camera) {
	if camera == nil {
		return
	}

	s.camera = camera
}

func sceneUpdate(s Scener, time float64) {
	s.Update(time)
}

func (s *Scene) Update(time float64) {
}

func (s *Scene) Draw(fb Framebuffer) {
}

func (s *Scene) OnResize(fb Framebuffer, width, height int) {
	fb.SetSize(width, height)
	fb.SetViewport(0, 0, width, height)

	s.camera.UpdateFBSize(width, height)
	projection := s.camera.GetProjection()
	fb.SetProjection(projection)
}

func (s *Scene) OnKeyPressed() {
}

func (s *Scene) OnKeyReleased() {
}

func (s *Scene) OnMouseMoved(x, y float32) {
}

func (s *Scene) OnMouseButtonPressed(button MouseButton, x, y float32) {
}

func (s *Scene) OnMouseButtonReleased(button MouseButton, x, y float32) {
}

func (s *Scene) OnRuneEntered(r rune) {
}
