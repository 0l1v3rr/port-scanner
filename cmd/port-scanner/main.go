package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	cl "github.com/0l1v3rr/port-scanner/internal/cli"
	help "github.com/0l1v3rr/port-scanner/internal/help"
	ips "github.com/0l1v3rr/port-scanner/pkg/ip"
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
		printScanner()
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
		} else if strings.HasPrefix(input, "set showclosed") {
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
		} else if strings.HasPrefix(input, "set ports") {
			if len(s) >= 3 {
				if strings.ToLower(s[2]) == "mostknown" {
					port = "mostknown"
				} else if strings.ToLower(s[2]) == "all" {
					port = "all"
				} else {
					fmt.Println("Please provide valid arguments!")
				}
			} else {
				fmt.Println("Please provide valid arguments!")
			}
		} else if strings.HasPrefix(input, "show details") {
			printDetails()
		} else if strings.HasPrefix(input, "run") {
			if strings.HasPrefix(input, "run specific-ports") {
				if len(s) >= 3 {
					ports := strings.Split(s[2], ";")
					var portsConverted []int
					for _, s := range ports {
						converted, cerr := strconv.Atoi(s)
						if cerr == nil {
							portsConverted = append(portsConverted, converted)
						}
					}
					if len(portsConverted) < 1 {
						fmt.Println("Please provide valid arguments!")
					} else {
						pt.ScanSpecificPort(portsConverted, protocol, target, showClosed, dialTime)
					}
				} else {
					fmt.Println("Please provide valid arguments!")
				}
			} else {
				if port == "mostknown" {
					pt.ScanMostKnownPorts(protocol, target, showClosed, dialTime)
				} else if port == "all" {
					pt.ScanAllPorts(protocol, target, showClosed, dialTime)
				}
			}
			fmt.Println()
		} else if strings.HasPrefix(input, "clear") {
			fmt.Print("\033[H\033[2J")
		} else if strings.HasPrefix(input, "show interfaces") {
			fmt.Println()
			ips.Ips()
			fmt.Println()
		} else if strings.HasPrefix(input, "ping") {
			if len(s) >= 2 {
				cl.Ping(s[1])
			} else {
				fmt.Println("Please provide valid arguments!")
			}
		} else if strings.HasPrefix(input, "motd") {
			help.PrintLogo()
			help.PrintMotd()
		} else if strings.HasPrefix(input, "help") {
			help.Help()
		} else if strings.HasPrefix(input, "reset") {
			reset()
		} else if strings.HasPrefix(input, "update") {
			if cl.IsUpdateNeeded() {
				cl.Update()
				time.Sleep(time.Second)
				return
			} else {
				fmt.Println("The port-scanner is up to date.")
			}
			fmt.Println()
		} else if input == "" || input == " " {
			continue
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
	fmt.Println()
}

func reset() {
	target = GetIP().String()
	protocol = "tcp"
	dialTime = 60
	port = "mostknown"
	showClosed = true
}

func printScanner() {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	fmt.Print(string(colorReset), "scanner(")
	fmt.Print(string(colorRed), target)
	fmt.Print(string(colorReset), ") > ")
}
