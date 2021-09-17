package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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

	start := time.Now()

	fmt.Printf("\nStarting port scanning... (%v)\n", ip)
	fmt.Println("PROTOCOL \tPORT \t\tSTATE \t\tSERVICE")
	if open {
		fmt.Printf("%v \t\t%v \t\topen \t\t%v \n", protocol, port, pt.PortServiceName(port))
	} else {
		fmt.Printf("%v \t\t%v \t\tclosed \t\t%v \n", protocol, port, pt.PortServiceName(port))
	}

	elapsed := time.Since(start)
	fmt.Printf("Done. Scanned in %v. \n", elapsed*1000)

}
