package main

import (
	"runtime"

	"github.com/Edgaru089/implot-go"
	"github.com/Edgaru089/implot-go-example/backend"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
)

const (
	width  = 1024
	height = 768
)

var (
	showDemoWindow = true
	showImPlotDemo = true
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, 1)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)
	win, err := glfw.CreateWindow(width, height, "Gl01", nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()
	glfw.SwapInterval(1)

	imgui.CreateContext(nil)
	implot.CreateContext()

	backend.Init(win)
	win.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		backend.MouseButtonCallback(button, action)
	})
	win.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		backend.KeyCallback(key, action)
	})
	win.SetCharCallback(func(w *glfw.Window, char rune) {
		backend.InputCallback(char)
	})
	win.SetScrollCallback(func(w *glfw.Window, xoff, yoff float64) {
		backend.MouseScrollCallback(xoff, yoff)
	})

	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Viewport(0, 0, width, height)
	win.SwapBuffers()

	for !win.ShouldClose() {
		backend.NewFrame()

		imgui.ShowDemoWindow(&showDemoWindow)
		implot.ShowDemoWindow(&showImPlotDemo)

		example()

		gl.Clear(gl.COLOR_BUFFER_BIT)
		backend.Render(win)
		win.SwapBuffers()
		glfw.PollEvents()
	}
}
