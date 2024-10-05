package main

import (
	"embed"
	"openHueApp/backend/dao"
	"openHueApp/backend/hue"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {

	hue := hue.NewHue()
	app := dao.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Hue",
		Width:            1024,
		Height:           768,
		DisableResize:    false,
		BackgroundColour: &options.RGBA{R: 99, G: 102, B: 106, A: 200},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			hue,
			app,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "My Application",
				Message: "© 2024 jakob ferber",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
