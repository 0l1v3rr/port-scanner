package ip

import (
	"strconv"
	"strings"
)

func Class(ip string) string {
	splitted := strings.Split(ip, ".")

	converted, _ := strconv.Atoi(splitted[0])
	if converted < 128 {
		return "A"
	}

	if converted < 192 {
		return "B"
	}

	if converted < 224 {
		return "C"
	}

	return "Undefined"
}
