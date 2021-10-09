package cli

import (
	"bytes"
	"fmt"
	"os/exec"
)

func UpdateNeeded() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	colorYellow := "\033[33m"

	cmd := exec.Command("git", "fetch", "--dry-run")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()

	if err != nil {
		fmt.Print(string(colorRed), " [!] Could not find Git installed.")
		fmt.Println(string(colorReset))
		return
	}

	if outb.String() == "" || len(outb.String()) < 1 {
		fmt.Print(string(colorGreen), " [!] The port-scanner is up to date.")
		fmt.Println(string(colorReset))
	} else {
		fmt.Print(string(colorYellow), " [!] You are using an outdated version of the port scanner.")
		fmt.Println(string(colorReset))
	}
}

func IsUpdateNeeded() bool {
	cmd := exec.Command("git", "fetch", "--dry-run")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()

	if err != nil {
		return false
	}

	if outb.String() == "" || len(outb.String()) < 1 {
		return false
	} else {
		return true
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
