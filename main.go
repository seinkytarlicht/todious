package main

import (
	"context"
	"embed"

	"github.com/seinkytarlicht/todious/internal/database"
	"github.com/seinkytarlicht/todious/internal/repository"
	"github.com/seinkytarlicht/todious/internal/service"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/.output/public
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "Todious",
		Description: "A demo Wails 3 and Nuxt 4",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:     "Todious",
		MinWidth:  1000,
		MinHeight: 618,
		Width:     1000,
		Height:    618,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
		Frameless:        true,
	})

	db := database.New()
	ctx := context.Background()
	todoRepo := repository.NewTodoRepository(db, ctx)
	todoService := service.NewTodoService(todoRepo, app)

	app.RegisterService(application.NewService(todoService))

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
