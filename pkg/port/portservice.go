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
	case 554:
		return "rtsp"
	case 546:
		return "dhcpv6"
	case 547:
		return "dhcpv6"
	case 560:
		return "rmonitor"
	case 563:
		return "nntp over ssl"
	case 587:
		return "smtp"
	case 591:
		return "filemaker"
	case 593:
		return "microsoft dcom"
	case 631:
		return "internet printing"
	case 636:
		return "dlap over ssl"
	case 639:
		return "msdp (pim)"
	case 646:
		return "ldp (mpls)"
	case 691:
		return "ms exchange"
	case 860:
		return "iscsi"
	case 873:
		return "rsync"
	case 902:
		return "vmware server"
	case 989:
		return "ftp over ssl"
	case 990:
		return "ftp over ssl"
	case 993:
		return "imap4 over ssl"
	case 995:
		return "pop3 over ssl"
	case 1025:
		return "microsoft rpc"
	case 1080:
		return "socks proxy/mydooom"
	case 1194:
		return "openvpn"
	case 1214:
		return "kazaa"
	case 1241:
		return "nessus"
	case 1311:
		return "deil openmanage"
	case 1337:
		return "waste"
	case 1433:
		return "microsoft sql"
	case 1434:
		return "microsoft sql"
	case 1512:
		return "wins"
	case 1589:
		return "cisco vqp"
	case 1701:
		return "l2tp"
	case 1723:
		return "ms pptp"
	case 1725:
		return "steam"
	case 1741:
		return "ciscoworks 2000"
	case 1755:
		return "ms media server"
	case 1812:
		return "radius"
	case 1813:
		return "radius"
	case 1863:
		return "msn"
	case 1985:
		return "cisco hsrp"
	case 2000:
		return "cisco scco"
	case 2002:
		return "cisco asc"
	case 2049:
		return "nfs"
	case 2082:
		return "cpanel"
	case 2083:
		return "cpanel"
	case 2100:
		return "oracle xdb"
	case 2222:
		return "directadmin"
	case 2302:
		return "halo"
	case 2483:
		return "oracle db"
	case 2484:
		return "oracle db"
	}

	return "unknown"
}
