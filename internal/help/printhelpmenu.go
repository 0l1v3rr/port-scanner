package help

import (
	"fmt"

	cl "github.com/0l1v3rr/port-scanner/internal/cli"
)

func PrintMotd() {
	colorYellow := "\033[33;1m"
	colorReset := "\033[0m"
	colorBlue := "\033[34;1m"
	colorCyan := "\033[36;1m"
	colorRed := "\033[31;1m"

	cl.UpdateNeeded()
	fmt.Print(string(colorBlue), " [>] Version: ")
	fmt.Println(string(colorReset), "port-scanner v2.3.0")
	fmt.Print(string(colorCyan), " [>] GitHub:  ")
	fmt.Println(string(colorReset), "https://github.com/0l1v3rr/port-scanner")
	fmt.Print(string(colorReset), " [x] Type ")
	fmt.Print(string(colorYellow), "help")
	fmt.Println(string(colorReset), "for more.")
	fmt.Print(string(colorReset), " [x] Type ")
	fmt.Print(string(colorRed), "exit")
	fmt.Println(string(colorReset), "for exit.")
	fmt.Println()
}

func PrintLogo() {
	colorCyan := "\033[36;1m"
	colorReset := "\033[0m"

	fmt.Println(string(colorCyan), "        ___  ___  ___ _____ ")
	fmt.Println(string(colorCyan), "       | _ \\/ _ \\| _ \\_   _|")
	fmt.Println(string(colorCyan), "       |  _/ (_) |   / | |  ")
	fmt.Println(string(colorCyan), "       |_|  \\___/|_|_\\ |_|  ")
	fmt.Println(string(colorCyan), "  ___  ___   _   _  _ _  _ ___ ___ ")
	fmt.Println(string(colorCyan), " / __|/ __| /_\\ | \\| | \\| | __| _ \\")
	fmt.Println(string(colorCyan), " \\__ \\ (__ / _ \\| .` | .` | _||   /")
	fmt.Println(string(colorCyan), " |___/\\___/_/ \\_\\_|\\_|_|\\_|___|_|_\\")
	fmt.Println(string(colorReset), "")
}

func Help() {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"
	fmt.Println()

	fmt.Println(string(colorReset), "COMMAND                       | DESCRIPTION")
	fmt.Println(string(colorReset), "------------------------------|-----------------------------")
	fmt.Print(string(colorCyan), " • set target <ip>            ")
	fmt.Println(string(colorReset), "| Sets the target IP address.")
	fmt.Print(string(colorCyan), " • set protocol <tcp/udp>     ")
	fmt.Println(string(colorReset), "| Sets the default protocol.")
	fmt.Print(string(colorCyan), " • set dialtime <sec>         ")
	fmt.Println(string(colorReset), "| Sets the dial timeout you want to use.")
	fmt.Print(string(colorCyan), " • set showclosed <true/false>")
	fmt.Println(string(colorReset), "| Toggle the display of open ports on and off.")
	fmt.Print(string(colorCyan), " • set ports <mostknown/all>  ")
	fmt.Println(string(colorReset), "| Sets the ports you want to scan.")
	fmt.Print(string(colorCyan), " • show details               ")
	fmt.Println(string(colorReset), "| Prints the details.")
	fmt.Print(string(colorCyan), " • clear                      ")
	fmt.Println(string(colorReset), "| Clears the terminal.")
	fmt.Print(string(colorCyan), " • motd                       ")
	fmt.Println(string(colorReset), "| Prints the banner.")
	fmt.Print(string(colorCyan), " • do <command>               ")
	fmt.Println(string(colorReset), "| Runs a terminal command.")
	fmt.Print(string(colorCyan), " • show interfaces            ")
	fmt.Println(string(colorReset), "| Shows all the interfaces and their IP.")
	fmt.Print(string(colorCyan), " • show mac                   ")
	fmt.Println(string(colorReset), "| Shows your mac address.")
	fmt.Print(string(colorCyan), " • update                     ")
	fmt.Println(string(colorReset), "| If it's needed, the program will update itself.")
	fmt.Print(string(colorCyan), " • ipinfo <IP>                ")
	fmt.Println(string(colorReset), "| Prints some information about the specified IP address.")
	fmt.Print(string(colorCyan), " • show ip                    ")
	fmt.Println(string(colorReset), "| Shows your IP addresses.")
	fmt.Print(string(colorCyan), " • run                        ")
	fmt.Println(string(colorReset), "| Runs the scan.")
	fmt.Print(string(colorCyan), " • run specific-ports <p1;p2;>")
	fmt.Println(string(colorReset), "| Runs the scan on the specified ports.")
	fmt.Print(string(colorCyan), " • reset                      ")
	fmt.Println(string(colorReset), "| Resets the given values to default.")

	fmt.Println()
}
