package dax

import (
	"testing"

	"github.com/dlespiau/dax/math"
)

func TestLookAt(t *testing.T) {
	return

	tests := []struct {
		position, source, result math.Vec3
	}{
		{math.Vec3{0, 0, 0}, math.Vec3{0, 0, 1}, math.Vec3{0, 0, -1}},
		{math.Vec3{0, 0, 0}, math.Vec3{0, 0, -1}, math.Vec3{0, 0, 1}},
		{math.Vec3{0, 0, 1}, math.Vec3{0, 0, 1}, math.Vec3{0, 0, -1}},
		{math.Vec3{0, 0, 1}, math.Vec3{0, 0, -1}, math.Vec3{0, 0, 1}},
		{math.Vec3{0, 0, -1}, math.Vec3{0, 0, 1}, math.Vec3{0, 0, -1}},
		{math.Vec3{0, 0, -1}, math.Vec3{0, 0, -1}, math.Vec3{0, 0, 1}},

		{math.Vec3{10, 0, 0}, math.Vec3{1, 0, 0}, math.Vec3{0, 0, 1}},

		{math.Vec3{10, 0, 0}, math.Vec3{0, 1, 0}, math.Vec3{0, 1, 0}},
	}

	camera := NewPerspectiveCamera(90, 800.0/600, 1, 100)

	for _, test := range tests {
		camera.SetPositionV(&test.position)
		camera.LookAt(&math.Vec3{0, 0, 0})

		rotation := camera.GetRotation()
		result := rotation.Rotate(&test.source)
		assertVec3(t, &test.result, &result, 1e-3)
	}
}
