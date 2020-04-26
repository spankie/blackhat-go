package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func scan(address string, port int) {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		fmt.Printf("Port %d is closed\n", port)
	} else {
		fmt.Printf("Port %d is open\n", port)
	}
	wg.Done()
}

func main() {
	// 100 works fine
	// 1000 pauses on my laptop. the program does not complete for a while
	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go scan("scanme.nmap.org", i)
	}
	wg.Wait()
}
