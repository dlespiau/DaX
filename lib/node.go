package dax

import (
	"github.com/dlespiau/dax/math"
)

type Node struct {
	// Grapher
	parent   Grapher
	children []Grapher

	position math.Vec3
	rotation math.Quaternion
	scale    math.Vec3
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
}

func (n *Node) SetPositionV(position *math.Vec3) {
	n.position = *position
}

func (n *Node) TranslateX(tx float32) {
	n.position[0] += tx
}

func (n *Node) TranslateY(ty float32) {
	n.position[1] += ty
}

func (n *Node) TranslateZ(tz float32) {
	n.position[2] += tz
}

func (n *Node) GetRotation() *math.Quaternion {
	return &n.rotation
}

func (n *Node) SetRotation(q *math.Quaternion) {
	n.rotation = *q
}

func (n *Node) RotateAroundAxis(axis *math.Vec3, angle float32) {
	q := math.QuatRotate(angle, axis)
	n.rotation.MulWith(&q)
}

func (n *Node) RotateX(angle float32) {
	n.RotateAroundAxis(&math.Vec3{1, 0, 0}, angle)
}

func (n *Node) RotateY(angle float32) {
	n.RotateAroundAxis(&math.Vec3{0, 1, 0}, angle)
}

func (n *Node) RotateZ(angle float32) {
	n.RotateAroundAxis(&math.Vec3{0, 0, 1}, angle)
}

func (n *Node) GetScale() *math.Vec3 {
	return &n.scale
}

func (n *Node) SetScale(sx, sy, sz float32) {
	n.scale[0] = sx
	n.scale[1] = sy
	n.scale[2] = sz
}

func (n *Node) SetScaleV(s *math.Vec3) {
	n.scale = *s
}

func (n *Node) ScaleX(sx float32) {
	n.scale[0] *= sx
}

func (n *Node) ScaleY(sy float32) {
	n.scale[1] *= sy
}

func (n *Node) ScaleZ(sz float32) {
	n.scale[2] *= sz
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
