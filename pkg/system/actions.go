package system

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ExitApp() {
	ClearScreen()
	fmt.Println("Bye, See you later!")
	os.Exit(0)
}
