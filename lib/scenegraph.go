package dax

type Grapher interface {
	SetParent(parent Grapher)
	GetParent() Grapher
	AddChild(child Grapher)
	GetChildren() []Grapher
}

type nodeStack struct {
	nodes []Grapher
}

func (s *nodeStack) Init() {
	s.nodes = make([]Grapher, 0, 16)
}

func (s *nodeStack) Empty() bool {
	return len(s.nodes) == 0
}

func (s *nodeStack) Push(n Grapher) {
	s.nodes = append(s.nodes, n)
}

func (s *nodeStack) Pop() Grapher {
	i := len(s.nodes) - 1
	n := s.nodes[i]
	s.nodes[i] = nil
	s.nodes = s.nodes[:i]
	return n
}

type SceneGraph struct {
	Node
}

func NewSceneGraph() *SceneGraph {
	sg := new(SceneGraph)
	sg.Init()
	return sg
}

func (sg *SceneGraph) Init() {
	sg.Node.Init()
}

func (sg *SceneGraph) updateWorldTransform() {
	sg.Node.updateWorldTransform(false)
}

func (sg *SceneGraph) Update(time float64) {
	sg.updateWorldTransform()
}

// Depth-first pre-order traversal of the SceneGraph
func (sg *SceneGraph) Traverse() <-chan Grapher {
	ch := make(chan Grapher)

	go func() {
		var stack nodeStack

		stack.Init()
		stack.Push(sg)

		for !stack.Empty() {
			n := stack.Pop()
			ch <- n
			children := n.GetChildren()
			for i := len(children) - 1; i >= 0; i-- {
				stack.Push(children[i])
			}
		}
		close(ch)
	}()

	return ch
}

func (sg *SceneGraph) Draw(fb Framebuffer) {
	fb.render().drawSceneGraph(fb, sg)
}
