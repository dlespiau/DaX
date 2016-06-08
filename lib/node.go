package dax

import (
	"github.com/dlespiau/dax/math"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Position math.Vec3
	Rotation math.Quaternion
	Scale    math.Vec3
}

func (n *Node) Init() {
	n.Rotation.Iden()
	n.Scale = math.Vec3{1, 1, 1}
}

func (n *Node) SetPosition(x, y, z float32) {
	n.Position[0] = x
	n.Position[1] = y
	n.Position[2] = z
}

func (n *Node) TranslateX(tx float32) {
	n.Position[0] += tx
}

func (n *Node) TranslateY(ty float32) {
	n.Position[1] += ty
}

func (n *Node) TranslateZ(tz float32) {
	n.Position[2] += tz
}

func (n *Node) SetRotation(q math.Quaternion) {
	n.Rotation = q
}

func (n *Node) RotateAroundAxis(axis *math.Vec3, angle float32) {
	q := math.QuatRotate(angle, axis)
	n.Rotation.MulWith(&q)
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

func (n *Node) SetScale(sx, sy, sz float32) {
	n.Scale[0] = sx
	n.Scale[1] = sy
	n.Scale[2] = sz
}

func (n *Node) ScaleX(sx float32) {
	n.Scale[0] *= sx
}

func (n *Node) ScaleY(sy float32) {
	n.Scale[1] *= sy
}

func (n *Node) ScaleZ(sz float32) {
	n.Scale[2] *= sz
}

func (n *Node) AddChild(child *Node) {
	child.Parent = n
	n.Children = append(n.Children, child)
}
