package dax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChild(t *testing.T) {
	var p, c Node

	p.AddChild(&c)
	assert.Equal(t, c.Parent, &p)
	assert.Equal(t, len(p.Children), 1)
	assert.Equal(t, p.Children[0], &c)
}
