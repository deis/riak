package riak

import (
	"net"
	"testing"
)

func TestGetDiscoveryIP(t *testing.T) {
	// getting the default hostname won't work b/c we're not in a k8s cluster with that service running
	addr, err := getDiscoveryIP(discoveryHostName)
	if err == nil {
		t.Errorf("getting IP for %s hostname was expected to fail", discoveryHostName)
	}
	if addr != "" {
		t.Errorf("failed call for hostname %s returned non-empty address %s", discoveryHostName, addr)
	}

	addr, err = getDiscoveryIP("google.com")
	if err != nil {
		t.Errorf("getting IP for google.com hostname was expected to not fail")
	}
	if addr == "" {
		t.Errorf("call for hostname google.com returned empty address")
	}
	addr, err = getDiscoveryIP("localhost")
	if err != ErrNoDiscoveryAddrs {
		t.Errorf("getting IP from loopback address was expected to fail; got %s", addr)
	}
	addr, err = getDiscoveryIP(localIP())
	if err != ErrNoDiscoveryAddrs {
		t.Errorf("getting IP from local IP address was expected to fail; got %s", addr)
	}

}

func TestGetLocalIP(t *testing.T) {
	// We live in a connected world. These tests assume you have a non-loopback address
	if localIP() == "" {
		t.Errorf("could not retrieve the local IP address. Are you connected to a network?")
	}
	if ip := net.ParseIP(localIP()); ip == nil {
		t.Errorf("could not parse local IP to a net.IP")
	}
}
