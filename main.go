package main

import (
	"b64/clogger"
	"b64/codec"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	codec := codec.NewCodec()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "Base64 Encoder/Decoder",
		// Frameless:          true,
		Width:              1024,
		Height:             768,
		Logger:             clogger.InitLogger(),
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 37, G: 37, B: 37, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			codec,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
