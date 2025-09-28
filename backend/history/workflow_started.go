package history

import (
	"github.com/mohan2020coder/go-workflows/backend/metadata"
	"github.com/mohan2020coder/go-workflows/backend/payload"
	"github.com/mohan2020coder/go-workflows/core"
)

type ExecutionStartedAttributes struct {
	Queue core.Queue `json:"queue,omitempty"`

	Name string `json:"name,omitempty"`

	Metadata *metadata.WorkflowMetadata `json:"metadata,omitempty"`

	Inputs []payload.Payload `json:"inputs,omitempty"`

	WorkflowSpanID [8]byte `json:"workflowSpanID,omitempty"`
}
