package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	cl "github.com/0l1v3rr/port-scanner/internal/cli"
	help "github.com/0l1v3rr/port-scanner/internal/help"
	pt "github.com/0l1v3rr/port-scanner/pkg/port"
)

var (
	target     string = GetIP().String()
	protocol   string = "tcp"
	dialTime   int    = 60
	port       string = "mostknown"
	showClosed bool   = true
)

func main() {
	cl.Clear()
	reader := bufio.NewReader(os.Stdin)

	help.PrintLogo()
	help.PrintMotd()

	for {
		fmt.Printf("scanner(%v) > ", target)
		input, _ := reader.ReadString('\n')
		input = TrimSuffix(input, "\n")

		if input == "exit" {
			break
		}

		s := strings.Split(input, " ")
		if strings.HasPrefix(input, "set target") {
			if len(s) >= 3 {
				target = s[2]
			} else {
				fmt.Println("Please provide valid arguments!")
			}
		} else if strings.HasPrefix(input, "set protocol") {
			if len(s) >= 3 {
				if s[2] != "tcp" && s[2] != "udp" {
					fmt.Println("Please provide valid arguments!")
				} else {
					protocol = s[2]
				}
			} else {
				fmt.Println("Please provide valid arguments!")
			}
		} else if strings.HasPrefix(input, "set dialtime") {
			if len(s) >= 3 {
				c, e := strconv.Atoi(s[2])
				if e != nil {
					fmt.Println("Please provide valid arguments!")
				} else {
					dialTime = c
				}
			} else {
				fmt.Println("Please provide valid arguments!")
			}
		} else if strings.HasPrefix(input, "set showClosed") || strings.HasPrefix(input, "set showclosed") {
			if len(s) >= 3 {
				if s[2] == "true" {
					showClosed = true
				} else if s[2] == "false" {
					showClosed = false
				} else {
					fmt.Println("Please provide valid arguments!")
				}
			} else {
				fmt.Println("Please provide valid arguments!")
			}
		} else if strings.HasPrefix(input, "show details") {
			printDetails()
		} else if input == "run" {
			if port == "mostknown" {
				pt.ScanMostKnownPorts(protocol, target, showClosed, dialTime)
			} else if port == "all" {
				pt.ScanAllPorts(protocol, target, showClosed, dialTime)
			}
			fmt.Println()
		} else {
			invalidCmd()
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

func TrimSuffix(s, suffix string) string {
	s = s[:len(s)-len(suffix)]
	return s
}

func printDetails() {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	fmt.Println(string(colorReset), "")
	fmt.Print(string(colorReset), " • Target:          ")
	fmt.Println(string(colorRed), target)
	fmt.Print(string(colorReset), " • Protocol:        ")
	fmt.Println(string(colorRed), protocol)
	fmt.Print(string(colorReset), " • Dial timeout:    ")
	fmt.Println(string(colorRed), dialTime)
	fmt.Print(string(colorReset), " • Port(s) to scan: ")
	fmt.Println(string(colorRed), port)
	fmt.Print(string(colorReset), " • Show closed:     ")
	fmt.Println(string(colorRed), showClosed)
	fmt.Println(string(colorReset), "")
}

func invalidCmd() {
	colorYellow := "\033[33m"
	colorReset := "\033[0m"
	fmt.Print("Unknown command. Type ")
	fmt.Print(string(colorYellow), "help")
	fmt.Println(string(colorReset), "for help.")
}

/*func scan(prot string, ipod string, po int) {
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
	pt.PrintAddresses(prot, ipod, port, dialTime)

	elapsed := time.Since(start)
	fmt.Printf("Done. Scanned in %v. \n", elapsed)
}*/
