package handler

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/internal/ui"
	"github.com/fr13n8/Bacterio/internal/utils"
	"github.com/fr13n8/Bacterio/pkg/network"
)

func (h *Handler) Listen(v []string) {
	if !utils.Contains(v, "host=") {
		color.Yellow(" [!] You should set a host!")
		return
	}
	if !utils.Contains(v, "port=") {
		color.Yellow(" [!] You should set a port!")
		return
	}

	address := utils.SplitAfterIndex(utils.Find(v, "host="), '=')
	port := utils.SplitAfterIndex(utils.Find(v, "port="), '=')

	h.services.Server.CreateServer(address, port).HandleConnects()

	p := prompt.New(
		h.ServerExecutor,
		ui.ServerCompleter,
		prompt.OptionPrefix(fmt.Sprintf("server [%s] > ", network.GetLocalIP().String())),
		prompt.OptionPrefixTextColor(prompt.White),
	)
	p.Run()
}

func (h *Handler) ShowConnects() {
	h.services.Server.ShowConnects()
}

func (h *Handler) Connect(v []string) {

	target := h.services.Server.SetTarget(v)
	if target == nil {
		return
	}

	p := prompt.New(
		h.TargetExecutor,
		ui.TargetCompleter,
		prompt.OptionPrefix(fmt.Sprintf("%s@%s > ", target.Hostname, target.Username)),
		prompt.OptionPrefixTextColor(prompt.Yellow),
	)
	p.Run()
}
