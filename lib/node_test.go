package dax

import (
	"testing"

	"github.com/dlespiau/dax/math"

	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	n := NewNode()
	n.SetPosition(1, 2, 3)
	assertVec3(t, &math.Vec3{1, 2, 3}, n.GetPosition(), 1e-6)
}

func TestPositionV(t *testing.T) {
	n := NewNode()
	v := &math.Vec3{1, 2, 3}
	n.SetPositionV(v)
	assertVec3(t, v, n.GetPosition(), 1e-6)
}

func TestTranslate(t *testing.T) {
	n := NewNode()
	n.TranslateX(10)
	assertVec3(t, &math.Vec3{10, 0, 0}, n.GetPosition(), 1e-6)

	n.TranslateY(14)
	assertVec3(t, &math.Vec3{10, 14, 0}, n.GetPosition(), 1e-6)

	n.TranslateZ(21)
	assertVec3(t, &math.Vec3{10, 14, 21}, n.GetPosition(), 1e-6)
}

func TestRotation(t *testing.T) {
	n := NewNode()
	q := math.QuatRotate(math.Pi/2, &math.Vec3{1, 1, 1})
	n.SetRotation(&q)
	assertQuat(t, &q, n.GetRotation(), 1e-6)
}

func TestRotate(t *testing.T) {
	n := NewNode()
	n.RotateX(math.Pi / 2)
	q := math.QuatRotate(math.Pi/2, &math.Vec3{1, 0, 0})
	assertQuat(t, &q, n.GetRotation(), 1e-6)

	n = NewNode()
	n.RotateY(math.Pi / 2)
	q = math.QuatRotate(math.Pi/2, &math.Vec3{0, 1, 0})
	assertQuat(t, &q, n.GetRotation(), 1e-6)

	n = NewNode()
	n.RotateZ(math.Pi / 2)
	q = math.QuatRotate(math.Pi/2, &math.Vec3{0, 0, 1})
	assertQuat(t, &q, n.GetRotation(), 1e-6)
}

func TestSetScale(t *testing.T) {
	n := NewNode()
	assertVec3(t, &math.Vec3{1, 1, 1}, n.GetScale(), 1e-6)

	n.SetScale(1, 2, 3)
	assertVec3(t, &math.Vec3{1, 2, 3}, n.GetScale(), 1e-6)

	n.SetScaleV(&math.Vec3{4, 5, 6})
	assertVec3(t, &math.Vec3{4, 5, 6}, n.GetScale(), 1e-6)
}

func TestScale(t *testing.T) {
	n := NewNode()
	n.ScaleX(2)
	assertVec3(t, &math.Vec3{2, 1, 1}, n.GetScale(), 1e-6)

	n.ScaleY(2)
	assertVec3(t, &math.Vec3{2, 2, 1}, n.GetScale(), 1e-6)

	n.ScaleZ(2)
	assertVec3(t, &math.Vec3{2, 2, 2}, n.GetScale(), 1e-6)
}

func TestAddChild(t *testing.T) {
	p := NewNode()
	c := NewNode()

	p.AddChild(c)
	assert.Equal(t, c.GetParent(), p)

	children := p.GetChildren()
	assert.Equal(t, len(children), 1)
	assert.Equal(t, children[0], c)
}
