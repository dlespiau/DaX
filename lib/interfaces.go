package dax

// Mesher is an object that can produce a Mesh.
type Mesher interface {
	GetMesh() *Mesh
}

// Drawer is an object that can draw on a Framebuffer.
type Drawer interface {
	Draw(fb Framebuffer)
}

// Getter gets a value.
type Getter interface {
	Get() interface{}
}

// Setter sets a value.
type Setter interface {
	Set(interface{})
}

// Namer is an object with a name.
type Namer interface {
	Name() string
}
