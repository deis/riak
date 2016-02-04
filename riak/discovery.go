package riak

import (
	"errors"
	"net"
)

const (
	discoveryHostName = "deis-riak-discovery"
)

var (
	ErrNoDiscoveryAddrs = errors.New("no discovery addresses found")
)

func getDiscoveryIP(hostName string) (string, error) {
	discoveryAddrs, err := net.LookupHost(hostName)
	if err != nil {
		return "", err
	}
	for _, addr := range discoveryAddrs {
		if addr != localIP() && !net.ParseIP(addr).IsLoopback() {
			return addr, nil
		}
	}
	return "", ErrNoDiscoveryAddrs
}

// localIP returns the non loopback local IP of the host
func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
