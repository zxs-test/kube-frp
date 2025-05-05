package ipallocator

import (
	"fmt"
	"net"
	"sync"
)

type IPAllocator struct {
	mu      sync.Mutex
	network *net.IPNet
	used    map[string]bool
}

func NewIPAllocator(cidr string) (*IPAllocator, error) {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %v", err)
	}

	return &IPAllocator{
		network: network,
		used:    make(map[string]bool),
	}, nil
}

func (a *IPAllocator) Allocate() (net.IP, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Start from the first usable IP in the network
	ip := make(net.IP, len(a.network.IP))
	copy(ip, a.network.IP)
	ip = ip.To4()

	// Skip network address
	ip[3]++

	for {
		// Check if we've reached the broadcast address
		if ip.Equal(a.network.IP) {
			return nil, fmt.Errorf("no available IP addresses in range")
		}

		// Check if this IP is already in use
		if !a.used[ip.String()] {
			a.used[ip.String()] = true
			return ip, nil
		}

		// Increment IP
		for i := 3; i >= 0; i-- {
			ip[i]++
			if ip[i] != 0 {
				break
			}
		}
	}
}

func (a *IPAllocator) Release(ip net.IP) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.used, ip.String())
}
