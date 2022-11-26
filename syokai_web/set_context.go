// context.Context型の値にデータを設定する

package main

import (
	"context"
	"fmt"
)

type TraceID string

const ZeroTraceID = ""

type traceIDKey struct{}

func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
	// キーには空構造体を用いて独自型を使うのが一般的。衝突させないため
}

func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
