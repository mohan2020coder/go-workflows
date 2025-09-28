package workflow

import (
	"github.com/mohan2020coder/go-workflows/internal/sync"
)

type (
	Context   = sync.Context
	WaitGroup = sync.WaitGroup
)

// NewWaitGroup creates a new WaitGroup instance.
func NewWaitGroup() WaitGroup {
	return sync.NewWaitGroup()
}

// Go spawns a workflow "goroutine".
func Go(ctx Context, f func(ctx Context)) {
	sync.Go(ctx, f)
}
