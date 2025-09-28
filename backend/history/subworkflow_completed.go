package history

import "github.com/mohan2020coder/go-workflows/backend/payload"

type SubWorkflowCompletedAttributes struct {
	Result payload.Payload `json:"result,omitempty"`
}
