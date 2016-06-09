package dax

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type renderer struct {
	program uint32
}

func newRenderer() *renderer {
	r := new(renderer)

	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	r.program = program

	return r
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
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

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
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
}
` + "\x00"

/*
var fragmentShader = `
#version 330
uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;
void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "0x00"
*/

const fragmentShader = `
#version 330
out vec4 outputColor;
void main() {
    outputColor = vec4(.8, .8, .8, 1);
}
` + "\x00"

func (r *renderer) drawPolyline(fb Framebuffer, p *Polyline) {
	if p.Size() == 0 {
		return
	}

	mesh := NewMesh()
	defer mesh.Destroy()
	mesh.AddAttribute("position", p.vertices, 3)
	mesh.Bind()

	gl.UseProgram(r.program)

	gl.BindFragDataLocation(r.program, 0, gl.Str("outputColor\x00"))

	position := uint32(gl.GetAttribLocation(r.program, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	mvp := gl.GetUniformLocation(r.program, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvp, 1, false, &fb.Projection()[0])

	gl.DrawArrays(gl.LINE_STRIP, 0, int32(p.Size()))
}

func (r *renderer) drawSceneGraph(fb Framebuffer, sg *SceneGraph) {

}
