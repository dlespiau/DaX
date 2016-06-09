package dax

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type AttributeBuffer struct {
	name        string
	nComponents int
	data        []float32
	vbo         uint32
}

func NewAttributeBuffer(name string, size int, nComponents int) *AttributeBuffer {
	ab := new(AttributeBuffer)
	ab.Init(name, size, nComponents)
	return ab
}

func (ab *AttributeBuffer) Init(name string, size int, nComponents int) {
	data := make([]float32, size*nComponents, size*nComponents)
	ab.InitFromData(name, data, nComponents)
}

func (ab *AttributeBuffer) InitFromData(name string, data []float32, nComponents int) {
	ab.name = name
	ab.nComponents = nComponents
	ab.data = data

	gl.GenBuffers(1, &ab.vbo)
}

func (ab *AttributeBuffer) Destroy() {
	gl.DeleteBuffers(1, &ab.vbo)
}

func (ab *AttributeBuffer) SetX(index int, x float32) {
	ab.data[index*ab.nComponents+0] = x
}

func (ab *AttributeBuffer) GetX(index int) (x float32) {
	x = ab.data[index*ab.nComponents+0]
	return
}

func (ab *AttributeBuffer) SetXY(index int, x, y float32) {
	ab.data[index*ab.nComponents+0] = x
	ab.data[index*ab.nComponents+1] = y
}

func (ab *AttributeBuffer) GetXY(index int) (x, y float32) {
	x = ab.data[index*ab.nComponents+0]
	y = ab.data[index*ab.nComponents+1]
	return
}

func (ab *AttributeBuffer) SetXYZ(index int, x, y, z float32) {
	ab.data[index*ab.nComponents+0] = x
	ab.data[index*ab.nComponents+1] = y
	ab.data[index*ab.nComponents+2] = z
}

func (ab *AttributeBuffer) GetXYZ(index int) (x, y, z float32) {
	x = ab.data[index*ab.nComponents+0]
	y = ab.data[index*ab.nComponents+1]
	z = ab.data[index*ab.nComponents+2]
	return
}

func (ab *AttributeBuffer) SetXYZW(index int, x, y, z, w float32) {
	ab.data[index*ab.nComponents+0] = x
	ab.data[index*ab.nComponents+1] = y
	ab.data[index*ab.nComponents+2] = z
	ab.data[index*ab.nComponents+3] = w

}

func (ab *AttributeBuffer) GetXYZW(index int) (x, y, z, w float32) {
	x = ab.data[index*ab.nComponents+0]
	y = ab.data[index*ab.nComponents+1]
	z = ab.data[index*ab.nComponents+2]
	w = ab.data[index*ab.nComponents+3]
	return
}

func (ab *AttributeBuffer) Upload() {
	gl.BindBuffer(gl.ARRAY_BUFFER, ab.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(ab.data)*4, gl.Ptr(ab.data), gl.STATIC_DRAW)
}

const (
	has_position = 1 << iota
	has_color
)

type mesh struct {
	flags      uint32
	vao        uint32
	attributes []AttributeBuffer
}

func NewMesh() *mesh {
	m := new(mesh)

	gl.GenVertexArrays(1, &m.vao)

	return m
}

func (m *mesh) Destroy() {
	for _, ab := range m.attributes {
		ab.Destroy()
	}
	gl.DeleteVertexArrays(1, &m.vao)
}

func (m *mesh) getAttribute(name string) *AttributeBuffer {
	for _, ab := range m.attributes {
		if ab.name == name {
			return &ab
		}
	}

	return nil
}

func (m *mesh) AddAttribute(name string, data []float32, nComponents int) {
	m.Bind()

	ab := m.getAttribute(name)
	if ab != nil {
		ab.Destroy()
	} else {
		var buffer AttributeBuffer

		m.attributes = append(m.attributes, buffer)
		ab = &m.attributes[len(m.attributes)-1]
	}

	ab.InitFromData(name, data, nComponents)
	ab.Upload()
}

func (m *mesh) Bind() {
	gl.BindVertexArray(m.vao)
}
