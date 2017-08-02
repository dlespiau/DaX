package main

import (
	"fmt"
	"strings"

	dax "github.com/dlespiau/dax/lib"
)

// Category is a way to group examples by common area.
type Category int

const (
	// CategoryWinsys is for window system examples.
	CategoryWinsys Category = iota
	// CategoryGraphics is for examples drawing things on the screen.
	CategoryGraphics
)

// String implements Stringer for Category
func (cat Category) String() string {
	switch cat {
	case CategoryWinsys:
		return "winsys"
	case CategoryGraphics:
		return "gfx"
	}
	return fmt.Sprintf("Category(%d)", cat)
}

// Example defines a Scene showing off the usage of the DaX API.
type Example struct {
	Category    Category
	Name        string
	Description string
	Scene       dax.Scener
}

// ID provides a string to identify an example.
func (e *Example) ID() string {
	category := strings.ToLower(e.Category.String())
	name := strings.ToLower(strings.Replace(e.Name, " ", "-", -1))
	return category + "-" + name
}
