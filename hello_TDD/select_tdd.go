package main

import (
	"fmt"
	"net/http"
	"time"
)

//func Racer(url string, url2 string) string {
//	aDuration := measureResponseTime(url)
//	bDuration := measureResponseTime(url2)
//
//	if aDuration > bDuration {
//		return url2
//	}
//
//	return url
//}
//
//func measureResponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}

var tenSecondTimeout = 10 * time.Second

func Racer(url, url2 string) (string, error) {
	return ConfigurableRacer(url, url2, tenSecondTimeout)
}

func ConfigurableRacer(url, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url):
		return url, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout")
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
