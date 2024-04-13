package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {

	target := "testhtml5.vulnweb.com"
	var addresses []string
	wg := sync.WaitGroup{}

	for port := 1; port <= 1024; port++ {
		addresses = append(addresses, fmt.Sprintf("%s:%d", target, port))
	}

	wg.Add(len(addresses))

	for _, address := range addresses {
		go checkPort(address)
	}

	wg.Wait()
}

func checkPort(address string) bool {
	conn, err := net.DialTimeout("tcp", address, time.Millisecond)
	fmt.Printf("Checking address %s\n", address)
	if err != nil {
		return false
	}
	conn.Close()
	fmt.Printf("Address %s is open\n", address)
	return true
}
