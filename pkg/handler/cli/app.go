package handler

func (h *Handler) Exit() {
	h.services.ExitApp()
}
