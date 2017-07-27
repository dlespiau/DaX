package dax

// Mesher is an object that can produce a Mesh.
type Mesher interface {
	GetMesh() *Mesh
}

// Drawer is an object that can draw on a Framebuffer.
type Drawer interface {
	Draw(fb Framebuffer)
}
