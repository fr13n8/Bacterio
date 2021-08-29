package handler

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fr13n8/Bacterio/pkg/utils"
)

func (h *Handler) Build(v []string) {
	params := map[string]string{
		"host":   "",
		"port":   "",
		"os":     "",
		"output": "",
		"attrs":  "",
	}
	vals, err := utils.Validate(v, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res, err := h.services.App.Build(vals)
	if err != nil {
		fmt.Println(err.Error())
	}

	color.Green(res)
}
