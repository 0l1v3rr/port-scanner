package cli

import (
	"fmt"
	"os/exec"
	"strings"
)

func UpdateNeeded() {
	colorRed := "\033[31;1m"
	colorGreen := "\033[32;1m"
	colorReset := "\033[0m"
	colorYellow := "\033[33;1m"

	out, err := exec.Command("git", "fetch", "--dry-run").Output()

	if err != nil {
		fmt.Print(string(colorRed), " [!] An error occurred. Maybe you don't have git installed.")
		fmt.Println(string(colorReset))
		return
	}

	if strings.Contains(string(out), "remote: ") {
		fmt.Print(string(colorYellow), " [!] You are using an outdated version of the port scanner.")
		fmt.Println(string(colorReset))
	} else {
		fmt.Print(string(colorGreen), " [!] The port-scanner is up to date.")
		fmt.Println(string(colorReset))
	}

}

func Update() {
	colorGreen := "\033[32;1m"
	colorReset := "\033[0m"

	cmd := exec.Command("git", "pull")
	err := cmd.Run()

	if err != nil {
		fmt.Println("An error occurred.")
		return
	}

	fmt.Print(string(colorGreen), "Successful update! Restart the program.")
	fmt.Println(string(colorReset))

}
