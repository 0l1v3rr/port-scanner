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
		return "socks proxy"
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
		return "cpanel ssl"
	case 2086:
		return "whm"
	case 2087:
		return "whm ssl"
	case 2095:
		return "webmail"
	case 2096:
		return "webmail ssl"
	case 2100:
		return "oracle xdb"
	case 2121:
		return "ccproxy-ftp"
	case 2222:
		return "directadmin"
	case 2302:
		return "halo"
	case 2483:
		return "oracle db"
	case 2484:
		return "oracle db"
	case 2745:
		return "bagle.h"
	case 2967:
		return "symantec av"
	case 3050:
		return "interbase db"
	case 3074:
		return "xbox live"
	case 3124:
		return "http proxy"
	case 3127:
		return "mydoom"
	case 3128:
		return "http proxy"
	case 3222:
		return "glbp"
	case 3260:
		return "iscsi target"
	case 3306:
		return "mysql"
	case 3389:
		return "terminal server"
	case 3689:
		return "itunes"
	case 3690:
		return "subversion"
	case 3724:
		return "world of warcraft"
	case 3784:
		return "ventrilo"
	case 3785:
		return "ventrilo"
	case 4333:
		return "msql"
	case 4444:
		return "blaster"
	case 4664:
		return "google desktop"
	case 4672:
		return "emule"
	case 4899:
		return "radmin"
	case 5000:
		return "upnp"
	case 5001:
		return "slingbox"
	case 5004:
		return "rtp"
	case 5005:
		return "rtp"
	case 5050:
		return "yahoo messenger"
	case 5060:
		return "sip"
	case 5190:
		return "aim/icq"
	case 5222:
		return "xmpp/jabber"
	case 5223:
		return "xmpp/jabber"
	case 5432:
		return "postgresql"
	case 5500:
		return "vnc server"
	case 5554:
		return "sasser"
	case 5800:
		return "vnc over http"
	case 5900:
		return "vnc server"
	case 6000:
		return "x11"
	case 6001:
		return "x11"
	case 6112:
		return "battle.net"
	case 6129:
		return "dameware"
	case 6257:
		return "winmx"
	case 6346:
		return "gnutella"
	case 6347:
		return "gnutella"
	case 6500:
		return "gamespy arcade"
	case 6566:
		return "sane"
	case 6588:
		return "analogx"
	case 6665:
		return "irc"
	case 6679:
		return "irc over ssl"
	case 6699:
		return "napster"
	case 6881:
		return "bittorrent"
	case 6999:
		return "bittorrent"
	case 6891:
		return "windows live"
	case 6970:
		return "quicktime"
	case 7212:
		return "ghostsurf"
	case 7648:
		return "cu-seeme"
	case 7649:
		return "cu-seeme"
	case 8000:
		return "internet radio"
	case 8080:
		return "http proxy"
	case 8086:
		return "kaspersky av"
	case 8087:
		return "kaspersky av"
	case 8118:
		return "privoxy"
	case 8200:
		return "vmware server"
	case 8500:
		return "adobe coldfusion"
	case 8767:
		return "teamspeak"
	case 8866:
		return "bagle.b"
	case 9100:
		return "hp jetdirect"
	case 9101:
		return "bacula"
	case 9102:
		return "bacula"
	case 9103:
		return "bacula"
	case 9119:
		return "mxit"
	case 9418:
		return "git"
	case 9800:
		return "webdav"
	case 9898:
		return "dabber"
	case 9988:
		return "rbot/spybot"
	case 9999:
		return "webmin"
	case 10000:
		return "webmin"
	case 10113:
		return "netlq"
	case 11371:
		return "openpgp"
	case 12035:
		return "second life"
	case 12036:
		return "second life"
	case 12345:
		return "netbus"
	case 13720:
		return "netbackup"
	case 13721:
		return "netbackup"
	case 14567:
		return "battlefield"
	case 15118:
		return "dipnet/oddbob"
	case 19226:
		return "adminsecure"
	case 19638:
		return "ensim"
	case 20000:
		return "usermin"
	case 24800:
		return "synergy"
	case 25999:
		return "xfine"
	case 27015:
		return "half-life"
	case 27374:
		return "sub7"
	case 28960:
		return "call of duty"
	case 31337:
		return "back orifice"
	case 33434:
		return "traceroute"
	}

	return "unknown"
}
