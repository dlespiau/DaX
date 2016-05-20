package dax

import (
	"testing"

	math "github.com/dlespiau/dax/math"
	"github.com/stretchr/testify/assert"
)

func testEqualF(t *testing.T, expected, value float32, epsilon float64) {
	if math.Abs(expected) < 1e-3 {
		assert.True(t, math.Abs(expected-value) < float32(epsilon))
	} else {
		t.Logf("Expected %f, value %f\n", expected, value)
		t.Logf("ratio %f\n", (expected-value)/expected)
		assert.InEpsilon(t, expected, value, epsilon)
	}
}

func TestHSL(t *testing.T) {
	tests := []struct {
		color   string
		r, g, b uint8
		h, s, l float32
	}{
		{"#000000", 0, 0, 0, 0, 0, 0},
		{"#FFFFFF", 255, 255, 255, 0, 0, 1},
		{"#FF0000", 255, 0, 0, 0, 1, 1. / 2},
		{"#00FF00", 0, 255, 0, 120. / 360, 1, 1. / 2},
		{"#0000FF", 0, 0, 255, 240. / 360, 1, 1. / 2},
		{"#FFFF00", 255, 255, 0, 60. / 360, 1, 1. / 2},
		{"#00FFFF", 0, 255, 255, 180. / 360, 1, 1. / 2},
		{"#FF00FF", 255, 0, 255, 300. / 360, 1, 1. / 2},
		{"#C0C0C0", 192, 192, 192, 0, 0, 3. / 4},
		{"#808080", 128, 128, 128, 0, 0, 1. / 2},
		{"#800000", 128, 0, 0, 0, 1, 1. / 4},
		{"#808000", 128, 128, 0, 60. / 360, 1, 1. / 4},
		{"#008000", 0, 128, 0, 120. / 360, 1, 1. / 4},
		{"#800080", 128, 0, 128, 300. / 360, 1, 1. / 4},
		{"#008080", 0, 128, 128, 180. / 360, 1, 1. / 4},
		{"#000080", 0, 0, 128, 240. / 360, 1, 1. / 4},
	}

	for i, test := range tests {
		var color Color

		color.FromRGBu8(test.r, test.g, test.b)
		h, s, l := color.ToHSL()
		t.Logf("Test case #%d", i)
		assert.Equal(t, test.h, h)
		assert.Equal(t, test.s, s)
		testEqualF(t, test.l, l, 1e-2)
	}

	for i, test := range tests {
		var color Color

		color.FromHSL(test.h, test.s, test.l)
		t.Logf("Test case #%d", i)
		testEqualF(t, color.R, float32(test.r)/255, 1e-2)
		testEqualF(t, color.G, float32(test.g)/255, 1e-2)
		testEqualF(t, color.B, float32(test.b)/255, 1e-2)
		assert.Equal(t, float32(1.0), color.A, 1e-2)
	}
}
