package history

import "github.com/mohan2020coder/go-workflows/backend/payload"

type TraceStartedAttributes struct {
	SpanID payload.Payload `json:"spanID"`
}
