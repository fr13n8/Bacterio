package handler

import (
	"strings"

	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/pkg/service"
	"github.com/fr13n8/Bacterio/pkg/system"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Executor(input string) {
	values := strings.Fields(input)
	for _, v := range values {
		switch v {
		case "build":
			h.Build(values)
			return
		case "listen":
			h.Listen(values)
			return
		case "exit":
			h.Exit()
		default:
			color.Red(" [!] Invalid parameter!")
			// util.Sleep(3)
			system.ClearScreen()
			return
		}
	}
}
