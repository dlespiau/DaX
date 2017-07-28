package dax

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	polylineMaterial = "dax-polyline"
)

type renderer struct {
	programs map[string]uint32
}

func newRenderer() *renderer {
	return &renderer{
		programs: make(map[string]uint32),
	}
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cSources, free := gl.Strs(source)
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

	return program, nil
}

/*
var vertexShader = `
#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
`
*/

const vertexShader = `
#version 330 core

in vec3 position;

uniform mat4 mvp;

void main(){
	gl_Position = mvp * vec4(position, 1.0f);
}`

const fragmentShader = `
#version 330
uniform vec4 color;
out vec4 outputColor;
void main() {
    outputColor = color;
}`

func (r *renderer) makePolylineProgram() uint32 {
	if p, ok := r.programs[polylineMaterial]; ok {
		return p
	}

	v := NewVertexShader(vertexShader)
	s := NewFragmentShader(fragmentShader)
	p, err := makeProgram(v, s)
	if err != nil {
		panic(err)
	}
	r.programs[polylineMaterial] = p
	return p
}

func (r *renderer) drawPolyline(fb Framebuffer, p *Polyline) {
	if p.Size() == 0 {
		return
	}

	program := r.makePolylineProgram()

	mesh := NewMesh()
	defer mesh.Destroy()
	mesh.AddAttribute("position", p.vertices, 3)
	mesh.Bind()

	gl.UseProgram(program)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	position := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	mvp := gl.GetUniformLocation(program, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvp, 1, false, &fb.Projection()[0])

	color := gl.GetUniformLocation(program, gl.Str("color\x00"))
	whiteish := (&Color{.8, .8, .8, 1}).Vec4()
	gl.Uniform4fv(color, 1, &whiteish[0])

	gl.DrawArrays(gl.LINE_STRIP, 0, int32(p.Size()))
}

func (r *renderer) drawSceneGraph(fb Framebuffer, sg *SceneGraph) {

}
