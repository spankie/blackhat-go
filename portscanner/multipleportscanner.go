package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i < 1024; i++ {
		_, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
		if err != nil {
			fmt.Printf("Port %d is closed\n", i)
		} else {
			fmt.Printf("Port %d is open\n", i)
		}
	}
}
