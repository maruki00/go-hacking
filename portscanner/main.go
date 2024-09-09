package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

func scanPort(address string, port int) error {
	host := fmt.Sprintf("%s:%d", address, port)
	con, err := net.Dial("tcp", host)
}

func main() {

	flag.String("host", "127.0.0.1", "provide a host to scan")
	flag.String("proc", "tcp", "provide a protocol")
	flag.String("ports", "1-1000", "ports to scan example : [ '1-1000' or '80,81,9000' ] ")
	flag.Int("thread", 1, "how many thread do you wanna use")
	flag.Parse()

	wg * sync.WaitGroup

}
