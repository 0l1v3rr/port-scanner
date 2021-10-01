package port

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func PrintAddresses(protocol string, hostname string, port int, dialtime int) {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Duration(dialtime)*time.Second)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(60 * time.Second)
			ScanPort(protocol, hostname, port, dialtime)
		} else {
			return
		}
	}
	defer conn.Close()

	fmt.Printf("Remote address: %s\n", conn.RemoteAddr().String())
	fmt.Printf("Local address:  %s\n", conn.LocalAddr().String())
}
