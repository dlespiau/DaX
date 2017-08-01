package material

import dax "github.com/dlespiau/dax/lib"

// Color is the simplest material possible, just a color.
type Color struct {
	dax.BaseMaterial
	color dax.Color
}

var _ dax.Material = &Color{}

// NewColor creates a new Color material.
func NewColor(color *dax.Color) *Color {
	return &Color{
		color: *color,
	}
}

const colorFragmentShader = `
#version 330
uniform vec4 color;
out vec4 outputColor;
void main() {
    outputColor = color;
}`

// ID is part of the Material interface.
func (m *Color) ID() string {
	return "-dax-material-color"
}

// GetFragmentShader is part of the Material interface.
func (m *Color) GetFragmentShader() *dax.FragmentShader {
	s := dax.NewFragmentShader(colorFragmentShader)
	s.AddUniform(dax.VariableKindVec4, "color")

	return s
}
