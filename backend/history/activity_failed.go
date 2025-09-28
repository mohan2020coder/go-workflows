package history

import "github.com/mohan2020coder/go-workflows/internal/workflowerrors"

type ActivityFailedAttributes struct {
	Error *workflowerrors.Error `json:"error,omitempty"`
}
