package util

import (
	"orbit_nft/logger"
	"time"
)

// WaitUntil wait function for a d duration, if function is not finished yet, continue
func WaitUntil(f func(), d time.Duration) {
	w := make(chan bool)
	go func() {
		f()
		w <- true
	}()

	select {
	case <-w:
		return
	case <-time.After(d):
		logger.Warnf("function execute more than %v", d)
		return
	}
}
