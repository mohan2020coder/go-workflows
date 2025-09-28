package executor

import (
	"context"

	"github.com/mohan2020coder/go-workflows/core"
)

type Cache interface {
	Store(ctx context.Context, instance *core.WorkflowInstance, workflow WorkflowExecutor) error
	Evict(ctx context.Context, instance *core.WorkflowInstance) error
	Get(ctx context.Context, instance *core.WorkflowInstance) (WorkflowExecutor, bool, error)
	StartEviction(ctx context.Context)
}
