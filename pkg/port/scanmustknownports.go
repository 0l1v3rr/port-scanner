package port

import (
	"fmt"
	"strconv"
	"time"
)

var ports [100]int = [100]int{21, 22, 23, 25, 53, 69, 79, 80, 110, 111, 139, 143, 443, 445, 513, 514, 995, 1080, 2049, 2082, 2121, 2745, 3306, 3127, 4444, 5432, 5554, 5900, 6000, 8866, 9898, 9988}

func ScanMostKnownPorts(protocol string, ip string, showClosed bool) {

	start := time.Now()
	scanned := 0
	opened := 0

	fmt.Printf("\nStarting port scanning... (%v)\n", ip)
	fmt.Println("PORT \t\tSTATE \t\tSERVICE")

	for _, port := range ports {
		if port != 0 {
			scanned++
			if ScanPort(protocol, ip, port) {
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
	}

	elapsed := time.Since(start)
	if opened < 1 && !showClosed {
		fmt.Println("- \t\t- \t\t-")
	}
	fmt.Printf("Done. Scanned in %v. \n", elapsed)
	fmt.Printf("Scanned ports: %v\n", scanned)
	fmt.Printf("Open ports: %v\n", opened)

}
