package chans

import (
	"errors"
	"testing"
	"time"
)

const (
	maxWaitDir = 200 * time.Millisecond
)

func joinErrsCh(c1, c2 <-chan error) <-chan error {
	ret := make(chan error)
	go func() {
		ret <- JoinErrs(c1, c2)
	}()
	return ret
}

func TestJoinErrs(t *testing.T) {
	err1 := errors.New("testing error 1")
	err2 := errors.New("testing error 2")
	c1 := make(chan error)
	c2 := make(chan error)
	go func() {
		c1 <- err1
	}()
	select {
	case err := <-joinErrsCh(c1, c2):
		if err != err1 {
			t.Errorf("error returned from JoinErrs was %s, expected %s", err, err1)
		}
	case <-time.After(maxWaitDir):
		t.Errorf("waited %s but got no response from JoinErrs", maxWaitDir)
	}

	go func() {
		c2 <- err2
	}()
	select {
	case err := <-joinErrsCh(c1, c2):
		if err != err2 {
			t.Errorf("error returned from JoinErrs was %s, expected %s", err, err2)
		}
	case <-time.After(maxWaitDir):
		t.Errorf("waited %s but got no response from JoinErrs", maxWaitDir)
	}
}
