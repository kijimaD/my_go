// select文を使って待機する

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-task:
				fmt.Println("get", i)
			default:
				fmt.Println("キャンセルされていない")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second) // ここで待つ
	for i := 0; 5 > i; i++ {
		task <- i // チャネルに値が入る
	}
	cancel() // 完了
}
