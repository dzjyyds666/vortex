package vortex

import (
	"context"
	"testing"
)

func Test_Vortex(t *testing.T) {
	ctx := context.Background()
	BootStrap(ctx)
}
