package handler

import (
	"fmt"

	"github.com/fatih/color"
)

func (h *Handler) Build(values []string) {
	res, err := h.services.App.Build(values)
	if err != nil {
		fmt.Println(err.Error())
	}

	color.Green(res)
}
