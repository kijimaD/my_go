package gen

import (
	"context"
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator(ctx, 1, 2, "xxxxxxxx", 3)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}
