package service

type Build interface {
	Build([]string) (string, error)
}

type Cmd interface {
}

type Service struct {
	Build
	Cmd
}

func NewService() *Service {
	return &Service{
		Build: NewBuildService(),
		Cmd:   NewCmdService(),
	}
}
