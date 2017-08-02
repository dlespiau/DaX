package dax

import (
	"testing"

	"github.com/dlespiau/dax/math"
	"github.com/stretchr/testify/assert"
)

type dummyOpaqueMaterial struct {
	BaseMaterial
}

type dummerMesher struct{}

func (m *dummerMesher) GetMesh() *Mesh { return NewMesh() }

func createDummyNode() *Node {
	mr := NewMeshRenderer(&dummerMesher{}, &dummyOpaqueMaterial{})
	return NewNode().AddComponent(mr)
}

var _ Material = &dummyOpaqueMaterial{}

//  y
//  ^
//  |
//  |  a  b  c
//  +--+--+--+--> x
//
//  a, b: two nodes
//    - a is placed at (0.1, 0, 0)
//    - b is a child of a (0.1, 0, 0) relatively to a
//  c: camera, placed at (0.3, 0, 0), looking at (0, 0, 0)

type testCtx struct {
	sg   *SceneGraph
	a, b *Node
	c    Camera
}

func buildTestSceneGraph() *testCtx {
	sg := NewSceneGraph()

	a := createDummyNode()
	a.SetPosition(0.1, 0, 0)
	sg.AddChild(a)

	b := createDummyNode()
	b.SetPosition(0.1, 0, 0)
	a.AddChild(b)

	// This camera won't scale nor skew the distances we define!
	// We also want a and b to be in its frustrum.
	c := NewOrthographicCamera(-1, 1, -1, 1, 1, -1)
	c.SetPosition(0.3, 0, 0)
	c.LookAt(&math.Vec3{0, 0, 0})

	sg.updateWorldTransform()

	return &testCtx{
		sg: sg,
		a:  a,
		b:  b,
		c:  c,
	}
}

func TestOpaqueFrontToBack(t *testing.T) {
	ctx := buildTestSceneGraph()
	cameraTransform := cameraTransform(ctx.c)
	nodes := opaqueFrontToBack(ctx.sg, cameraTransform)

	assert.Equal(t, 2, len(nodes))

	assert.Equal(t, ctx.b, nodes[0].node)
	assertFloat(t, -0.1, nodes[0].z, 1e-6)

	assert.Equal(t, ctx.a, nodes[1].node)
	assertFloat(t, -0.2, nodes[1].z, 1e-6)
}
