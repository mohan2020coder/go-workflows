package history

import "github.com/mohan2020coder/go-workflows/backend/payload"

type SignalReceivedAttributes struct {
	Name string          `json:"name,omitempty"`
	Arg  payload.Payload `json:"arg,omitempty"`
}
