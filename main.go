package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "OnlyGood",
		Width:  1280,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
            TitleBar: &mac.TitleBar{
                TitlebarAppearsTransparent: false,
                HideTitle:                  true,
                HideTitleBar:               false,
                FullSizeContent:            false,
                UseToolbar:                 true,
                HideToolbarSeparator:       false,
            },
            Appearance:           mac.DefaultAppearance,
            WebviewIsTransparent: true,
            WindowIsTranslucent:  true,
            About: &mac.AboutInfo{
                Title:   "OnlyGood",
                Message: "© Mitchell Technologies Limited",
                // Icon:    icon,
            },
        },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
