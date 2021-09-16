package help

import "fmt"

func PrintHelpMenu() {
	fmt.Println(" ───────────────────────────────────")
	fmt.Println("\tPORT SCANNER HELP MENU")
	fmt.Println(" ───────────────────────────────────")
	fmt.Println(" --protocol \t The protocol (tcp/udp). Default is TCP.")
	fmt.Println("")
}
