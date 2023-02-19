package progress

import (
	"context"
	"fmt"
	"testing"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	init := progressWriter{false, &progressReader{}, map[string]interface{}{"a": 1}}
	ctx = context.WithValue(ctx, contextKey, &init)

	ctx.Value(contextKey)
	// fmt.Printf("%#v", v)

	pw, _, _ := NewFromContext(ctx)
	fmt.Printf("%#v", pw)

	// oneOffProgress(ctx, "iii")
}
