package util

import (
	"errors"
)

var (
	ErrTimeout = errors.New("timeout")
)

func Poll(tickDur time.Duration, timeout time.Duration, fn func() bool) error {
	timeoutCh := time.After(timeout)
	tickCh := time.Tick(tickDur)
	if fn() {
		return nil
	}
	for {
		select {
		case <-timeoutCh:
			return ErrTimeout
		case <-tickCh:
			if fn() {
				return nil
			}
		}
	}
}
