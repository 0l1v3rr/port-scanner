package main

import (
	"flag"
	"fmt"
	"os"

	hm "github.com/0l1v3rr/port-scanner/internal/help"
	validip "github.com/0l1v3rr/port-scanner/pkg/ip"
	pt "github.com/0l1v3rr/port-scanner/pkg/port"
)

const (
	defaultProtocol string = "tcp"
	defaultIp       string = "localhost"
)

var (
	protocol string
	ip       string
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			hm.PrintHelpMenu()
			return
		}
	}

	flag.StringVar(&protocol, "protocol", defaultProtocol, "The protocol")
	flag.StringVar(&ip, "ip", defaultIp, "The IP Address you want to scan.")
	flag.Parse()

	if protocol != "udp" && protocol != "tcp" {
		fmt.Println("Error: Invalid protocol.")
		return
	}

	if !validip.IsValidIp(ip) {
		fmt.Println("Error: Invalid IP.")
		return
	}

	port := 80
	open := pt.ScanPort(protocol, ip, port)

	fmt.Println("PORT \tSTATE \tSERVICE")
	if open {
		fmt.Printf("%v/%v \topen \thttp", port, protocol)
	} else {
		fmt.Printf("%v/%v \tclosed \thttp", port, protocol)
	}

}
