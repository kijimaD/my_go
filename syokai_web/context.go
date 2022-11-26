// 処理を中断する

package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		return
	}
	fmt.Println("キャンセルされていない")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx) // 表示される
	cancel()
	child(ctx) // 表示されない
}
