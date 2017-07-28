package dax

import "github.com/dlespiau/dax/math"

// UniformKind defines the kind of uniform.
type UniformKind int

const (
	// UniformKindFloat is a float uniform.
	UniformKindFloat UniformKind = iota
	// UniformKindVec2 is a vec2 uniform.
	UniformKindVec2
	// UniformKindVec3 is a vec3 uniform.
	UniformKindVec3
	// UniformKindVec4 is a vec4 uniform.
	UniformKindVec4
	// UniformKindMatrix is a matrix uniform.
	UniformKindMatrix
	uniformKindMax
)

// Uniform is a shader parameter, constant per draw call.
type Uniform interface {
	Namer
	Kind() UniformKind
	Getter
	Setter
}

type baseUniform struct {
	name string // NUL-terminated string (GL drivers needs C strings)
	kind UniformKind
}

// Name returns the name of the uniform.
func (u *baseUniform) Name() string {
	return u.name[:len(u.name)-1]
}

// Kind returns the kind of uniform
func (u *baseUniform) Kind() UniformKind {
	return u.kind
}

type floatUniform struct {
	baseUniform
	val float32
}

func (u *floatUniform) Get() interface{} {
	return u.val
}

func (u *floatUniform) Set(v interface{}) {
	u.val = v.(float32)
}

type vec2Uniform struct {
	baseUniform
	val math.Vec2
}

func (u *vec2Uniform) Get() interface{} {
	return u.val
}

func (u *vec2Uniform) Set(v interface{}) {
	u.val = v.(math.Vec2)
}

type vec3Uniform struct {
	baseUniform
	val math.Vec3
}

func (u *vec3Uniform) Get() interface{} {
	return u.val
}

func (u *vec3Uniform) Set(v interface{}) {
	u.val = v.(math.Vec3)
}

type vec4Uniform struct {
	baseUniform
	val math.Vec4
}

func (u *vec4Uniform) Get() interface{} {
	return u.val
}

func (u *vec4Uniform) Set(value interface{}) {
	switch v := value.(type) {
	case Color:
		u.val = v.Vec4()
	default:
		u.val = v.(math.Vec4)
	}
}

type matrixUniform struct {
	baseUniform
	val math.Mat4
}

func (u *matrixUniform) Get() interface{} {
	return u.val
}

func (u *matrixUniform) Set(v interface{}) {
	u.val = v.(math.Mat4)
}

func createUniform(kind UniformKind, name string) Uniform {
	var u Uniform

	switch kind {
	case UniformKindFloat:
		u = &floatUniform{
			baseUniform: baseUniform{
				kind: UniformKindFloat,
				name: name + "\x00",
			},
		}
	case UniformKindVec2:
		u = &vec2Uniform{
			baseUniform: baseUniform{
				kind: UniformKindVec2,
				name: name + "\x00",
			},
		}
	case UniformKindVec3:
		u = &vec3Uniform{
			baseUniform: baseUniform{
				kind: UniformKindVec3,
				name: name + "\x00",
			},
		}
	case UniformKindVec4:
		u = &vec4Uniform{
			baseUniform: baseUniform{
				kind: UniformKindVec4,
				name: name + "\x00",
			},
		}
	case UniformKindMatrix:
		u = &matrixUniform{
			baseUniform: baseUniform{
				kind: UniformKindMatrix,
				name: name + "\x00",
			},
		}
	}

	return u
}

type baseShader struct {
	source   string // NUL-terminated string (GL drivers needs C strings)
	uniforms []Uniform
}

// Uniform returns the uniform named name.
func (s *baseShader) Uniform(name string) Uniform {
	for i := range s.uniforms {
		if name == s.uniforms[i].Name() {
			return s.uniforms[i]
		}
	}
	return nil
}

// AddUniform adds a uniform to a shader.
func (s *baseShader) AddUniform(kind UniformKind, name string) Uniform {
	u := createUniform(kind, name)
	s.uniforms = append(s.uniforms, u)
	return u
}

// VertexShader is a program that runs for each vertex.
type VertexShader struct {
	baseShader
}

// NewVertexShader creates a vertex shader.
func NewVertexShader(source string) *VertexShader {
	return &VertexShader{
		baseShader: baseShader{
			source: source + "\x00",
		},
	}
}

// FragmentShader is a program that runs for each fragment.
type FragmentShader struct {
	baseShader
}

// NewFragmentShader creates a fragment shader.
func NewFragmentShader(source string) *FragmentShader {
	return &FragmentShader{
		baseShader: baseShader{

			source: source + "\x00",
		},
	}
}
