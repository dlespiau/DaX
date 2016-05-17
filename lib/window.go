package dax

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Window struct {
	app           *Application
	name          string
	width, height int
	fb            Framebuffer
	scene         Scener
	glfwWindow    *glfw.Window
}

func newWindow(app *Application, name string, width, height int) *Window {
	window := new(Window)
	window.app = app
	window.name = name
	window.width = width
	window.height = height

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfwWindow, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		panic(err)
	}
	window.glfwWindow = glfwWindow
	glfwWindow.MakeContextCurrent()

	glfw.SwapInterval(1)

	// create OnScreen object
	window.fb = newOnScreen(width, height)

	// window events
	glfwWindow.SetCloseCallback(onClose)
	glfwWindow.SetSizeCallback(onResize)

	// key events
	glfwWindow.SetKeyCallback(onKeyEvent)
	glfwWindow.SetCharCallback(onRuneEvent)

	// mouse events
	glfwWindow.SetMouseButtonCallback(OnMouseButton)
	glfwWindow.SetCursorPosCallback(OnMouseMoved)

	// Install the default scene
	window.SetScene(new(Scene))

	return window
}

func (w *Window) Update() {
	w.scene.Update()
}

func (w *Window) Draw() {
	c := w.scene.BackgroundColor()

	gl.ClearColor(c.r, c.g, c.b, c.a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	w.scene.Draw(w.fb)
}

func (w *Window) Close() {
	w.glfwWindow.SetShouldClose(true)
}

func onResize(w *glfw.Window, width, height int) {
	window := getWindow(w)
	window.width = width
	window.height = height
	window.scene.OnResize(window.fb, width, height)
}

func onClose(w *glfw.Window) {
	window := getWindow(w)
	window.scene.TearDown()
}

func onKeyEvent(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	window := getWindow(w)

	if action == glfw.Press {
		window.scene.OnKeyPressed()
	} else if action == glfw.Release {
		window.scene.OnKeyReleased()
	}
}

func OnMouseMoved(w *glfw.Window, x, y float64) {
	window := getWindow(w)
	window.scene.OnMouseMoved(float32(x), float32(y))
}

func OnMouseButton(w *glfw.Window, button glfw.MouseButton,
	action glfw.Action, mod glfw.ModifierKey) {
	window := getWindow(w)
	x, y := w.GetCursorPos()
	if action == glfw.Press {
		window.scene.OnMouseButtonPressed(MouseButton(button), float32(x), float32(y))
	} else if action == glfw.Release {
		window.scene.OnMouseButtonReleased(MouseButton(button), float32(x), float32(y))
	}
}

func onRuneEvent(w *glfw.Window, r rune) {
	window := getWindow(w)
	window.scene.OnRuneEntered(r)
}

func (w *Window) SetScene(s Scener) {
	if w.scene != nil {
		w.scene.TearDown()
	}

	if s != nil {
		w.scene = s
	} else {
		// fallback to the default scene, maintaining the invariant
		// that we always have valid scene
		w.scene = new(Scene)
	}
	sceneSetup(w.scene)
	w.scene.OnResize(w.fb, w.width, w.height)
}

func (w *Window) Screenshot() *image.RGBA {
	return w.fb.Screenshot()
}

func (w *Window) ScreenshotToFile(filename string) {
	img := w.fb.Screenshot()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err.Error())
	}
}
