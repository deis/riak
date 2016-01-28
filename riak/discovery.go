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
	} else if len(discoveryAddrs) == 0 {
		return "", ErrNoDiscoveryAddrs
	}
	return discoveryAddrs[0], nil
}
