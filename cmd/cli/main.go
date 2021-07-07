package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/fr13n8/Bacterio/internal/ui"
	handler "github.com/fr13n8/Bacterio/pkg/handler/cli"
	"github.com/fr13n8/Bacterio/pkg/network"
	service "github.com/fr13n8/Bacterio/pkg/service/cli"
)

func main() {
	services := service.NewService()
	handler := handler.NewHandler(services)

	ui.StartUi("dev")
	p := prompt.New(
		handler.AppExecutor,
		ui.AppCompleter,
		prompt.OptionPrefix(fmt.Sprintf(" %s > ", network.GetLocalIP().String())),
		prompt.OptionPrefixTextColor(prompt.White),
	)
	p.Run()
}
