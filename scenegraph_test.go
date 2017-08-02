package dax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	var a, b, c, d, e, f, g, h, i Node
	sg := NewSceneGraph()
	preOrder := [...]Grapher{sg, &f, &b, &a, &d, &c, &e, &g, &i, &h}

	sg.AddChild(&f)
	f.AddChild(&b)
	f.AddChild(&g)
	b.AddChild(&a)
	b.AddChild(&d)
	d.AddChild(&c)
	d.AddChild(&e)
	g.AddChild(&i)
	i.AddChild(&h)

	idx := 0
	for n := range sg.Traverse() {
		assert.Equal(t, n, preOrder[idx])
		idx++
	}
	assert.Equal(t, len(preOrder), idx)
}
