package help

import "fmt"

func PrintMotd() {
	colorYellow := "\033[33m"
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorCyan := "\033[36m"
	colorRed := "\033[31m"

	fmt.Println(string(colorGreen), "• port-scanner v2.0.1")
	fmt.Println(string(colorCyan), "• https://github.com/0l1v3rr/port-scanner")
	fmt.Print(string(colorReset), " • Type ")
	fmt.Print(string(colorYellow), "help")
	fmt.Println(string(colorReset), "for more.")
	fmt.Print(string(colorReset), " • Type ")
	fmt.Print(string(colorRed), "exit")
	fmt.Println(string(colorReset), "for exit.")
	fmt.Println()
}

func PrintLogo() {
	colorCyan := "\033[36m"
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

	fmt.Print(string(colorCyan), "• set target <ip>")
	fmt.Println(string(colorReset), " - Sets the target IP address.")
	fmt.Print(string(colorCyan), "• set protocol <tcp/udp>")
	fmt.Println(string(colorReset), " - Sets the default protocol.")
	fmt.Print(string(colorCyan), "• set dialtime <sec>")
	fmt.Println(string(colorReset), " - Sets the dial timeout you want to use.")
	fmt.Print(string(colorCyan), "• set showclosed <true/false>")
	fmt.Println(string(colorReset), " - Toggle the display of open ports on and off.")
	fmt.Print(string(colorCyan), "• set ports <mostknown/all>")
	fmt.Println(string(colorReset), " - Sets the ports you want to scan.")
	fmt.Print(string(colorCyan), "• show details")
	fmt.Println(string(colorReset), " - Prints the details.")
	fmt.Print(string(colorCyan), "• clear")
	fmt.Println(string(colorReset), " - Clears the terminal.")
	fmt.Print(string(colorCyan), "• motd")
	fmt.Println(string(colorReset), " - Prints the banner.")
	fmt.Print(string(colorCyan), "• run")
	fmt.Println(string(colorReset), " - Runs the scan.")

	fmt.Println()
}
