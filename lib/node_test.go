package dax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChild(t *testing.T) {
	p := NewNode()
	c := NewNode()

	p.AddChild(c)
	assert.Equal(t, c.GetParent(), p)

	children := p.GetChildren()
	assert.Equal(t, len(children), 1)
	assert.Equal(t, children[0], c)
}
