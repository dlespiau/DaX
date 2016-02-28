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

type Application struct {
	name   string
	window *Window
}

var _app *Application

func (app *Application) addWindow(window *Window) {
	// TODO: support multiple windows
	app.window = window
}

func (app *Application) Run() {
	for !app.window.glfwWindow.ShouldClose() {
		app.window.Update()
		app.window.Draw()
		app.window.glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
}

func (app *Application) CreateWindow(name string, width, height int) *Window {
	_app = app

	window := newWindow(app, name, width, height)
	app.addWindow(window)

	return window
}

func getWindow(w *glfw.Window) *Window {
	return _app.window
}
