package ui

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v4"
	"os"
	"runtime"
	"soundTool/env"
	"soundTool/platform"
	"soundTool/renderer"
	"soundTool/ui/forms"
	"soundTool/wrapper"
)

const windowWidth = 1600
const windowHeight = 900

func RunUI() {
	var err error
	env.NewPlatform, err = platform.NewPlatform(windowWidth, windowHeight, "ImguiWidgetTester")
	//var atlas imgui.FontAtlas

	context := imgui.CreateContext(nil)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	keep := true
	version := "#version 430"
	osName := runtime.GOOS
	switch osName {
	//Macos
	case "darwin":
		version = "#version 150"
	default:
	}
	imguiRenderer, err := wrapper.NewOpenGL3(version)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer env.NewPlatform.Dispose()
	defer context.Destroy()
	defer imguiRenderer.Dispose()
	env.ImguiWrapping = wrapper.NewImgui(env.NewPlatform, imguiRenderer, context)
	imgui.StyleColorsClassic()
	env.Init()
	mainRunner := func() {
		/*imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding,6)
		imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding,12)
		imgui.PushStyleVarFloat(imgui.StyleVarPopupRounding,6)
		imgui.PushStyleVarFloat(imgui.StyleVarGrabRounding,12)
		defer imgui.PopStyleVar()
		defer imgui.PopStyleVar()
		defer imgui.PopStyleVar()
		defer imgui.PopStyleVar()*/
		if keep {
			imgui.ShowDemoWindow(&keep)
		}
		//Add your imgui forms call functions here
		forms.MenuBar(env.NewPlatform.DisplaySize())
		forms.RenderMenu()
	}
	env.ImguiWrapping.Run(mainRunner, renderer.Ticker)
	env.DeInit()
}
