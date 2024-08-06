package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(slowURL, fastURL string) (winner string, error error) {
	return ConfigurableRacer(slowURL, fastURL, tenSecondTimeout)
}

func ConfigurableRacer(slowURL, fastURL string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(slowURL):
		return slowURL, nil
	case <-ping(fastURL):
		return fastURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", slowURL, fastURL)
	}
	// slowURLDuration := measureResponseTime(slowURL)
	//
	// fastURLDuration := measureResponseTime(fastURL)
	//
	// if slowURLDuration < fastURLDuration {
	// 	return slowURL
	// }
	//
	// return fastURL
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()

	http.Get(url)

	return time.Since(start)
}
