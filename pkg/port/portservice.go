package port

func PortServiceName(port int) string {
	switch port {
	case 7:
		return "echo"
	case 19:
		return "chargen"
	case 20:
		return "ftp"
	case 21:
		return "ftp"
	case 22:
		return "ssh"
	case 23:
		return "telnet"
	case 25:
		return "smtp"
	case 43:
		return "whois"
	case 49:
		return "tacacs"
	case 53:
		return "dns"
	case 67:
		return "dhcp/bootp"
	case 68:
		return "dhcp/bootp"
	case 69:
		return "tftp"
	case 70:
		return "gopher"
	case 79:
		return "finger"
	case 80:
		return "http"
	case 88:
		return "kerberos"
	case 102:
		return "ms exchange"
	case 110:
		return "pop3"
	case 111:
		return "rpcbind"
	case 113:
		return "ident"
	case 119:
		return "nntp"
	case 123:
		return "ntp"
	case 135:
		return "microsoft-rpc"
	case 139:
		return "netbios-ssn"
	case 143:
		return "imap4"
	case 161:
		return "snmp"
	case 162:
		return "snmp"
	case 177:
		return "xdmcp"
	case 179:
		return "bgp"
	case 201:
		return "appletalk"
	case 264:
		return "bgmp"
	case 318:
		return "tsp"
	case 319:
		return "ldap"
	case 411:
		return "direct connect"
	case 412:
		return "direct connect"
	case 443:
		return "https"
	case 445:
		return "microsoft-ds"
	case 464:
		return "kerberos"
	case 465:
		return "smtp over ssl"
	case 497:
		return "retrospect"
	case 500:
		return "isakmp"
	case 512:
		return "rexec"
	case 513:
		return "login"
	case 514:
		return "syslong"
	case 515:
		return "ldp/lpr"
	case 520:
		return "rip"
	case 521:
		return "ripng (ipv6)"
	case 540:
		return "uucp"
	}

	return "unknown"
}
