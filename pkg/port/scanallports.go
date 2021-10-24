package port

import (
	"fmt"
	"strconv"
	"time"
)

func ScanAllPorts(protocol string, ip string, showClosed bool, dialTime int) {
	start := time.Now()
	scanned := 0
	opened := 0

	fmt.Printf("\nStarting port scanning... (%v)\n", ip)
	PrintAddresses(ip)
	fmt.Println("PORT \t\tSTATE \t\tSERVICE")

	for port := 1; port < 65535; port++ {
		scanned++
		if ScanPort(protocol, ip, port, dialTime) {
			opened++
			if len(strconv.Itoa(port)) < 3 {
				fmt.Printf("%v/%v \t\topen \t\t%v\n", port, protocol, PortServiceName(port))
			} else {
				fmt.Printf("%v/%v \topen \t\t%v\n", port, protocol, PortServiceName(port))
			}
		} else {
			if showClosed {
				if len(strconv.Itoa(port)) < 3 {
					fmt.Printf("%v/%v \t\tclosed \t\t%v\n", port, protocol, PortServiceName(port))
				} else {
					fmt.Printf("%v/%v \tclosed \t\t%v\n", port, protocol, PortServiceName(port))
				}
			}
		}
	}

	elapsed := time.Since(start)
	if opened < 1 && !showClosed {
		fmt.Println("- \t\t- \t\t-")
	}
	fmt.Printf("Done. Scanned in %v. \n", elapsed)
	fmt.Printf("Scanned ports: %v\n", scanned)
	fmt.Printf("Open ports: %v\n", opened)
}
