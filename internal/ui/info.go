package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func ShowInfo(version string) {
	fmt.Println("")
	color.Yellow(" version " + version)
	color.Cyan(" By efr13nd")
	fmt.Println("")
	color.White(" Please use `tab` to autocomplete commands.")
	color.White(" Type `exit` to quit this program.")
	fmt.Println("")
}
