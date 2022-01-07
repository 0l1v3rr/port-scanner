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

	fmt.Printf("Remote address: \u001b[31;1m%s\u001b[0m\n", conn.RemoteAddr().String())
}
