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

	Update()
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

func sceneSetup(s Scener) {
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
	s.camera = camera
}

func (s *Scene) Update() {
}

func (s *Scene) Draw(fb Framebuffer) {
}

func (s *Scene) OnResize(fb Framebuffer, width, height int) {
	fb.SetSize(width, height)
	fb.SetViewport(0, 0, width, height)

	var camera Camera
	if s.camera != nil {
		camera = s.camera
	} else {
		camera = newOrthographicCamera(0, float32(width),
			float32(height), 0, -1, 1)
	}
	projection := camera.GetProjection()
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
