package dax

import (
	"github.com/dlespiau/dax/math"
)

type Camera interface {
	GetProjection() *math.Mat4
}

type BaseCamera struct {
	Node
	projection math.Mat4
}

func (camera *BaseCamera) GetProjection() *math.Mat4 {
	return &camera.projection
}

type orthographicCamera struct {
	BaseCamera
}

func NewOrthographicCamera(left, right, bottom, top, near, far float32) *orthographicCamera {
	camera := new(orthographicCamera)

	camera.projection = math.Ortho(left, right, bottom, top, near, far)
	return camera
}

type perspectiveCamera struct {
	BaseCamera
}

func NewPerspectiveCamera(fovy, aspect, near, far float32) *perspectiveCamera {
	camera := new(perspectiveCamera)

	camera.projection = math.Perspective(fovy, aspect, near, far)
	return camera
}
