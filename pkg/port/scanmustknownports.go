package port

import (
	"fmt"
	"time"
)

func ScanMostKnownPorts(protocol string, ip string, showClosed bool) {

	start := time.Now()
	scanned := 0
	opened := 0

	fmt.Printf("\nStarting port scanning... (%v)\n", ip)
	if showClosed {
		fmt.Println("PORT \t\tSTATE \t\tSERVICE")
	}

	actualport := 21
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 22
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 23
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 25
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 53
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 69
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 79
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 80
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \t\topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \t\tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 110
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 111
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 139
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 143
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 443
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 445
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 513
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 514
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 995
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 1080
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 2049
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 2082
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 2121
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 2745
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 3306
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 3127
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 4444
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 5432
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 5554
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 5900
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 6000
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 8866
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 9898
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	actualport = 9988
	scanned++
	if ScanPort(protocol, ip, actualport) {
		opened++
		fmt.Printf("%v/%v \topen \t\t%v\n", actualport, protocol, PortServiceName(actualport))
	} else {
		if showClosed {
			fmt.Printf("%v/%v \tclosed \t\t%v\n", actualport, protocol, PortServiceName(actualport))
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done. Scanned in %v. \n", elapsed)
	if opened < 1 && !showClosed {
		fmt.Println("There's no opened port.")
	}
	if showClosed {
		fmt.Printf("Scanned ports: %v\n", scanned)
		fmt.Printf("Open ports: %v\n", opened)
	}

}
