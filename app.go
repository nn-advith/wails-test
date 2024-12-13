package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

func (a *App) QuitApp() {
	if a.ctx != nil {
		runtime.Quit(a.ctx)
	}
}

func (a *App) About() {
	if a.ctx != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    "info",
			Title:   "Information",
			Message: "This is an informational message.",
			Buttons: []string{"OK"},
		})
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	// temp := clogger.InitLogger()
	// temp.Log("INFO", "This is a sample test log")
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
