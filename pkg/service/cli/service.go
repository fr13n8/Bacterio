package service

import "github.com/fr13n8/Bacterio/internal/models"

type App interface {
	ClearScreen()
	ExitApp()
	Build([]string) (string, error)
}

type Server interface {
	CreateServer(address, port string) *serverService
	ShowConnects()
	HandleConnects()
	SetTarget(v []string) *models.Connect
	GetInformation() ([]byte, error)
	RunCommand(cmd string) ([]byte, error)
	RunStiller() (*models.Message, error)
	TakeScreenshot() ([]byte, error)
}

type Service struct {
	App
	Server
}

func NewService() *Service {
	return &Service{
		App:    NewAppService(),
		Server: NewServerService(),
	}
}
