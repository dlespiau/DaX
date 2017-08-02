package dax

import "github.com/dlespiau/dax/math"

// VariableKind defines the type of a variable in a shader.
type VariableKind int

const (
	// VariableKindFloat is a float uniform.
	VariableKindFloat VariableKind = iota
	// VariableKindVec2 is a vec2 uniform.
	VariableKindVec2
	// VariableKindVec3 is a vec3 uniform.
	VariableKindVec3
	// VariableKindVec4 is a vec4 uniform.
	VariableKindVec4
	// VariableKindMat4 is a 4x4 matrix uniform.
	VariableKindMat4
	variableKindMax
)

type baseVariable struct {
	name string
	kind VariableKind
}

// Name returns the name of the uniform.
func (u *baseVariable) Name() string {
	return u.name
}

// Kind returns the kind of uniform
func (u *baseVariable) Kind() VariableKind {
	return u.kind
}

// Attribute is per-vertex data as input to the vertex shader.
type Attribute struct {
	baseVariable
}

// Uniform is a shader parameter, constant per draw call.
type Uniform interface {
	Namer
	Kind() VariableKind
	Getter
	Setter
}

// builtin uniforms are those magical uniforms we'll detect and upload
// automatically in the core rendering engine.

var builtinUniforms = [...]string{
	"mvp",
}

type builtinUniform struct {
	baseVariable
}

func (u *builtinUniform) Get() interface{} {
	panic("can't get builtin uniforms")
}

func (u *builtinUniform) Set(v interface{}) {
	panic("can't set builtin uniforms")
}

type floatUniform struct {
	baseVariable
	val float32
}

func (u *floatUniform) Get() interface{} {
	return u.val
}

func (u *floatUniform) Set(v interface{}) {
	u.val = v.(float32)
}

type vec2Uniform struct {
	baseVariable
	val math.Vec2
}

func (u *vec2Uniform) Get() interface{} {
	return u.val
}

func (u *vec2Uniform) Set(v interface{}) {
	u.val = v.(math.Vec2)
}

type vec3Uniform struct {
	baseVariable
	val math.Vec3
}

func (u *vec3Uniform) Get() interface{} {
	return u.val
}

func (u *vec3Uniform) Set(v interface{}) {
	u.val = v.(math.Vec3)
}

type vec4Uniform struct {
	baseVariable
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

type mat4Uniform struct {
	baseVariable
	val math.Mat4
}

func (u *mat4Uniform) Get() interface{} {
	return u.val
}

func (u *mat4Uniform) Set(v interface{}) {
	u.val = v.(math.Mat4)
}

func createUniform(kind VariableKind, name string) Uniform {
	var u Uniform

	// Builtin uniforms are a bit special.
	for i := range builtinUniforms {
		if builtinUniforms[i] == name {
			return &builtinUniform{
				baseVariable: baseVariable{
					kind: kind,
					name: name,
				},
			}
		}
	}

	switch kind {
	case VariableKindFloat:
		u = &floatUniform{
			baseVariable: baseVariable{
				kind: VariableKindFloat,
				name: name,
			},
		}
	case VariableKindVec2:
		u = &vec2Uniform{
			baseVariable: baseVariable{
				kind: VariableKindVec2,
				name: name,
			},
		}
	case VariableKindVec3:
		u = &vec3Uniform{
			baseVariable: baseVariable{
				kind: VariableKindVec3,
				name: name,
			},
		}
	case VariableKindVec4:
		u = &vec4Uniform{
			baseVariable: baseVariable{
				kind: VariableKindVec4,
				name: name,
			},
		}
	case VariableKindMat4:
		u = &mat4Uniform{
			baseVariable: baseVariable{
				kind: VariableKindMat4,
				name: name,
			},
		}
	}

	return u
}

type baseShader struct {
	source   string
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
func (s *baseShader) AddUniform(kind VariableKind, name string) Uniform {
	u := createUniform(kind, name)
	s.uniforms = append(s.uniforms, u)
	return u
}

// VertexShader is a program that runs for each vertex.
type VertexShader struct {
	baseShader
	attributes []Attribute
}

// NewVertexShader creates a vertex shader.
func NewVertexShader(source string) *VertexShader {
	return &VertexShader{
		baseShader: baseShader{
			source: source,
		},
	}
}

// AddAttribute adds an attribute to a vertex shader.
func (vs *VertexShader) AddAttribute(kind VariableKind, name string) *Attribute {
	a := Attribute{
		baseVariable: baseVariable{
			kind: kind,
			name: name,
		},
	}
	vs.attributes = append(vs.attributes, a)
	return &a
}

// FragmentShader is a program that runs for each fragment.
type FragmentShader struct {
	baseShader
}

// NewFragmentShader creates a fragment shader.
func NewFragmentShader(source string) *FragmentShader {
	return &FragmentShader{
		baseShader: baseShader{

			source: source,
		},
	}
}
