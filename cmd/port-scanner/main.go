package main

import (
	"flag"
	"fmt"
	"os"

	hm "github.com/0l1v3rr/port-scanner/internal/help"
	pt "github.com/0l1v3rr/port-scanner/pkg/port"
)

const (
	defaultProtocol string = "tcp"
)

var (
	protocol string
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			hm.PrintHelpMenu()
			return
		}
	}

	flag.StringVar(&protocol, "protocol", defaultProtocol, "The protocol.")
	flag.Parse()

	if protocol != "udp" && protocol != "tcp" {
		fmt.Println("Error: Invalid protocol.")
		return
	}

	port := 80
	open := pt.ScanPort(protocol, "localhost", port)

	fmt.Println("PORT \tSTATE \tSERVICE")
	if open {
		fmt.Printf("%v/%v \topen \thttp", port, protocol)
	} else {
		fmt.Printf("%v/%v \tclosed \thttp", port, protocol)
	}

}
