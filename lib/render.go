package dax

import (
	"fmt"
	"os"
	"sort"
	"unsafe"

	"github.com/dlespiau/dax/math"

	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	polylineMaterial = "-dax-material-polyline"
)

type uploadInput struct {
	fb Framebuffer
}

type glUploader interface {
	upload(input uploadInput)
}

type glUniform struct {
	uniform  Uniform
	location int32
}

// A uniform from the library user
type glUniformUser glUniform

func (u *glUniformUser) upload(input uploadInput) {
	panic("not implemented")
}

// A uniform that upload the ModelViewProjection matrix
type glUniformMVP glUniform

func (u *glUniformMVP) upload(input uploadInput) {
	cameraTransform := cameraTransform(input.fb.GetCamera())
	gl.UniformMatrix4fv(u.location, 1, false, &cameraTransform[0])
}

type glAttributeBuffer struct {
	buffer *AttributeBuffer
	id     uint32
}

// upload needs to be called after the vao has been bound!
func (vbo *glAttributeBuffer) upload() {
	ab := vbo.buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
	gl.BufferData(gl.ARRAY_BUFFER, len(ab.Data)*4, gl.Ptr(ab.Data), gl.STATIC_DRAW)
}

type glIndexBuffer struct {
	buffer *IndexBuffer
	id     uint32
}

// upload needs to be called after the has been bound!
func (vbo *glIndexBuffer) upload() {
	if vbo.id == 0 {
		// A mesh doesn't always use indices.
		return
	}

	ib := vbo.buffer
	var size int
	var ptr unsafe.Pointer

	if ib.data16 != nil {
		size = len(ib.data16) * 2
		ptr = gl.Ptr(ib.data16)
	} else {
		size = len(ib.data32) * 4
		ptr = gl.Ptr(ib.data32)
	}

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo.id)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, size, ptr, gl.STATIC_DRAW)
}

type glVAO struct {
	id      uint32
	vbos    []glAttributeBuffer
	indices glIndexBuffer
}

func newVAOFromMesh(mesh *Mesh) *glVAO {
	vao := &glVAO{}

	gl.GenVertexArrays(1, &vao.id)

	for i := range mesh.attributes {
		ab := &mesh.attributes[i]

		vbo := glAttributeBuffer{
			buffer: ab,
		}
		gl.GenBuffers(1, &vbo.id)

		vao.vbos = append(vao.vbos, vbo)
	}

	if mesh.indices.Len() > 0 {
		vao.indices.buffer = &mesh.indices
		gl.GenBuffers(1, &vao.indices.id)
	}

	return vao
}

func (vao *glVAO) bind() {
	gl.BindVertexArray(vao.id)
}

func (vao *glVAO) upload() {
	for i := range vao.vbos {
		vao.vbos[i].upload()
	}
	vao.indices.upload()
}

func (vao *glVAO) destroy() {
	for i := range vao.vbos {
		vbo := &vao.vbos[i]
		gl.DeleteBuffers(1, &vbo.id)
	}
	if vao.indices.id != 0 {
		gl.DeleteBuffers(1, &vao.indices.id)
	}
	gl.DeleteVertexArrays(1, &vao.id)
}

type glProgram struct {
	id        uint32
	vs        *VertexShader
	fs        *FragmentShader
	uploaders []glUploader
}

func glVertexMode(mode VertexMode) uint32 {
	switch mode {
	case VertexModePoints:
		return gl.POINTS
	case VertexModeLineStrip:
		return gl.LINE_STRIP
	case VertexModeLineLoop:
		return gl.LINE_LOOP
	case VertexModeLines:
		return gl.LINES
	case VertexModeTriangleStrip:
		return gl.TRIANGLE_STRIP
	case VertexModeTriangleFan:
		return gl.TRIANGLE_FAN
	case VertexModeTriangles:
		return gl.TRIANGLES
	default:
		panic("Unknown vertex mode")
	}
}

func glIndexType(ib *IndexBuffer) uint32 {
	if ib.data16 != nil {
		return gl.UNSIGNED_SHORT
	}
	return gl.UNSIGNED_INT
}

type renderer struct {
	// material ID -> glProgram
	programs map[string]*glProgram
	// The only vs we currently have :/
	vs *VertexShader
}

const vertexShader = `
#version 330 core

in vec3 position;

uniform mat4 mvp;

void main(){
	gl_Position = mvp * vec4(position, 1.0f);
}`

func newRenderer() *renderer {
	vs := NewVertexShader(vertexShader)
	vs.AddAttribute(VariableKindVec3, "position")
	vs.AddUniform(VariableKindMat4, "mvp")

	return &renderer{
		programs: make(map[string]*glProgram),
		vs:       vs,
	}
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cSources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, cSources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		//log := strings.Repeat('\x00', int(logLength+1))
		//gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		log := "TODO"
		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func makeProgram(v *VertexShader, f *FragmentShader) (uint32, error) {
	vertexShader, err := compileShader(v.source, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(f.source, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		//log := strings.Repeat("\x00", int(logLength+1))
		log := "TODO"
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	// XXX: we probably want to make that a FS property. Or define it in the shader
	// source itself.
	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	return program, nil
}

const fragmentShader = `
#version 330
uniform vec4 color;
out vec4 outputColor;
void main() {
    outputColor = color;
}`

func (r *renderer) makePolylineProgram() *glProgram {
	if p, ok := r.programs[polylineMaterial]; ok {
		return p
	}

	s := NewFragmentShader(fragmentShader)
	p, err := makeProgram(r.vs, s)
	if err != nil {
		panic(err)
	}
	program := &glProgram{
		id: p,
		vs: r.vs,
		fs: s,
	}
	r.programs[polylineMaterial] = program
	return program
}
func (r *renderer) drawPolyline(fb Framebuffer, p *Polyline) {
	if p.Size() == 0 {
		return
	}

	program := r.makePolylineProgram()

	mesh := NewMesh()
	mesh.AddAttribute("position", p.vertices, 3)
	vao := newVAOFromMesh(mesh)

	defer vao.destroy()

	vao.bind()
	vao.upload()

	gl.UseProgram(program.id)

	position := uint32(gl.GetAttribLocation(program.id, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	mvp := gl.GetUniformLocation(program.id, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvp, 1, false, &fb.GetCamera().GetProjection()[0])

	color := gl.GetUniformLocation(program.id, gl.Str("color\x00"))
	whiteish := (&Color{.8, .8, .8, 1}).Vec4()
	gl.Uniform4fv(color, 1, &whiteish[0])

	gl.DrawArrays(gl.LINE_STRIP, 0, int32(p.Size()))
}

type zNode struct {
	node *Node
	mr   *MeshRenderer
	z    float32
}

type frontToBack []zNode

func (a frontToBack) Len() int           { return len(a) }
func (a frontToBack) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a frontToBack) Less(i, j int) bool { return a[i].z > a[j].z }

type backToFront []zNode

func (a backToFront) Len() int           { return len(a) }
func (a backToFront) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a backToFront) Less(i, j int) bool { return a[i].z < a[j].z }

func getMeshRenderer(node *Node) *MeshRenderer {
	var mr *MeshRenderer
	var ok bool
	for i := range node.components {
		if mr, ok = node.components[i].(*MeshRenderer); ok {
			break
		}
	}

	return mr
}

func opaqueFrontToBack(sg *SceneGraph, cameraTransform *math.Mat4) []zNode {
	var nodes []zNode
	for g := range sg.Traverse() {
		node, ok := g.(*Node)
		if !ok {
			continue
		}

		// If the material needs blending, we can't draw it in this pass. We'll have to
		// draw it back to front
		mr := getMeshRenderer(node)
		if mr == nil {
			continue
		}
		if mr.material.GetBlending().Enabled {
			continue
		}

		// Compute world position of the node.
		origin := math.Vec4{0, 0, 0, 1}
		position := node.worldTransform.AsMat4().Mul4x1(&origin)

		// And get the the node position from the camera pov.
		transformed := cameraTransform.Mul4x1(&position)

		nodes = append(nodes, zNode{
			node: node,
			mr:   mr,
			z:    transformed.Z(),
		})
	}

	// Sort the nodes by z
	sort.Sort(frontToBack(nodes))

	return nodes
}

// Compute the camera transform: projection . worldTransform^-1.
func cameraTransform(c Camera) *math.Mat4 {
	cameraTransform := *c.GetProjection()

	// The camera may either be part of the scene (part of the scene graph) or not.
	// XXX: we don't check that the root of the tree the Camera is part of is
	// indeed the scenegraph we are drawing.
	cameraNode := c.AsNode()
	var cameraWorldInverse math.Mat4
	if cameraNode.parent == nil {
		cameraWorldInverse = cameraNode.GetTransform().Inverse()
	} else {
		cameraWorldInverse = cameraNode.worldTransform.AsMat4().Inverse()
	}
	cameraTransform.Mul4With(&cameraWorldInverse)
	return &cameraTransform
}

func createUniformUploader(program *glProgram, uniform Uniform) glUploader {
	name := uniform.Name()
	location := gl.GetUniformLocation(program.id, gl.Str(name+"\x00"))
	if location == -1 {
		panic("no uniform called " + name)
	}

	switch name {
	case "mvp":
		return &glUniformMVP{
			uniform:  uniform,
			location: location,
		}
	default:
		return &glUniformUser{
			uniform:  uniform,
			location: location,
		}
	}
}

func collectUniforms(program *glProgram, uniforms []Uniform) {
	for _, u := range uniforms {
		uploader := createUniformUploader(program, u)
		program.uploaders = append(program.uploaders, uploader)
	}
}

func (r *renderer) programForMaterial(m Material) *glProgram {
	if p, ok := r.programs[m.ID()]; ok {
		return p
	}

	vs := NewVertexShader(vertexShader)
	fs := m.GetFragmentShader()
	p, err := makeProgram(vs, fs)
	if err != nil {
		panic(err)
	}
	program := &glProgram{
		id: p,
		vs: vs,
		fs: fs,
	}

	gl.UseProgram(p)

	// Cache uniform locations
	collectUniforms(program, vs.uniforms)
	collectUniforms(program, fs.uniforms)

	r.programs[m.ID()] = program
	return program
}

func (r *renderer) drawSceneGraph(fb Framebuffer, sg *SceneGraph) {
	c := fb.GetCamera()

	// Update all world transform matrices.
	sg.updateWorldTransform()

	// Render opaque geometry, front to back to limit overdraw thanks to early z
	// discard.
	cameraTransform := cameraTransform(c)
	nodes := opaqueFrontToBack(sg, cameraTransform)
	for i := range nodes {
		node := &nodes[i]

		// cameraTransform * node.worldTransform
		mvp := &math.Mat4{}

		mesh := node.mr.mesher.GetMesh()
		vao := newVAOFromMesh(mesh)
		defer vao.destroy()
		vao.bind()

		program := r.programForMaterial(node.mr.material)
		gl.UseProgram(program.id)

		// Upload each attribute buffer and link them to the vertex shader.
		vsAttr := &r.vs.attributes[0]
		for i := range vao.vbos {
			vbo := vao.vbos[i]
			ab := vbo.buffer
			if ab.Name != vsAttr.name {
				continue
			}

			vbo.upload()

			location := gl.GetAttribLocation(program.id, gl.Str(vsAttr.name+"\x00"))
			if location == -1 {
				// XXX: reports errors better
				fmt.Fprintf(os.Stderr, "couldn't find attribute %s", vsAttr.name)
			}

			index := uint32(location)
			gl.EnableVertexAttribArray(index)
			gl.VertexAttribPointer(index, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
		}

		// Upload indices.
		vao.indices.upload()

		// Upload uniforms
		mvp.Mul4Of(cameraTransform, node.node.worldTransform.AsMat4())
		location := gl.GetUniformLocation(program.id, gl.Str("mvp\x00"))
		gl.UniformMatrix4fv(location, 1, false, &mvp[0])

		color := gl.GetUniformLocation(program.id, gl.Str("color\x00"))
		whiteish := (&Color{.8, .8, .8, 1}).Vec4()
		gl.Uniform4fv(color, 1, &whiteish[0])

		// Draw. The index array is already bound above.
		gl.DrawElements(
			glVertexMode(mesh.GetVertexMode()),
			int32(mesh.indices.Len()),
			glIndexType(&mesh.indices),
			gl.PtrOffset(0))
	}
}
