package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func PrintAddresses(hostname string) {
	address := hostname + ":" + strconv.Itoa(80)
	conn, err := net.DialTimeout("tcp", address, time.Duration(5)*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	fmt.Printf("Remote address: %s\n", conn.RemoteAddr().String())
}
