package port

import (
	"net"
	"strconv"
	"strings"
	"time"
)

func ScanPort(protocol string, hostname string, port int, dialtime int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Duration(dialtime)*time.Second)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(60 * time.Second)
			ScanPort(protocol, hostname, port, dialtime)
		} else {
			return false
		}
	}
	defer conn.Close()

	return true
}
