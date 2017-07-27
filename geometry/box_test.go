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

func TestNewBoxNoOptions(t *testing.T) {
	box := NewBox(10, 20, 30)
	assert.Equal(t, float32(10), box.Width)
	assert.Equal(t, float32(20), box.Height)
	assert.Equal(t, float32(30), box.Depth)
	assert.Equal(t, 1, box.NumWidthSegments)
	assert.Equal(t, 1, box.NumHeightSegments)
	assert.Equal(t, 1, box.NumDepthSegments)
}

func TestBoxMesh(t *testing.T) {
	box := NewBox(10, 20, 30, BoxOptions{
		NumWidthSegments:  2,
		NumHeightSegments: 2,
		NumDepthSegments:  2,
	})

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
