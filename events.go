package dax

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

type MouseButton glfw.MouseButton

// Mouse buttons
const (
	MouseButton1      MouseButton = MouseButton(glfw.MouseButton1)
	MouseButton2      MouseButton = MouseButton(glfw.MouseButton2)
	MouseButton3      MouseButton = MouseButton(glfw.MouseButton3)
	MouseButton4      MouseButton = MouseButton(glfw.MouseButton4)
	MouseButton5      MouseButton = MouseButton(glfw.MouseButton5)
	MouseButton6      MouseButton = MouseButton(glfw.MouseButton6)
	MouseButton7      MouseButton = MouseButton(glfw.MouseButton7)
	MouseButton8      MouseButton = MouseButton(glfw.MouseButton8)
	MouseButtonLast   MouseButton = MouseButton(glfw.MouseButtonLast)
	MouseButtonLeft   MouseButton = MouseButton(glfw.MouseButtonLeft)
	MouseButtonRight  MouseButton = MouseButton(glfw.MouseButtonRight)
	MouseButtonMiddle MouseButton = MouseButton(glfw.MouseButtonMiddle)
)

func (b MouseButton) String() string {
	switch b {
	case MouseButtonLeft:
		return "left"
	case MouseButtonRight:
		return "right"
	case MouseButtonMiddle:
		return "middle"
	}
	return string(b)
}
