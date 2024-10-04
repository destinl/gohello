package utils

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func TcpScanner() {
	hostname := flag.String("hostname", "", "hostname to scan")
	startPort := flag.Int("startPort", 80, "start port to scan")
	endPort := flag.Int("endPort", 100, "end port to scan")
	timeout := flag.Duration("timeout", 200*time.Millisecond, "timeout for each port")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("Open ports: %v\n", ports)
}

func isOpen(hostname string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", hostname, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}
