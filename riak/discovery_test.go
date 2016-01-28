package riak

import (
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
}
