package workflow

import (
	"github.com/mohan2020coder/go-workflows/internal/workflowstate"
)

// WorkflowInstance returns the current workflow instance.
func WorkflowInstance(ctx Context) *Instance {
	wfState := workflowstate.WorkflowState(ctx)
	return wfState.Instance()
}
