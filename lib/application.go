package dax

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw: ", err)
	}

	if err := gl.Init(); err != nil {
		panic(err)
	}
}

type application struct {
	Name string

	windows map[*glfw.Window]*Window
}

// NewApplication creates an Application. The Application object is the top
// level object from which everything else in DaX is derived.
func NewApplication(name string) *application {
	app := new(application)
	app.Name = name
	app.windows = make(map[*glfw.Window]*Window)
	return app
}

var _app *application

func (app *application) addWindow(window *Window) {
	// TODO: support multiple windows
	app.windows[window.glfwWindow] = window
}

// Run enters the application main loop.
func (app *application) Run() {
	for _, window := range app.windows {
		for !window.glfwWindow.ShouldClose() {
			window.Update()
			window.Draw()
			window.glfwWindow.SwapBuffers()
			glfw.PollEvents()
		}
	}
}

// CreateWindow creates a window on which scene will be drawn.
func (app *application) CreateWindow(name string, width, height int) *Window {
	_app = app

	window := newWindow(app, name, width, height)
	app.addWindow(window)

	return window
}

func getWindow(w *glfw.Window) *Window {
	return _app.windows[w]
}
