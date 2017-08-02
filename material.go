package dax

// Material is used to paint geometries.
type Material interface {
	// ID is the material ID. This must be unique across all materials.
	ID() string
	// GetFragmentShader returns the fragment shader used by this Material.
	GetFragmentShader() *FragmentShader
	// GetBlending returns the blending state of the Material.
	GetBlending() *Blending
	// GetDepthTest returns the depth test state of the Material.
	GetDepthTest() *DepthTest
}

// BlendingMode is the blending mode of a Material.
type BlendingMode int

const (
	BlendingAdd BlendingMode = iota
	BlendingSubstract
	BlendingReverseSubstract
	BlendingMin
	BlendingMax
)

// BlendingFunc is the blending function of a Material.
type BlendingFunc int

const (
	BlendingOne BlendingFunc = iota
	BlendingZero

	BlendingSrcColor
	BlendingDstColor
	BlendingOneMinusSrcColor
	BlendingOneMinusDstColor

	BlendingSrcAlpha
	BlendingDstAlpha
	BlendingOneMinusSrcAlpha
	BlendingOneMinusDstAlpha

	BlendingConstantColor
	BlendingOneMinusConstantColor
	BlendingConstantAlpha
	BlendingOneMinusConstantAlpha
)

// Blending holds the entire blending state of a Material.
type Blending struct {
	Enabled   bool
	ModeRGB   BlendingMode
	ModeAlpha BlendingMode
	SrcRGB    BlendingFunc
	DstRGB    BlendingFunc
	SrcAlpha  BlendingFunc
	DstAlpha  BlendingFunc
	Color     Color
}

// DepthTestFunc is the depth test function of a Material.
type DepthTestFunc int

const (
	DepthTestNever DepthTestFunc = iota
	DepthTestLess
	DepthTestGreater
	DepthTestEqual
	DepthTestAlways
	DepthTestLessOrEqual
	DepthTestGreaterOrEqual
	DepthTestNotEqual
)

// DepthTest holds the entire depth test state of a Material.
type DepthTest struct {
	Enabled bool
	Write   bool
	Func    DepthTestFunc
}

// BaseMaterial holds the common material state and can be used to implement
// custom materials.
type BaseMaterial struct {
	Blending  Blending
	DepthTest DepthTest
}

// ID is part of the Material interface.
func (m *BaseMaterial) ID() string {
	return "-dax-material-base"
}

// GetFragmentShader is part of the Material interface.
func (m *BaseMaterial) GetFragmentShader() *FragmentShader {
	return NewFragmentShader(`
#version 330
out vec4 outputColor;
void main() {
    outputColor = vec4(1,1,1,1);
}`)
}

// GetBlending is part of the Material interface.
func (m *BaseMaterial) GetBlending() *Blending {
	return &m.Blending
}

// GetDepthTest is part of the Material interface.
func (m *BaseMaterial) GetDepthTest() *DepthTest {
	return &m.DepthTest
}

var _ Material = &BaseMaterial{}
