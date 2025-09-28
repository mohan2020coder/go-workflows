package history

import (
	"github.com/mohan2020coder/go-workflows/core"
)

type SubWorkflowCancellationRequestedAttributes struct {
	SubWorkflowInstance *core.WorkflowInstance `json:"sub_workflow_instance,omitempty"`
}
