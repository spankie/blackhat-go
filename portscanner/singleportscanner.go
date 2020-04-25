package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		fmt.Println("Port 80 is open")
	} else {
		fmt.Println("Port 80 is not open")
	}
}
