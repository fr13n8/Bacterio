package handler

import "github.com/fatih/color"

func (h *Handler) Build(values []string) {
	res, err := h.services.Build.Build(values)
	if err != nil {
		color.Red(err.Error())
	}

	color.Green(res)
}
