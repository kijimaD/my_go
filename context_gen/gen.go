package gen

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		userID, authToken, traceID := ctx.Value("userID").(int), ctx.Value("authToken").(string), ctx.Value("traceID").(int)
		fmt.Println("log: ", userID, authToken, traceID) // log:  2 xxxxxxxx 3
		fmt.Println("generator closed")
	}()
	return out
}

func output(ctx context.Context) {
	log := ctx.Value(ctxkey).(map[string]string)
	fmt.Println(log)
}

var ctxkey = "Log"

func add(ctx context.Context, key string, value string) context.Context {
	result := map[string]string{}
	m := ctx.Value(ctxkey).(map[string]string)
	for k, v := range m {
		result[k] = v
	}
	result[key] = value
	ctx = context.WithValue(ctx, ctxkey, result)
	return ctx
}
