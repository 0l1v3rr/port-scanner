package ip

import (
	"fmt"
	"net"
)

func Ips() {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("An error occurred.")
		return
	}

	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				fmt.Print(string(colorCyan), " • ")
				fmt.Print(string(colorCyan), i.Name)
				fmt.Print(": ")
				fmt.Println(string(colorReset), v)

			case *net.IPNet:
				fmt.Print(string(colorCyan), " • ")
				fmt.Print(string(colorCyan), i.Name)
				fmt.Print(": ")
				fmt.Println(string(colorReset), v)
			}

		}
	}

}
