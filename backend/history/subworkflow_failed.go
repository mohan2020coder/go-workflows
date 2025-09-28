package history

import "github.com/mohan2020coder/go-workflows/internal/workflowerrors"

type SubWorkflowFailedAttributes struct {
	Error *workflowerrors.Error `json:"error,omitempty"`
}
