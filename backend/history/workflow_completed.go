package history

import (
	"github.com/mohan2020coder/go-workflows/backend/payload"
	"github.com/mohan2020coder/go-workflows/internal/workflowerrors"
)

type ExecutionCompletedAttributes struct {
	Result payload.Payload       `json:"result,omitempty"`
	Error  *workflowerrors.Error `json:"error,omitempty"`
}
