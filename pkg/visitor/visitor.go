package visitor

import (
	"fmt"
	"net"
	"sync"

	v1 "github.com/imneov/kube-frp/pkg/config/v1"
	"github.com/imneov/kube-frp/pkg/util/ipallocator"
)

var (
	ipAllocators  = make(map[string]*ipallocator.IPAllocator)
	ipAllocatorMu sync.Mutex
)

func getIPAllocator(cidr string) (*ipallocator.IPAllocator, error) {
	ipAllocatorMu.Lock()
	defer ipAllocatorMu.Unlock()

	if allocator, ok := ipAllocators[cidr]; ok {
		return allocator, nil
	}

	allocator, err := ipallocator.NewIPAllocator(cidr)
	if err != nil {
		return nil, err
	}

	ipAllocators[cidr] = allocator
	return allocator, nil
}

func allocateIPAndPort(cfg *v1.VisitorBaseConfig) (string, int, error) {
	var ip net.IP

	if cfg.IPRange != "" {
		allocator, err := getIPAllocator(cfg.IPRange)
		if err != nil {
			return "", 0, err
		}

		ip, err = allocator.Allocate()
		if err != nil {
			return "", 0, err
		}
	} else {
		ip = net.ParseIP(cfg.BindAddr)
		if ip == nil {
			return "", 0, fmt.Errorf("invalid bind address: %s", cfg.BindAddr)
		}
	}

	port := cfg.BindPort
	if cfg.AutoAssignPort {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			return "", 0, err
		}
		addr := listener.Addr().(*net.TCPAddr)
		port = addr.Port
		listener.Close()
	}

	return ip.String(), port, nil
}
