package dax

import (
	"testing"

	"github.com/dlespiau/dax/math"
)

func TestLookAt(t *testing.T) {

	tests := []struct {
		position, result math.Vec3
	}{
		{math.Vec3{1, 0, 0}, math.Vec3{-1, 0, 0}},
		{math.Vec3{-1, 0, 0}, math.Vec3{1, 0, 0}},
		// XXX: LookAt when on the "up" axis
		// {math.Vec3{0, 1, 0}, math.Vec3{0, -1, 0}},
		// {math.Vec3{0, -1, 0}, math.Vec3{0, 1, 0}},
		{math.Vec3{0, 0, 1}, math.Vec3{0, 0, -1}},
		{math.Vec3{0, 0, -1}, math.Vec3{0, 0, 1}},
	}

	camera := NewPerspectiveCamera(90, 800.0/600, 1, 100)

	for _, test := range tests {
		t.Logf("test: position %v", test.position)

		camera.SetPositionV(&test.position)
		camera.LookAt(&math.Vec3{0, 0, 0})

		transform := camera.AsNode().GetTransform()
		result4 := transform.Mul4x1(&math.Vec4{0, 0, -1})
		result := result4.Vec3()
		assertVec3(t, &test.result, &result, 1e-3)
	}
}
