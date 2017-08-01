package dax

// AttributeBuffer holds per-vertex attribute. There is one AttributeBuffer per
// kind of data we want to keep with each vertex.
type AttributeBuffer struct {
	Name          string
	NumComponents int
	Data          []float32
}

func NewAttributeBuffer(name string, size int, NumComponents int) *AttributeBuffer {
	ab := new(AttributeBuffer)
	ab.Init(name, size, NumComponents)
	return ab
}

func (ab *AttributeBuffer) Init(name string, size int, NumComponents int) {
	data := make([]float32, size*NumComponents, size*NumComponents)
	ab.InitFromData(name, data, NumComponents)
}

func (ab *AttributeBuffer) InitFromData(name string, data []float32, NumComponents int) {
	ab.Name = name
	ab.NumComponents = NumComponents
	ab.Data = data
}

// Len returns the number of elements in the AttributeBuffer. Because each
// element has a number of components, the length of the Data array is then
// Len() * NumComponents.
func (ab *AttributeBuffer) Len() int {
	return len(ab.Data) / ab.NumComponents
}

func (ab *AttributeBuffer) SetX(index int, x float32) {
	ab.Data[index*ab.NumComponents+0] = x
}

func (ab *AttributeBuffer) GetX(index int) (x float32) {
	x = ab.Data[index*ab.NumComponents+0]
	return
}

func (ab *AttributeBuffer) SetXY(index int, x, y float32) {
	ab.Data[index*ab.NumComponents+0] = x
	ab.Data[index*ab.NumComponents+1] = y
}

func (ab *AttributeBuffer) GetXY(index int) (x, y float32) {
	x = ab.Data[index*ab.NumComponents+0]
	y = ab.Data[index*ab.NumComponents+1]
	return
}

func (ab *AttributeBuffer) SetXYZ(index int, x, y, z float32) {
	ab.Data[index*ab.NumComponents+0] = x
	ab.Data[index*ab.NumComponents+1] = y
	ab.Data[index*ab.NumComponents+2] = z
}

func (ab *AttributeBuffer) GetXYZ(index int) (x, y, z float32) {
	x = ab.Data[index*ab.NumComponents+0]
	y = ab.Data[index*ab.NumComponents+1]
	z = ab.Data[index*ab.NumComponents+2]
	return
}

func (ab *AttributeBuffer) SetXYZW(index int, x, y, z, w float32) {
	ab.Data[index*ab.NumComponents+0] = x
	ab.Data[index*ab.NumComponents+1] = y
	ab.Data[index*ab.NumComponents+2] = z
	ab.Data[index*ab.NumComponents+3] = w

}

func (ab *AttributeBuffer) GetXYZW(index int) (x, y, z, w float32) {
	x = ab.Data[index*ab.NumComponents+0]
	y = ab.Data[index*ab.NumComponents+1]
	z = ab.Data[index*ab.NumComponents+2]
	w = ab.Data[index*ab.NumComponents+3]
	return
}

type IndexBuffer struct {
	data16 []uint16
	data32 []uint32
}

func (ib *IndexBuffer) Init(size int) {
	if size > 65536 {
		ib.data32 = make([]uint32, size, size)
	} else {
		ib.data16 = make([]uint16, size, size)
	}
}

func (ib *IndexBuffer) Len() int {
	if len(ib.data16) > 0 {
		return len(ib.data16)
	}
	return len(ib.data32)
}

func (ib *IndexBuffer) InitFromData(data []uint) {
	ib.Init(len(data))
	for i, v := range data {
		ib.Set(i, v)
	}
}

func (ib *IndexBuffer) Set(nth int, index uint) {
	if ib.data16 != nil {
		ib.data16[nth] = uint16(index)
	} else {
		ib.data32[nth] = uint32(index)
	}
}

// VertexMode defines how vertices should be interpreted by the draw call.
type VertexMode int

const (
	// VertexModePoints draws a single dot for each vertex.
	VertexModePoints VertexMode = iota
	VertexModeLineStrip
	VertexModeLineLoop
	VertexModeLines
	VertexModeTriangleStrip
	VertexModeTriangleFan
	VertexModeTriangles
)

type Mesh struct {
	flags      uint32
	mode       VertexMode
	attributes []AttributeBuffer
	indices    IndexBuffer
}

func NewMesh() *Mesh {
	m := &Mesh{
		mode: VertexModeTriangles,
	}

	return m
}

// GetVertexMode returns how vertices in the Mesh are interpreted. New meshes
// default to VertexModeTriangles.
func (m *Mesh) GetVertexMode() VertexMode {
	return m.mode
}

// SetVertexMode sets how vertices in the should be interpreted.
func (m *Mesh) SetVertexMode(mode VertexMode) {
	m.mode = mode
}

func (m *Mesh) GetAttribute(name string) *AttributeBuffer {
	for _, ab := range m.attributes {
		if ab.Name == name {
			return &ab
		}
	}

	return nil
}

func (m *Mesh) getNewAttribute(name string) *AttributeBuffer {
	ab := m.GetAttribute(name)
	if ab != nil {
		return ab
	}

	var buffer AttributeBuffer

	m.attributes = append(m.attributes, buffer)
	return &m.attributes[len(m.attributes)-1]
}

func (m *Mesh) AddAttribute(name string, data []float32, NumComponents int) {
	ab := m.getNewAttribute(name)
	ab.InitFromData(name, data, NumComponents)
}

func (m *Mesh) AddAttributeBuffer(buffer *AttributeBuffer) {
	ab := m.getNewAttribute(buffer.Name)
	*ab = *buffer
}

func (m *Mesh) HasIndices() bool {
	return m.indices.data16 != nil || m.indices.data32 != nil
}

func (m *Mesh) AddIndices(data []uint) {
	m.indices.InitFromData(data)
}
