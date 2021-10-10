package cli

import (
	"fmt"
	"os/exec"
	"strings"
)

func Ping(ip string) {
	out, _ := exec.Command("ping", ip, "-i 3", "-w 10").Output()

	if strings.Contains(string(out), "Destination Host Unreachable") || strings.Contains(string(out), "Ping request could not find host") || strings.Contains(string(out), "ping: unknown host") {
		fmt.Println("The host is unreachable.")
	} else {
		fmt.Println("The host is alive!")
	}
}
