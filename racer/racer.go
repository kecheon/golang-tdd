package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	aDuration := measureDuration(a)
	bDuration := measureDuration(b)

	if aDuration > bDuration {
		return b
	}
	return a
}

func measureDuration(url string) (duration time.Duration) {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

const tenSecondTimeout = 10 * time.Second

func RacerPing(a, b string, timeout time.Duration) (winner string, err error) {
	return ConfigurableRacer(a, b, timeout)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timeout error")
	}
}
