package ip

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PublicIP() {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"

	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("An unknown error occurred.")
		return
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("An unknown error occurred.")
		return
	}

	fmt.Print(string(colorCyan), " • Your public IP: ")
	fmt.Print(string(colorReset), "")
	fmt.Printf("%s\n", ip)
	fmt.Println()
}
