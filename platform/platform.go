package platform

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

// Platform is a holder for the glfw window
type Platform struct {
	Window *glfw.Window
}

// NewPlatform attempts to initialize a GLFW context.
func NewPlatform(windowWidth int, windowHeight int, name string) (*Platform, error) {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize glfw: %v", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, name, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("failed to create window: %v", err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	platform := &Platform{Window: window}

	return platform, nil
}

// Dispose cleans up the resources.
func (platform *Platform) Dispose() {
	platform.Window.Destroy()
	glfw.Terminate()
}

// ShouldStop returns true if the window is to be closed.
func (platform *Platform) ShouldStop() bool {
	return platform.Window.ShouldClose()
}

// ProcessEvents handles all pending window events.
func (platform *Platform) ProcessEvents() {
	glfw.PollEvents()
}

// IsFocused returns the focus status of the window
func (platform *Platform) IsFocused() bool {
	return platform.Window.GetAttrib(glfw.Focused) != 0
}

// DisplaySize returns the dimension of the display.
func (platform *Platform) DisplaySize() [2]float32 {
	w, h := platform.Window.GetSize()
	return [2]float32{float32(w), float32(h)}
}

// GetCursorPos returs the cursor x and y position
func (platform *Platform) GetCursorPos() (float64, float64) {
	return platform.Window.GetCursorPos()
}

// GetMousePress returns true if mouse buttons are currently pressed
func (platform *Platform) GetMousePress(mouseButton glfw.MouseButton) bool {
	return platform.Window.GetMouseButton(mouseButton) == glfw.Press
}

// GetMousePresses123 returns press status of mouse buttons 1,2 and 3
func (platform *Platform) GetMousePresses123() [3]bool {
	return [3]bool{platform.GetMousePress(glfw.MouseButton1),
		platform.GetMousePress(glfw.MouseButton2),
		platform.GetMousePress(glfw.MouseButton3)}
}

// FramebufferSize returns the dimension of the framebuffer.
func (platform *Platform) FramebufferSize() [2]float32 {
	w, h := platform.Window.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}

// PostRender performs a buffer swap.
func (platform *Platform) PostRender() {
	platform.Window.SwapBuffers()
}

// SetMouseButtonCallback sets a glfw compatible mouse callback function
func (platform *Platform) SetMouseButtonCallback(callback glfw.MouseButtonCallback) {
	platform.Window.SetMouseButtonCallback(callback)
}

// SetScrollCallback sets a glfw compatible scroll callback function
func (platform *Platform) SetScrollCallback(callback glfw.ScrollCallback) {
	platform.Window.SetScrollCallback(callback)
}

// SetKeyCallback sets a glfw compatible key callback function
func (platform *Platform) SetKeyCallback(callback glfw.KeyCallback) {
	platform.Window.SetKeyCallback(callback)
}

// SetCharCallback sets a glfw compatible char callback function
func (platform *Platform) SetCharCallback(callback glfw.CharCallback) {
	platform.Window.SetCharCallback(callback)
}

// ClipboardText returns the current clipboard text, if available.
func (platform *Platform) ClipboardText() (string, error) {
	return platform.Window.GetClipboardString()
}

// SetClipboardText sets the text as the current clipboard text.
func (platform *Platform) SetClipboardText(text string) {
	platform.Window.SetClipboardString(text)
}
