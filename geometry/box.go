package geometry

import (
	dax "github.com/dlespiau/dax/lib"
	"github.com/dlespiau/dax/math"
)

// Box is a rectangular cuboid centered around (0, 0, 0) with sizes Width,
// Height and Depth on the X, Y and Z axis respectively. Some control over the
// tesselation of each face is given through the number of segments on each
// direction.
type Box struct {
	Width, Height, Depth                                  float32
	NumWidthSegments, NumHeightSegments, NumDepthSegments int
}

// BoxOptions contains optional parameters for the Box constructors.
type BoxOptions struct {
	NumWidthSegments, NumHeightSegments, NumDepthSegments int
}

var defaultBox = Box{
	Width:             1.0,
	Height:            1.0,
	Depth:             1.0,
	NumWidthSegments:  1,
	NumHeightSegments: 1,
	NumDepthSegments:  1,
}

// NewBox creates a new box.
func NewBox(width, height, depth float32, options ...BoxOptions) *Box {
	box := defaultBox
	box.Width = width
	box.Height = height
	box.Depth = depth

	if len(options) == 0 {
		return &box
	}

	if options[0].NumWidthSegments > 0 {
		box.NumWidthSegments = options[0].NumWidthSegments
	}
	if options[0].NumHeightSegments > 0 {
		box.NumHeightSegments = options[0].NumHeightSegments
	}
	if options[0].NumDepthSegments > 0 {
		box.NumDepthSegments = options[0].NumDepthSegments
	}

	return &box
}

type boxContext struct {
	nVertices int
	positions []float32
	normals   []float32
	uvs       []float32
	indices   []uint
}

func buildPlane(ctx *boxContext,
	u, v, w int,
	udir, vdir float32,
	width, height, depth float32,
	gridX, gridY int) {

	segmentWidth := width / float32(gridX)
	segmentHeight := height / float32(gridY)

	widthHalf := width / 2
	heightHalf := height / 2
	depthHalf := depth / 2

	gridX1 := gridX + 1
	gridY1 := gridY + 1

	vertexCounter := 0

	vector := math.Vec3{}

	// generate vertices, normals and uvs

	for iy := 0; iy < gridY1; iy++ {
		y := float32(iy)*segmentHeight - heightHalf

		for ix := 0; ix < gridX1; ix++ {
			x := float32(ix)*segmentWidth - widthHalf

			// position
			vector[u] = x * udir
			vector[v] = y * vdir
			vector[w] = depthHalf

			ctx.positions = append(ctx.positions, vector[0], vector[1], vector[2])

			// normal
			vector[u] = 0
			vector[v] = 0
			if depth > 0 {
				vector[w] = 1
			} else {
				vector[w] = -1
			}

			ctx.normals = append(ctx.normals, vector[0], vector[1], vector[2])

			// uvs
			ctx.uvs = append(ctx.uvs, float32(ix)/float32(gridX), 1-(float32(iy)/float32(gridY)))

			// counters
			vertexCounter++

		}

	}

	// indices

	// 1. you need three indices to draw a single face
	// 2. a single segment consists of two faces
	// 3. so we need to generate six (2*3) indices per segment

	for iy := 0; iy < gridY; iy++ {
		for ix := 0; ix < gridX; ix++ {
			a := uint(ctx.nVertices + ix + gridX1*iy)
			b := uint(ctx.nVertices + ix + gridX1*(iy+1))
			c := uint(ctx.nVertices + (ix + 1) + gridX1*(iy+1))
			d := uint(ctx.nVertices + (ix + 1) + gridX1*iy)

			// faces
			ctx.indices = append(ctx.indices, a, b, d, b, c, d)
		}

	}

	// update total number of vertices
	ctx.nVertices += vertexCounter
}

// GetMesh is part of the dax.Mesher interface.
func (b *Box) GetMesh() *dax.Mesh {

	m := dax.NewMesh()

	ctx := &boxContext{}

	width := b.Width
	height := b.Height
	depth := b.Depth
	widthSegments := b.NumWidthSegments
	heightSegments := b.NumHeightSegments
	depthSegments := b.NumDepthSegments

	buildPlane(ctx, 2, 1, 0, -1, -1, depth, height, width, depthSegments, heightSegments)  // px
	buildPlane(ctx, 2, 1, 0, 1, -1, depth, height, -width, depthSegments, heightSegments)  // nx
	buildPlane(ctx, 0, 2, 1, 1, 1, width, depth, height, widthSegments, depthSegments)     // py
	buildPlane(ctx, 0, 2, 1, 1, -1, width, depth, -height, widthSegments, depthSegments)   // ny
	buildPlane(ctx, 0, 1, 2, 1, -1, width, height, depth, widthSegments, heightSegments)   // pz
	buildPlane(ctx, 0, 1, 2, -1, -1, width, height, -depth, widthSegments, heightSegments) // nz

	m.AddAttribute("position", ctx.positions, 3)
	m.AddAttribute("normal", ctx.normals, 3)
	m.AddAttribute("uv", ctx.uvs, 2)
	m.AddIndices(ctx.indices)

	return m
}
