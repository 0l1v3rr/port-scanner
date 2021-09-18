package help

import "fmt"

func PrintHelpMenu() {
	fmt.Println(" ───────────────────────────────────")
	fmt.Println("\tPORT SCANNER HELP MENU")
	fmt.Println(" ───────────────────────────────────")
	fmt.Println(" --protocol \t The protocol (tcp/udp). Default is TCP.")
	fmt.Println(" --ip \t\t The IP Address you want to scan. Default is 'localhost'.")
	fmt.Println(" --port \t The only port you want to scan.")
	fmt.Println("")
}
