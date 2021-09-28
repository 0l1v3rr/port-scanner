package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	hm "github.com/0l1v3rr/port-scanner/internal/help"
	validip "github.com/0l1v3rr/port-scanner/pkg/ip"
	pt "github.com/0l1v3rr/port-scanner/pkg/port"
)

var (
	defaultProtocol string = "tcp"
	defaultIp       string = GetIP().String()
	defaultDialTime int    = 60
	protocol        string
	dialTime        int
	ip              string
	domain          string
	port            int
	showClosed      bool
	allPorts        bool
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
	flag.IntVar(&dialTime, "dialtime", defaultDialTime, "The dial timeout you want to use.")
	flag.StringVar(&domain, "domain", "", "The domain name you want to scan.")
	flag.BoolVar(&showClosed, "closed", false, "With this flag, the app won't show the closed ports.")
	flag.BoolVar(&allPorts, "all", false, "With this flag, the app will scan all the ports from 1 to 65535.")
	flag.Parse()

	if protocol != "udp" && protocol != "tcp" {
		fmt.Println("Error: Invalid protocol.")
		return
	}

	if !validip.IsValidIp(ip) {
		fmt.Println("Error: Invalid IP.")
		return
	}

	if allPorts {
		if domain != "" {
			pt.ScanAllPorts(protocol, domain, !showClosed, dialTime)
		} else {
			pt.ScanAllPorts(protocol, ip, !showClosed, dialTime)
		}
	} else {
		if port == -123 {
			if domain != "" {
				pt.ScanMostKnownPorts(protocol, domain, !showClosed, dialTime)
			} else {
				pt.ScanMostKnownPorts(protocol, ip, !showClosed, dialTime)
			}
		} else {
			if port < 1 {
				fmt.Println("Error: Invalid port.")
				return
			}

			if domain != "" {
				scan(protocol, domain, port)
			} else {
				scan(protocol, ip, port)
			}

		}
	}

}

func GetIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func scan(prot string, ipod string, po int) {
	start := time.Now()

	fmt.Printf("\nStarting port scanning... (%v)\n", ipod)
	fmt.Println("PORT \t\tSTATE \t\tSERVICE")
	open := pt.ScanPort(prot, ipod, port, dialTime)

	if len(strconv.Itoa(port)) < 3 {
		if open {
			fmt.Printf("%v/%v \t\topen \t\t%v\n", po, prot, pt.PortServiceName(po))
		} else {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", po, prot, pt.PortServiceName(po))
		}
	} else {
		if open {
			fmt.Printf("%v/%v \topen \t\t%v\n", po, prot, pt.PortServiceName(po))
		} else {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", po, prot, pt.PortServiceName(po))
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done. Scanned in %v. \n", elapsed)
}
