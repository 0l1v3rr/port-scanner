package cli

import (
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	cmd2 := exec.Command("cls")
	cmd2.Stdout = os.Stdout
	cmd2.Run()
}
