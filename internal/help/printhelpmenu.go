package help

import "fmt"

func PrintMotd() {
	colorYellow := "\033[33m"
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorCyan := "\033[36m"

	fmt.Println(string(colorGreen), "• port-scanner v2.0.1")
	fmt.Println(string(colorCyan), "• https://github.com/0l1v3rr/port-scanner")
	fmt.Print(string(colorReset), " • Type ")
	fmt.Print(string(colorYellow), "help")
	fmt.Println(string(colorReset), "for more.")
	fmt.Print(string(colorReset), " • Type ")
	fmt.Print(string(colorYellow), "exit")
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
