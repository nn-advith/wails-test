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

	// appMenu := menu.NewMenu()

	// fileMenu := appMenu.AddSubmenu("File")
	// fileMenu.AddText("Open", nil, func(_ *menu.CallbackData) {
	// 	println("Open clicked!")
	// })
	// fileMenu.AddSeparator()
	// fileMenu.AddText("Exit", nil, func(_ *menu.CallbackData) {
	// 	fmt.Println("Exit clicked!")
	// 	app.QuitApp()
	// })

	// helpMenu := appMenu.AddSubmenu("Help")
	// helpMenu.AddText("About", nil, func(_ *menu.CallbackData) {
	// 	app.About()
	// 	// wails.InfoDialog("About", "Base64 Encoder/Decoder v1.0")
	// })

	// Create application with options
	err := wails.Run(&options.App{
		Title:              "Base64 Encoder/Decoder",
		Frameless:          true,
		Width:              1440,
		Height:             820,
		Logger:             clogger.InitLogger(),
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		CSSDragProperty:    "widows",
		CSSDragValue:       "1",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 30, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			codec,
		},
		// Menu: appMenu,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
