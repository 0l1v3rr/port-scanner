package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	protocol   string
	ip         string
	port       int
	showClosed bool
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
	flag.IntVar(&port, "port", -123, "The only port you want to scan.")
	flag.BoolVar(&showClosed, "closed", false, "With this flag, the app won't show the closed ports.")
	flag.Parse()

	if protocol != "udp" && protocol != "tcp" {
		fmt.Println("Error: Invalid protocol.")
		return
	}

	if !validip.IsValidIp(ip) {
		fmt.Println("Error: Invalid IP.")
		return
	}

	if port == -123 {
		pt.ScanMostKnownPorts(protocol, ip, !showClosed)
	} else {
		if port < 1 {
			fmt.Println("Error: Invalid port.")
			return
		}

		start := time.Now()

		fmt.Printf("\nStarting port scanning... (%v)\n", ip)
		fmt.Println("PORT \t\tSTATE \t\tSERVICE")
		open := pt.ScanPort(protocol, ip, port)

		if len(strconv.Itoa(port)) < 3 {
			if open {
				fmt.Printf("%v/%v \t\topen \t\t%v\n", port, protocol, pt.PortServiceName(port))
			} else {
				fmt.Printf("%v/%v \t\tclosed \t\t%v\n", port, protocol, pt.PortServiceName(port))
			}
		} else {
			if open {
				fmt.Printf("%v/%v \topen \t\t%v\n", port, protocol, pt.PortServiceName(port))
			} else {
				fmt.Printf("%v/%v \tclosed \t\t%v\n", port, protocol, pt.PortServiceName(port))
			}
		}

		elapsed := time.Since(start)
		fmt.Printf("Done. Scanned in %v. \n", elapsed)
	}

}
