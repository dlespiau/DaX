package dax

import (
	"github.com/dlespiau/dax/math"
)

type Camera interface {
	UpdateFBSize(width, height int)
	GetProjection() *math.Mat4
}

type BaseCamera struct {
	Node
	projection math.Mat4
}

func (c *BaseCamera) GetProjection() *math.Mat4 {
	return &c.projection
}

type orthographicCamera struct {
	BaseCamera
}

func NewOrthographicCamera(left, right, bottom, top, near, far float32) *orthographicCamera {
	c := new(orthographicCamera)

	c.projection = math.Ortho(left, right, bottom, top, near, far)
	return c
}

func (c *orthographicCamera) UpdateFBSize(width, height int) {
}

type perspectiveCamera struct {
	BaseCamera
	fovy, aspect, near, far float32
}

func (c *perspectiveCamera) updateProjection() {
	c.projection = math.Perspective(c.fovy, c.aspect, c.near, c.far)
}

func NewPerspectiveCamera(fovy, aspect, near, far float32) *perspectiveCamera {
	c := new(perspectiveCamera)
	c.fovy = fovy
	c.aspect = aspect
	c.near = near
	c.far = far

	c.updateProjection()

	return c
}

func (c *perspectiveCamera) UpdateFBSize(width, height int) {
	c.aspect = float32(width) / float32(height)
	c.updateProjection()
}
