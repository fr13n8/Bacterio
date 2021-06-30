package ui

import (
	"github.com/fr13n8/Bacterio/pkg/system"
)

func StartUi(ver string) {
	system.ClearScreen()
	ShowHeader()
	ShowInfo(ver)
}
