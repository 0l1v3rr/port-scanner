package cli

import (
	"fmt"
	"net"
)

func PrintMacAddr() {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"

	ifas, err := net.Interfaces()
	if err != nil {
		fmt.Println("An unknown error occurred.")
		return
	}

	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			fmt.Print(string(colorCyan), " • ")
			fmt.Print(string(colorCyan), ifa.Name)
			fmt.Print(": ")
			fmt.Println(string(colorReset), a)
		}
	}
}
