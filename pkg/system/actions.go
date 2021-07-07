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
	fmt.Println("always be careful don't leave a trace...")
	os.Exit(0)
}
