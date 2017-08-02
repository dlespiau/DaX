package dax

// MeshRenderer is a component rendering a Mesh with a Material.
type MeshRenderer struct {
	mesher   Mesher
	material Material
}

// NewMeshRenderer creates a new MeshRenderer.
func NewMeshRenderer(mesher Mesher, material Material) *MeshRenderer {
	return &MeshRenderer{
		mesher:   mesher,
		material: material,
	}
}

// Update implements Updater for MeshRenderer.
func (mr *MeshRenderer) Update(time float64) {

}

// Draw implements Drawer for MeshRenderer.
func (mr *MeshRenderer) Draw() {

}
