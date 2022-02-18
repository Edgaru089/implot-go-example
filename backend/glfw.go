package backend

import (
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
)

const (
	mouseButtonPrimary = iota
	mouseButtonSecondary
	mouseButtonTertiary
	mouseButtonCount
)

var (
	win *glfw.Window
	io  imgui.IO

	lastframe        time.Time
	mouseJustPressed [mouseButtonCount]bool
)

func Init(window *glfw.Window) {
	win = window
	io = imgui.CurrentIO()

	setKeymap()
	lastframe = time.Now()

	renderInit()
}

// NewFrame marks the begin of a render pass.
func NewFrame() {
	dsx, dsy := win.GetSize()
	io.SetDisplaySize(imgui.Vec2{X: float32(dsx), Y: float32(dsy)})

	now := time.Now()
	deltaTime := float32(time.Since(lastframe).Seconds())
	if deltaTime <= 0.0 {
		deltaTime = 1e-6
	}
	io.SetDeltaTime(deltaTime)
	lastframe = now

	if win.GetAttrib(glfw.Focused) != 0 {
		x, y := win.GetCursorPos()
		io.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
	}

	for i := 0; i < mouseButtonCount; i++ {
		down := mouseJustPressed[i] || (win.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		io.SetMouseButtonDown(i, down)
		mouseJustPressed[i] = false
	}
	imgui.NewFrame()
}

func setKeymap() {
	io.KeyMap(imgui.KeyTab, int(glfw.KeyTab))
	io.KeyMap(imgui.KeyLeftArrow, int(glfw.KeyLeft))
	io.KeyMap(imgui.KeyRightArrow, int(glfw.KeyRight))
	io.KeyMap(imgui.KeyUpArrow, int(glfw.KeyUp))
	io.KeyMap(imgui.KeyDownArrow, int(glfw.KeyDown))
	io.KeyMap(imgui.KeyPageUp, int(glfw.KeyPageUp))
	io.KeyMap(imgui.KeyPageDown, int(glfw.KeyPageDown))
	io.KeyMap(imgui.KeyHome, int(glfw.KeyHome))
	io.KeyMap(imgui.KeyEnd, int(glfw.KeyEnd))
	io.KeyMap(imgui.KeyInsert, int(glfw.KeyInsert))
	io.KeyMap(imgui.KeyDelete, int(glfw.KeyDelete))
	io.KeyMap(imgui.KeyBackspace, int(glfw.KeyBackspace))
	io.KeyMap(imgui.KeySpace, int(glfw.KeySpace))
	io.KeyMap(imgui.KeyEnter, int(glfw.KeyEnter))
	io.KeyMap(imgui.KeyEscape, int(glfw.KeyEscape))
	io.KeyMap(imgui.KeyA, int(glfw.KeyA))
	io.KeyMap(imgui.KeyC, int(glfw.KeyC))
	io.KeyMap(imgui.KeyV, int(glfw.KeyV))
	io.KeyMap(imgui.KeyX, int(glfw.KeyX))
	io.KeyMap(imgui.KeyY, int(glfw.KeyY))
	io.KeyMap(imgui.KeyZ, int(glfw.KeyZ))
}

var glfwButtonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: mouseButtonPrimary,
	glfw.MouseButton2: mouseButtonSecondary,
	glfw.MouseButton3: mouseButtonTertiary,
}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	mouseButtonPrimary:   glfw.MouseButton1,
	mouseButtonSecondary: glfw.MouseButton2,
	mouseButtonTertiary:  glfw.MouseButton3,
}

// MouseButtonCallback is the callback called when the mouse button changes.
func MouseButtonCallback(button glfw.MouseButton, action glfw.Action) {
	if index, known := glfwButtonIndexByID[button]; known && (action == glfw.Press) {
		mouseJustPressed[index] = true
	}
}

// MouseScrollCallback is called when scroll status changes.
func MouseScrollCallback(x, y float64) {
	io.AddMouseWheelDelta(float32(x), float32(y))
}

// KeyCallback is called when a key is pressed or released.
func KeyCallback(key glfw.Key, action glfw.Action) {
	if action == glfw.Press {
		io.KeyPress(int(key))
	}
	if action == glfw.Release {
		io.KeyRelease(int(key))
	}
	io.KeyCtrl(int(glfw.KeyLeftControl), int(glfw.KeyRightControl))
	io.KeyShift(int(glfw.KeyLeftShift), int(glfw.KeyRightShift))
	io.KeyAlt(int(glfw.KeyLeftAlt), int(glfw.KeyRightAlt))
	io.KeySuper(int(glfw.KeyLeftSuper), int(glfw.KeyRightSuper))
}

// InputCallback is called when a char is inputed (CharChange)
func InputCallback(input rune) {
	io.AddInputCharacters(string(input))
}
