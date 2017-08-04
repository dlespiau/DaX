package dax

// PropertyFlag defines flags that can be applied to properties.
type PropertyFlag uint64

const ()

// Property is a generic interface to a paramter. This parameter has a name and
// a value that can be queried and set.
type Property interface {
	Namer
	Is(PropertyFlag) bool
	Set(value interface{})
	Get() interface{}
}

// Tweakable is an objects with properties that can be tweaked.
type Tweakable interface {
	GetProperties() []Property
}

type baseProperty struct {
	name string
}

// Name returns the name of the property.
func (p *baseProperty) Name() string {
	return p.name
}

// FloatProperty is a property holding a float32.
type FloatProperty struct {
	baseProperty
	Value float32
}

// ColorProperty is a Property holding a Color.
type ColorProperty struct {
	baseProperty
	Value Color
}

// Init initialize a ColorPoperty.
func (p *ColorProperty) Init(name string, color *Color) *ColorProperty {
	p.name = name
	p.Value = *color
	return p
}
