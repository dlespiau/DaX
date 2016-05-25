package math

var (
	// MinNormal is the smallest normal value possible.
	MinNormal = float32(1.1754943508222875e-38) // 1 / 2**(127 - 1)
	// MinValue is the smallest non zero value possible.
	MinValue = float32(SmallestNonzeroFloat32)
	// MaxValue is the highest value a float32 can have.
	MaxValue = float32(MaxFloat32)

	// InfPos is the positive infinity value.
	InfPos = float32(Inf(1))
	// InfNeg is the positive infinity value.
	InfNeg = float32(Inf(-1))
)

// Epsilon is some tiny value that determines how precisely equal we want our
// floats to be. This is exported and left as a variable in case you want to
// change the default threshold for the purposes of certain methods (e.g.
// Unproject uses the default epsilon when determining if the determinant is
// "close enough" to zero to mean there's no inverse).
//
// This is, obviously, not mutex protected so be **absolutely sure** that no
// functions using Epsilon are being executed when you change this.
var Epsilon float32 = 1e-10

// FloatEqual is a safe utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
func FloatEqual(a, b float32) bool {
	return FloatEqualThreshold(a, b, Epsilon)
}

// FloatEqualThreshold is a utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
//
// This differs from FloatEqual in that it lets you pass in your comparison threshold, so that you can adjust the comparison value to your specific needs
func FloatEqualThreshold(a, b, epsilon float32) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	}

	diff := Abs(a - b)
	if a*b == 0 || diff < MinNormal { // If a or b are 0 or both are extremely close to it
		return diff < epsilon*epsilon
	}

	// Else compare difference
	return diff/(Abs(a)+Abs(b)) < epsilon
}

// IsClamped checks if a is clamped between low and high as if
// Clamp(a, low, high) had been called.
//
// In most cases it's probably better to just call Clamp
// without checking this since it's relatively cheap.
func IsClamped(a, low, high float32) bool {
	return a >= low && a <= high
}

// SetMin sets a to the Min(a, b).
func SetMin(a, b *float32) {
	if *b < *a {
		*a = *b
	}
}

// SetMax sets a to the Max(a, b).
func SetMax(a, b *float32) {
	if *a < *b {
		*a = *b
	}
}

// Round shortens a float32 value to a specified precision (number of digits
// after the decimal point) with "round half up" tie-braking rule. Half-way
// values (23.5) are always rounded up (24).
func Round(v float32, precision int) float32 {
	p := float32(precision)
	t := v * Pow(10, p)
	if t > 0 {
		return Floor(t+0.5) / Pow(10, p)
	}
	return Ceil(t-0.5) / Pow(10, p)
}
