package dax

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type VertexAttribute int

const (
	AttributePosition VertexAttribute = iota
	AttributeNormal
	AttributeTextureCoordinate
	AttributeColor
)

const (
	has_position = 1 << iota
	has_color
)

type mesh struct {
	flags uint32
	vao   uint32
	vbo   uint32
}

func NewMesh() *mesh {
	m := new(mesh)

	gl.GenVertexArrays(1, &m.vao)

	return m
}

func (m *mesh) Destroy() {
	gl.DeleteBuffers(1, &m.vbo)
	gl.DeleteVertexArrays(1, &m.vao)
}

func (m *mesh) SetAttribute(attrib VertexAttribute, data []float32) {
	m.Bind()

	gl.DeleteBuffers(1, &m.vbo)
	gl.GenBuffers(1, &m.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)

}

func (m *mesh) Bind() {
	gl.BindVertexArray(m.vao)
}
