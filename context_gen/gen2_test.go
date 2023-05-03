package gen

import (
	"context"
	"fmt"
	"testing"
)

// goroutineでないので、contextの値が伝播してないっぽい?
func TestOneffProgress(t *testing.T) {
	pw := progressWriter{reader: "4"}
	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, ctxkey, &pw)
	done := oneOffProgress(ctx, "exporting layers")
	_ = done(nil)

	fmt.Println(ctx.Value(ctxkey))
	// ここで値が変更されていてほしいが、できてない
}
