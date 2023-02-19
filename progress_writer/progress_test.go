package progress

import (
	"context"
	"fmt"
	"testing"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	c := FromContext(ctx)
	fmt.Printf("%#v", c)
}
