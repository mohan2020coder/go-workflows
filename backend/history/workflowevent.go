package history

import "github.com/mohan2020coder/go-workflows/core"

// WorkflowEvent is a event addressed for a specific workflow instance
type WorkflowEvent struct {
	WorkflowInstance *core.WorkflowInstance `json:"workflow_instance,omitempty"`

	HistoryEvent *Event `json:"history_event,omitempty"`
}
