package handler

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/fr13n8/Bacterio/internal/ui"
	"github.com/fr13n8/Bacterio/pkg/network"
	"github.com/fr13n8/Bacterio/pkg/utils"
)

func (h *Handler) Listen(v []string) {
	params := map[string]string{
		"host": "",
		"prot": "",
	}
	vals, err := utils.Validate(v, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	h.services.Server.CreateServer(vals["host"], vals["port"]).HandleConnects()

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
