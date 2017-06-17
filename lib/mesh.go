package dax

import (
	"unsafe"

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

type IndexBuffer struct {
	data16 []uint16
	data32 []uint32
	vbo    uint32
}

func (ib *IndexBuffer) Init(size int) {
	if size > 65536 {
		ib.data32 = make([]uint32, size, size)
	} else {
		ib.data16 = make([]uint16, size, size)
	}
	gl.GenBuffers(1, &ib.vbo)
}

func (ib *IndexBuffer) InitFromData(data []uint) {
	ib.Init(len(data))
	for i, v := range data {
		ib.Set(i, v)
	}
}

func (ib *IndexBuffer) Destroy() {
	gl.DeleteBuffers(1, &ib.vbo)
}

func (ib *IndexBuffer) Set(nth int, index uint) {
	if ib.data16 != nil {
		ib.data16[nth] = uint16(index)
	} else {
		ib.data32[nth] = uint32(index)
	}
}

func (ib *IndexBuffer) Upload() {
	var size int
	var ptr unsafe.Pointer

	if ib.data16 != nil {
		size = len(ib.data16) * 2
		ptr = gl.Ptr(ib.data16)
	} else {
		size = len(ib.data32) * 4
		ptr = gl.Ptr(ib.data32)
	}

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, size, ptr, gl.STATIC_DRAW)
}

const (
	has_position = 1 << iota
	has_color
)

type Mesh struct {
	flags      uint32
	vao        uint32
	attributes []AttributeBuffer
	indices    IndexBuffer
}

func NewMesh() *Mesh {
	m := new(Mesh)

	gl.GenVertexArrays(1, &m.vao)

	return m
}

func (m *Mesh) Destroy() {
	for _, ab := range m.attributes {
		ab.Destroy()
	}
	m.indices.Destroy()
	gl.DeleteVertexArrays(1, &m.vao)
}

func (m *Mesh) getAttribute(name string) *AttributeBuffer {
	for _, ab := range m.attributes {
		if ab.name == name {
			return &ab
		}
	}

	return nil
}

func (m *Mesh) getNewAttribute(name string) *AttributeBuffer {
	ab := m.getAttribute(name)
	if ab != nil {
		ab.Destroy()
	} else {
		var buffer AttributeBuffer

		m.attributes = append(m.attributes, buffer)
		ab = &m.attributes[len(m.attributes)-1]
	}

	return ab
}

func (m *Mesh) AddAttribute(name string, data []float32, nComponents int) {
	m.Bind()

	ab := m.getNewAttribute(name)
	ab.InitFromData(name, data, nComponents)
	ab.Upload()
}

func (m *Mesh) AddAttributeBuffer(buffer *AttributeBuffer) {
	m.Bind()

	ab := m.getNewAttribute(buffer.name)
	*ab = *buffer
	ab.Upload()
}

func (m *Mesh) HasIndices() bool {
	return m.indices.data16 != nil || m.indices.data32 != nil
}

func (m *Mesh) AddIndices(data []uint) {
	m.Bind()

	m.indices.InitFromData(data)
	m.indices.Upload()
}

func (m *Mesh) Bind() {
	gl.BindVertexArray(m.vao)
}
