package dax

import (
	"github.com/dlespiau/dax/math"
)

type Node struct {
	// Grapher
	parent   Grapher
	children []Grapher

	transformValid      bool
	worldTransformValid bool

	position math.Vec3
	rotation math.Quaternion
	scale    math.Vec3

	transform      math.Transform
	worldTransform math.Transform
}

func NewNode() *Node {
	n := new(Node)
	n.Init()
	return n
}

func (n *Node) Init() {
	n.rotation.Iden()
	n.scale = math.Vec3{1, 1, 1}
}

func (n *Node) GetPosition() *math.Vec3 {
	return &n.position
}

func (n *Node) SetPosition(x, y, z float32) {
	n.position[0] = x
	n.position[1] = y
	n.position[2] = z
	n.transformValid = false
}

func (n *Node) SetPositionV(position *math.Vec3) {
	n.position = *position
	n.transformValid = false
}

func (n *Node) Translate(tx, ty, tz float32) {
	n.position[0] += tx
	n.position[1] += ty
	n.position[2] += tz
	n.transformValid = false
}

func (n *Node) TranslateV(t *math.Vec3) {
	n.position[0] += t[0]
	n.position[1] += t[1]
	n.position[2] += t[2]
	n.transformValid = false
}

func (n *Node) TranslateX(tx float32) {
	n.position[0] += tx
	n.transformValid = false
}

func (n *Node) TranslateY(ty float32) {
	n.position[1] += ty
	n.transformValid = false
}

func (n *Node) TranslateZ(tz float32) {
	n.position[2] += tz
	n.transformValid = false
}

func (n *Node) GetRotation() *math.Quaternion {
	return &n.rotation
}

func (n *Node) SetRotation(q *math.Quaternion) {
	n.rotation = *q
	n.transformValid = false
}

func (n *Node) RotateAroundAxis(axis *math.Vec3, angle float32) {
	q := math.QuatRotate(angle, axis)
	n.rotation.MulWith(&q)
	n.transformValid = false
}

func (n *Node) RotateX(angle float32) {
	n.RotateAroundAxis(&math.Vec3{1, 0, 0}, angle)
	n.transformValid = false
}

func (n *Node) RotateY(angle float32) {
	n.RotateAroundAxis(&math.Vec3{0, 1, 0}, angle)
	n.transformValid = false
}

func (n *Node) RotateZ(angle float32) {
	n.RotateAroundAxis(&math.Vec3{0, 0, 1}, angle)
	n.transformValid = false
}

func (n *Node) GetScale() *math.Vec3 {
	return &n.scale
}

func (n *Node) SetScale(sx, sy, sz float32) {
	n.scale[0] = sx
	n.scale[1] = sy
	n.scale[2] = sz
	n.transformValid = false
}

func (n *Node) SetScaleV(s *math.Vec3) {
	n.scale = *s
	n.transformValid = false
}

func (n *Node) Scale(sx, sy, sz float32) {
	n.scale[0] *= sx
	n.scale[1] *= sy
	n.scale[2] *= sz
	n.transformValid = false
}

func (n *Node) ScaleV(s *math.Vec3) {
	n.scale[0] *= s[0]
	n.scale[1] *= s[1]
	n.scale[2] *= s[2]
	n.transformValid = false
}

func (n *Node) ScaleX(sx float32) {
	n.scale[0] *= sx
	n.transformValid = false
}

func (n *Node) ScaleY(sy float32) {
	n.scale[1] *= sy
	n.transformValid = false
}

func (n *Node) ScaleZ(sz float32) {
	n.scale[2] *= sz
	n.transformValid = false
}

func (n *Node) updateTransform() {
	if n.transformValid {
		return
	}

	n.transform.SetTranslateVec3(&n.position)
	n.transform.RotateQuat(&n.rotation)
	n.transform.ScaleVec3(&n.scale)
	n.transformValid = true
	n.worldTransformValid = false
}

func (n *Node) getTransform() *math.Transform {
	n.updateTransform()
	return &n.transform
}

func (n *Node) GetTransform() *math.Mat4 {
	n.updateTransform()
	return (*math.Mat4)(&n.transform)
}

// updateWorldTransform will update the transformation from node space to world
// space recursively on all nodes.
// force can be used to force the updates on children when a parent has changed
// its transform and we, then, need to update the world transform on that
// subtree.
func (n *Node) updateWorldTransform(force bool) {
	// Start by updating the local transform, and, as side effect,
	// worldTransformValid
	n.updateTransform()

	if !n.worldTransformValid || force {
		if n.parent == nil {
			// this node isn't parented (root or not part of a
			// scene graph)
			n.worldTransform = n.transform
		} else {
			// compose with parent transform
			parent := (n.parent).(*Node)
			world := (*math.Mat4)(&parent.worldTransform)
			local := (*math.Mat4)(&n.transform)

			(*math.Mat4)(&n.worldTransform).Mul4Of(world, local)
		}

		force = true
	}

	for _, child := range n.children {
		node := child.(*Node)
		node.updateWorldTransform(force)
	}
}

// Grapher implementation

func (n *Node) GetParent() Grapher {
	return n.parent
}

func (n *Node) SetParent(parent Grapher) {
	n.parent = parent
}

func (n *Node) AddChild(child Grapher) {
	child.SetParent(n)
	n.children = append(n.children, child)
}

func (n *Node) GetChildren() []Grapher {
	return n.children
}
