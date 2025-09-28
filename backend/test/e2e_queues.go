package test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mohan2020coder/go-workflows/client"
	"github.com/mohan2020coder/go-workflows/core"
	"github.com/mohan2020coder/go-workflows/worker"
	"github.com/mohan2020coder/go-workflows/workflow"
	"github.com/stretchr/testify/require"
)

var e2eQueueTests = []backendTest{
	{
		name: "Queues/OnlyPullsWorkflowsFromGivenQueues",
		f: func(t *testing.T, ctx context.Context, c *client.Client, w *worker.Worker, b TestBackend) {
			wf := func(ctx workflow.Context) (bool, error) {
				return true, nil
			}

			wfNever := func(ctx workflow.Context) (bool, error) {
				panic("should never be called")
			}

			register(t, ctx, w, []interface{}{wf, wfNever}, nil)

			instance, err := c.CreateWorkflowInstance(ctx, client.WorkflowInstanceOptions{
				InstanceID: uuid.NewString(),
			}, wf)
			require.NoError(t, err)

			neverInstance, err := c.CreateWorkflowInstance(ctx, client.WorkflowInstanceOptions{
				InstanceID: uuid.NewString(),
				Queue:      "Never",
			}, wfNever)
			require.NoError(t, err)

			ns, err := c.GetWorkflowInstanceState(ctx, neverInstance)
			require.NoError(t, err)
			require.Equal(t, core.WorkflowInstanceStateActive, ns)

			_, err = client.GetWorkflowResult[any](ctx, c, instance, time.Second*10)
			require.NoError(t, err)
		},
	},
}
