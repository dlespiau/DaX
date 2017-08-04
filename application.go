package dax

import (
	"log"
	"runtime"
	"sync"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize gl:", err)
	}
}

// Application object is the top level object from which everything else in DaX
// is derived.
type Application struct {
	Name string

	windows map[*glfw.Window]*Window
}

var appInstance *Application
var appOnce sync.Once

// NewApplication creates a new Application. This is a singleton.
func NewApplication(name string) *Application {
	appOnce.Do(func() {
		app := new(Application)
		app.Name = name
		app.windows = make(map[*glfw.Window]*Window)
		appInstance = app
	})
	return appInstance
}

func (app *Application) addWindow(window *Window) {
	// TODO: support multiple windows
	app.windows[window.glfwWindow] = window
}

// Run enters the application main loop.
func (app *Application) Run() {
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
func (app *Application) CreateWindow(name string, width, height int) *Window {
	window := newWindow(app, name, width, height)
	app.addWindow(window)

	return window
}

func getWindow(w *glfw.Window) *Window {
	return appInstance.windows[w]
}
