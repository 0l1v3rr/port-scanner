package help

import "fmt"

func PrintMotd() {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorCyan := "\033[36m"

	fmt.Println(string(colorGreen), "- port-scanner v2.0.1")
	fmt.Println(string(colorCyan), "- https://github.com/0l1v3rr/port-scanner")
	fmt.Print(string(colorReset), " - Type ")
	fmt.Print(string(colorRed), "help")
	fmt.Println(string(colorReset), "for more.")
	fmt.Print(string(colorReset), " - Type ")
	fmt.Print(string(colorRed), "exit")
	fmt.Println(string(colorReset), "for exit.")
	fmt.Println()
}

func PrintLogo() {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"

	fmt.Println(string(colorCyan), "   ____________   _  ___  _________ ")
	fmt.Println(string(colorCyan), "  / __/ ___/ _ | / |/ / |/ / __/ _ \\")
	fmt.Println(string(colorCyan), " _\\ \\/ /__/ __ |/    /    / _// , _/")
	fmt.Println(string(colorCyan), "/___/\\___/_/ |_/_/|_/_/|_/___/_/|_| ")
	fmt.Println(string(colorReset), "")
}
