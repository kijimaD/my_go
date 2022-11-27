// レスポンスを返すのがどちらが早いか調べる
package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// ================
// 並列実行する

var tenSecondTimeout = 10 * time.Second

func Racer2(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// タイムアウトを引数指定してテストできる
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// 最初にチャネルに書き込む方はコードが実行されURLを返す
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
