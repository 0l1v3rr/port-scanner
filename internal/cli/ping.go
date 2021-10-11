package cli

import (
	"fmt"
	"os/exec"
)

func Ping(ip string) {
	out, err := exec.Command("ping", ip, "-i 3", "-w 10").Output()

	if err != nil {
		fmt.Println("Unknown host.")
		return
	}

	if out != nil {
		fmt.Println(string(out))
	}
}
