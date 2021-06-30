package handler

import "github.com/fr13n8/Bacterio/pkg/system"

func (h *Handler) Exit() {
	system.ExitApp()
}
