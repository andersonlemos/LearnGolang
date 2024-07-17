package runner

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const durationLimit = 10 * time.Second

func ConfigurableRunner(a, b string, duration time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(durationLimit):
		return "", errors.New(fmt.Sprintf("timed out waiting for %s and %s", a, b))
	}
}

func Runner(a, b string) (string, error) {
	return ConfigurableRunner(a, b, durationLimit)
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		_, _ = http.Get(URL)
		ch <- true
	}()
	return ch
}

func measureTimeWithDelay(URL string) time.Duration {
	startTime := time.Now()
	_, _ = http.Get(URL)
	return time.Since(startTime)
}
