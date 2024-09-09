package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

func scanPort(address string, proc string, port int) bool {
	host := fmt.Sprintf("%s:%d", address, port)
	con, err := net.Dial(proc, host)
	if err != nil {
		return false
	}
	con.Close()
	return true
}

func ScanChunkPorts(wg *sync.WaitGroup, host string, proc string, ports []int) {

	defer wg.Done()
	for _, port := range ports {
		if scanPort(host, proc, port) {
			fmt.Printf("Port %d is Open.\n", port)
		}
	}

}

func run(host string, proc string, thread int, ports []int) {
	var wg sync.WaitGroup

	chunk := len(ports) / thread
	for i := 0; i < thread; i++ {
		wg.Add(1)
		start := i * chunk
		end := (i * chunk) + chunk

		go ScanChunkPorts(&wg, host, proc, ports[start:end])
	}
	wg.Wait()
}

func main() {

	host := flag.String("host", "127.0.0.1", "provide a host to scan")
	proc := flag.String("proc", "tcp", "provide a protocol")
	ports := flag.String("ports", "1-1000", "ports to scan example : [ '1-1000' or '80,81,9000' ] ")
	thread := flag.Int("thread", 1, "how many thread do you wanna use")
	flag.Parse()

	var ps []int
	if strings.Contains(*ports, "-") {
		tmp := strings.Split(*ports, "-")
		start, err1 := strconv.ParseInt(tmp[0], 10, 64)
		end, err2 := strconv.ParseInt(tmp[1], 10, 64)
		if err1 != nil || err2 != nil {
			panic("invalid port number")
		}
		for i := start; i <= end; i++ {
			ps = append(ps, int(i))
		}
	} else if strings.Contains(*ports, ",") {
		tmp := strings.Split(*ports, ",")
		for _, item := range tmp {
			val, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				panic("invalid port number")
			}
			ps = append(ps, int(val))
		}

	} else {
		val, err := strconv.ParseInt(*ports, 10, 64)
		if err != nil {
			panic("invalid port number")
		}
		ps = append(ps, int(val))
	}

	run(*host, *proc, *thread, ps)
	// fmt.Println("result : ", *host, *proc, *ports, *thread, ps)

}
