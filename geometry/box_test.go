package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPlane(t *testing.T) {
	ctx := &boxContext{}

	buildPlane(ctx, 2, 1, 0, -1, -1, 2, 2, 2, 2, 2)

	assert.Equal(t, 9*3, len(ctx.positions))
	assert.Equal(t, 9*3, len(ctx.normals))
	assert.Equal(t, 9*2, len(ctx.uvs))
}
func TestBoxMesh(t *testing.T) {
	box := &Box{
		Width:             10,
		Height:            20,
		Depth:             30,
		NumWidthSegments:  2,
		NumHeightSegments: 2,
		NumDepthSegments:  2,
	}

	valid := box.Validate()
	assert.True(t, valid)

	m := box.GetMesh()

	positions := m.GetAttribute("position")
	assert.NotNil(t, positions)
	// 9 vertices, 6 faces, 3 components per vertex
	assert.Equal(t, 9*6*3, len(positions.Data))

	normals := m.GetAttribute("normal")
	assert.NotNil(t, normals)
	// 9 vertices, 6 faces, 3 components per vertex
	assert.Equal(t, 9*6*3, len(normals.Data))

	uvs := m.GetAttribute("uv")
	assert.NotNil(t, uvs)
	// 9 vertices, 6 faces, 2 components per vertex
	assert.Equal(t, 9*6*2, len(uvs.Data))
}
