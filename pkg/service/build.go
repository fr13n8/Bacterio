package service

type BuildService struct {
}

func NewBuildService() *BuildService {
	return &BuildService{}
}

func (build *BuildService) Build([]string) (string, error) {
	return "xuy", nil
}
