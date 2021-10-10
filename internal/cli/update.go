package cli

import (
	"fmt"
	"os/exec"
	"strings"
)

func UpdateNeeded() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	colorYellow := "\033[33m"

	out, err := exec.Command("git", "fetch", "--dry-run").Output()

	if err != nil {
		fmt.Print(string(colorRed), " [!] Could not find Git installed.")
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

func IsUpdateNeeded() bool {
	out, err := exec.Command("git", "fetch", "--dry-run").Output()

	if err != nil {
		return false
	}

	if strings.Contains(string(out), "remote: ") {
		return true
	} else {
		return false
	}
}

func Update() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	cmd := exec.Command("git", "pull")
	err := cmd.Run()

	if err != nil {
		fmt.Print(string(colorRed), "An error occurred.")
		fmt.Println(string(colorReset))
		return
	}

	fmt.Print(string(colorGreen), "Successful update! Restart the program.")
	fmt.Println(string(colorReset))

}
