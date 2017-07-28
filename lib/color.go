package dax

import (
	math "github.com/dlespiau/dax/math"
)

// Color represents a color encoded in RGBA.
type Color struct {
	R, G, B, A float32
}

// FromRGBA initializes a color from (r,g,b,a) values. Components should be
// between 0 and 1.
func (color *Color) FromRGBA(r, g, b, a float32) {
	color.R = r
	color.G = g
	color.B = b
	color.A = a
}

// FromRGB initializes a color from (r,g,b) values. Components should be
// between 0 and 1. The alpha is initiazed to 1 (fully opaque).
func (color *Color) FromRGB(r, g, b float32) {
	color.FromRGBA(r, g, b, 1.0)
}

func u8toF(x uint8) float32 {
	return float32(x) / 255.
}

// FromRGBAu8 initializes a color from (r,g,b,a) values. Components should be
// between 0 and 255.
func (color *Color) FromRGBAu8(r, g, b, a uint8) {
	color.FromRGBA(u8toF(r), u8toF(g), u8toF(b), u8toF(a))
}

// FromRGBu8 initializes a color from (r,g,b) values. Components should be
// between 0 and 255. The alpha is initiazed to 1 (ie 255, fully opaque).
func (color *Color) FromRGBu8(r, g, b uint8) {
	color.FromRGBAu8(r, g, b, 255)
}

// XXX: The RGB <-> HSL functions could do with some benchmarking and ideas
// from: http://lolengine.net/blog/2013/01/13/fast-rgb-to-hsv
func hue2rgb(p, q, t float32) float32 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1./6 {
		return p + (q-p)*6*t
	}
	if t < 1./2 {
		return q
	}
	if t < 2./3 {
		return p + (q-p)*(2./3-t)*6
	}
	return p
}

// FromHSL initializes a color from (h,s,l) values. h, s, l are between 0 and 1.
// Conversion formula is adapted from:
// http://en.wikipedia.org/wiki/HSL_color_space
func (color *Color) FromHSL(h, s, l float32) {
	color.A = 1.0

	if s == 0 {
		// achromatic
		color.R = l
		color.G = l
		color.B = l
		return
	}

	var q float32
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	color.R = hue2rgb(p, q, h+1./3)
	color.G = hue2rgb(p, q, h)
	color.B = hue2rgb(p, q, h-1./3)
}

// ToHSL converts a color to HSL. Returned (h,s,l) components are between 0 and
// 1. Conversion formula adapted from:
// http://en.wikipedia.org/wiki/HSL_color_space.
func (color *Color) ToHSL() (h, s, l float32) {
	r := color.R
	g := color.G
	b := color.B

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))

	h = float32(0.0)
	s = float32(0.0)
	l = (max + min) / 2

	if max == min {
		return
	}

	d := max - min
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}

	switch max {
	case r:
		var k float32
		if g < b {
			k = float32(6)
		} else {
			k = float32(0)
		}
		h = (g-b)/d + k
	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}
	h /= 6

	return
}
