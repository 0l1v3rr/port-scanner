package ip

import (
	"fmt"
	"strings"
)

func FirstIP(ip string) string {
	c := Class(ip)

	splitted := strings.Split(ip, ".")

	if c == "A" {
		return fmt.Sprintf("%v.0.0.1", splitted[0])
	}

	if c == "B" {
		return fmt.Sprintf("%v.%v.0.1", splitted[0], splitted[1])
	}

	return fmt.Sprintf("%v.%v.%v.1", splitted[0], splitted[1], splitted[2])
}

func LastIP(ip string) string {
	c := Class(ip)

	splitted := strings.Split(ip, ".")

	if c == "A" {
		return fmt.Sprintf("%v.255.255.254", splitted[0])
	}

	if c == "B" {
		return fmt.Sprintf("%v.%v.255.254", splitted[0], splitted[1])
	}

	return fmt.Sprintf("%v.%v.%v.254", splitted[0], splitted[1], splitted[2])
}

func Broadcast(ip string) string {
	c := Class(ip)

	splitted := strings.Split(ip, ".")

	if c == "A" {
		return fmt.Sprintf("%v.255.255.255", splitted[0])
	}

	if c == "B" {
		return fmt.Sprintf("%v.%v.255.255", splitted[0], splitted[1])
	}

	return fmt.Sprintf("%v.%v.%v.255", splitted[0], splitted[1], splitted[2])
}

func Prefix(ip string) string {
	c := Class(ip)

	splitted := strings.Split(ip, ".")

	if c == "A" {
		return fmt.Sprintf("%v.0.0.0", splitted[0])
	}

	if c == "B" {
		return fmt.Sprintf("%v.%v.0.0", splitted[0], splitted[1])
	}

	return fmt.Sprintf("%v.%v.%v.0", splitted[0], splitted[1], splitted[2])
}
