# Port Scanner
This is a simple terminal application that allows you to find opened ports on your network.<br>
Inspired by [nmap](https://github.com/nmap/nmap)

## Setup
First, you need to download [Go](https://golang.org/dl/)<br>
Then clone this repository:
```sh
git clone https://github.com/0l1v3rr/port-scanner.git
cd port-scanner
```
On Linux or Mac, use the `make run` command, to run the app.<br>
On Windows: `go run cmd\port-scanner\main.go`

## Flags
`--protocol` -> tcp/udp - default is tcp.<br>
**Example usage:** `go run cmd\port-scanner\main.go --protocol udp`
<br><br>
`--ip` -> The IP address you want to scan. Default is your device's IP.<br>
**Example usage:** `go run cmd\port-scanner\main.go --ip 192.168.0.1`
<br><br>
`--port` -> The only port you want to scan. By default, the app will scan the most known ports.<br>
**Example usage:** `go run cmd\port-scanner\main.go --port 443`
<br><br>
`--closed` -> With this flag, the app won't show the closed ports.<br>
**Example usage:** `go run cmd\port-scanner\main.go --closed`

### You can combine the flags.
For example: `go run cmd\port-scanner\main.go --ip 192.168.0.1 --port 23`<br>
The line above means that you are scanning port **23** of IP **192.168.0.1**.