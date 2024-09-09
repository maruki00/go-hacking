package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// func scanPort(address string, port int) error {
// 	host := fmt.Sprintf("%s:%d", address, port)
// 	con, err := net.Dial("tcp", host)
// }

// func run()

func main() {

	host := flag.String("host", "127.0.0.1", "provide a host to scan")
	proc := flag.String("proc", "tcp", "provide a protocol")
	ports := flag.String("ports", "1-1000", "ports to scan example : [ '1-1000' or '80,81,9000' ] ")
	thread := flag.Int("thread", 1, "how many thread do you wanna use")
	flag.Parse()

	// wg * sync.WaitGroup

	var ps []int64
	if strings.Contains(*ports, "-") {
		tmp := strings.Split(*ports, "-")
		start, err1 := strconv.ParseInt(tmp[0], 10, 64)
		end, err2 := strconv.ParseInt(tmp[1], 10, 64)
		if err1 != nil || err2 != nil {
			panic("invalid port number")
		}
		for i := start; i <= end; i++ {
			ps = append(ps, i)
		}
	} else if strings.Contains(*ports, ",") {
		tmp := strings.Split(*ports, "-")
		for _, item := range tmp {
			val, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				panic("invalid port number")
			}
			ps = append(ps, val)
		}

	} else {
		val, err := strconv.ParseInt(*ports, 10, 64)
		if err != nil {
			panic("invalid port number")
		}
		ps = append(ps, val)
	}

	fmt.Println("result : ", *host, *proc, *ports, *thread, ps)

}
