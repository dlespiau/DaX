package geometry

import (
	"github.com/dlespiau/dax"
	"github.com/dlespiau/dax/math"
)

type Sphere struct {
	radius                  float32
	nVSegments, nHSegments  int
	phiStart, phiLength     float32
	thetaStart, thetaLength float32
}

func NewSphere(radius float32, nVSegments, nHSegments int) *Sphere {
	s := new(Sphere)
	s.Init(radius, nVSegments, nHSegments)
	return s
}

func (s *Sphere) InitFull(radius float32, nVSegments, nHSegments int,
	phiStart, phiLength, thetaStart, thetaLength float32) {
	s.radius = radius
	s.nVSegments = nVSegments
	s.nHSegments = nHSegments
	s.phiStart = phiStart
	s.phiLength = phiLength
	s.thetaStart = thetaStart
	s.thetaLength = thetaLength
}

func (s *Sphere) Init(radius float32, nVSegments, nHSegments int) {
	const angle float32 = 2 * float32(math.Pi)

	s.InitFull(radius, nVSegments, nHSegments, 0, angle, 0, angle)
}

func (s *Sphere) GetMesh() *dax.Mesh {
	m := dax.NewMesh()
	var positions, normals, uvs dax.AttributeBuffer

	thetaEnd := s.thetaStart + s.thetaLength
	vertexCount := (s.nVSegments + 1) * (s.nHSegments + 1)

	positions.Init("position", vertexCount, 3)
	normals.Init("normal", vertexCount, 3)
	uvs.Init("uvs", vertexCount, 2)

	index := 0
	vertices := make([][]uint, s.nHSegments+1, s.nHSegments+1)
	normal := math.Vec3{}

	for y := 0; y <= s.nHSegments; y++ {

		verticesRow := make([]uint, s.nVSegments+1, s.nVSegments+1)

		v := float32(y) / float32(s.nHSegments)

		for x := 0; x <= s.nVSegments; x++ {

			u := float32(x) / float32(s.nVSegments)

			px := -s.radius * math.Cos(s.phiStart+u*s.phiLength) * math.Sin(s.thetaStart+v*s.thetaLength)
			py := s.radius * math.Cos(s.thetaStart+v*s.thetaLength)
			pz := s.radius * math.Sin(s.phiStart+u*s.phiLength) * math.Sin(s.thetaStart+v*s.thetaLength)

			normal.Set(px, py, pz)
			normal.Normalize()

			positions.SetXYZ(index, px, py, pz)
			normals.SetXYZ(index, normal[0], normal[1], normal[2])
			uvs.SetXY(index, u, 1-v)

			verticesRow[x] = uint(index)

			index++
		}

		vertices[y] = verticesRow

	}

	indices := make([]uint, vertexCount, vertexCount)
	i := 0

	for y := 0; y < s.nHSegments; y++ {

		for x := 0; x < s.nVSegments; x++ {

			v1 := vertices[y][x+1]
			v2 := vertices[y][x]
			v3 := vertices[y+1][x]
			v4 := vertices[y+1][x+1]

			if y != 0 || s.thetaStart > 0 {
				indices[i] = v1
				i++
				indices[i] = v2
				i++
				indices[i] = v4
				i++
			}
			if y != s.nHSegments-1 || thetaEnd < math.Pi {
				indices[i] = v2
				i++
				indices[i] = v3
				i++
				indices[i] = v4
				i++
			}
		}
	}

	m.AddIndices(indices)
	m.AddAttributeBuffer(&positions)
	m.AddAttributeBuffer(&normals)
	m.AddAttributeBuffer(&uvs)

	return m
}
