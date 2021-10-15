package ip

import (
	"strconv"
	"strings"
)

func Type(ip string) string {
	splitted := strings.Split(ip, ".")

	converted1, _ := strconv.Atoi(splitted[0])
	converted2, _ := strconv.Atoi(splitted[1])

	if converted1 == 10 {
		return "Private"
	}

	if converted1 == 127 {
		return "Loopback"
	}

	if converted1 == 172 && converted2 == 16 {
		return "Private"
	}

	if converted1 == 192 && converted2 == 168 {
		return "Private"
	}

	if converted1 == 0 {
		return "Invalid"
	}

	if converted1 == 169 && converted2 == 254 {
		return "Private"
	}

	return "Public"
}
