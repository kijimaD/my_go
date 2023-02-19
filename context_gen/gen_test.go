package gen

import (
	"context"
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ctx = context.WithValue(ctx, "userID", 2)
	ctx = context.WithValue(ctx, "authToken", "xxxxxx")
	ctx = context.WithValue(ctx, "traceID", 3)
	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}

func TestOutput(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, ctxkey, map[string]string{})
	ctx = add(ctx, "job1", "aaa")
	ctx = add(ctx, "job2", "bbb")
	output(ctx)
}
