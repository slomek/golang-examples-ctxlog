package ctxlog

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), "label", "depdep")
	Println(ctx, "Hello world!")
}
