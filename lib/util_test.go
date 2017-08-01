package dax

import (
	"fmt"
	"os"
	"testing"

	"github.com/dlespiau/dax/math"
	"github.com/stretchr/testify/assert"
)

func assertFloat(t *testing.T, exp, val, epsilon float32) {
	if math.FloatEqualThreshold(exp, val, epsilon) {
		return
	}

	fmt.Fprintf(os.Stderr, "Expected %v, got %v\n", exp, val)
	assert.True(t, false)
}

func assertVec2(t *testing.T, exp, val *math.Vec2, epsilon float32) {
	if math.FloatEqualThreshold(exp[0], val[0], epsilon) &&
		math.FloatEqualThreshold(exp[1], val[1], epsilon) {
		return
	}

	fmt.Fprintf(os.Stderr, "Expected %v, got %v\n", exp, val)
	assert.True(t, false)
}

func assertVec3(t *testing.T, exp, val *math.Vec3, epsilon float32) {
	if math.FloatEqualThreshold(exp[0], val[0], epsilon) &&
		math.FloatEqualThreshold(exp[1], val[1], epsilon) &&
		math.FloatEqualThreshold(exp[2], val[2], epsilon) {
		return
	}

	fmt.Fprintf(os.Stderr, "Expected %v, got %v\n", exp, val)
	assert.True(t, false)
}

func assertVec4(t *testing.T, exp, val *math.Vec4, epsilon float32) {
	if math.FloatEqualThreshold(exp[0], val[0], epsilon) &&
		math.FloatEqualThreshold(exp[1], val[1], epsilon) &&
		math.FloatEqualThreshold(exp[2], val[2], epsilon) &&
		math.FloatEqualThreshold(exp[3], val[3], epsilon) {
		return
	}

	fmt.Fprintf(os.Stderr, "Expected %v, got %v\n", exp, val)
	assert.True(t, false)
}

func assertQuat(t *testing.T, exp, val *math.Quaternion, epsilon float32) {
	if exp.EqualThreshold(val, epsilon) {
		return
	}

	fmt.Fprintf(os.Stderr, "Expected %v, got %v\n", exp, val)
	assert.True(t, false)
}
