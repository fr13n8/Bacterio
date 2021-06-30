package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/fr13n8/Bacterio/internal/ui"
	"github.com/fr13n8/Bacterio/pkg/handler"
	"github.com/fr13n8/Bacterio/pkg/network"
	"github.com/fr13n8/Bacterio/pkg/service"
)

func main() {
	services := service.NewService()
	handler := handler.NewHandler(services)

	ui.StartUi("dev")
	p := prompt.New(
		handler.Executor,
		ui.HostCompleter,
		prompt.OptionPrefix(network.GetLocalIP().String()+" >"),
		prompt.OptionPrefixTextColor(prompt.White),
	)
	p.Run()
}
