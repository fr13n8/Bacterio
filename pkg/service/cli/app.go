package service

import "github.com/fr13n8/Bacterio/pkg/system"

type AppService struct {
}

func NewAppService() *AppService {
	return &AppService{}
}

func (c *AppService) ClearScreen() {
	system.ClearScreen()
}

func (c *AppService) ExitApp() {
	system.ExitApp()
}

func (c *AppService) Build(map[string]string) (string, error) {
	return "xuy", nil
}
