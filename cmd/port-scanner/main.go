package main

import (
	"fmt"

	pt "github.com/0l1v3rr/port-scanner/pkg/port"
)

func main() {
	port := 80
	open := pt.ScanPort("tcp", "localhost", port)

	if open {
		fmt.Printf("The port %v is opened.", port)
	} else {
		fmt.Printf("The port %v is closed.", port)
	}

}
