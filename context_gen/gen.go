package gen

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int, userID int, authToken string, traceID int) <-chan int {
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
		fmt.Println("log: ", userID, authToken, traceID) // log:  2 xxxxxxxx 3
		fmt.Println("generator closed")
	}()
	return out
}
