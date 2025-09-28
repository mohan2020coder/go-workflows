package test

import (
	"context"

	"github.com/mohan2020coder/go-workflows/backend"
	"github.com/mohan2020coder/go-workflows/backend/history"
)

type TestBackend interface {
	backend.Backend

	GetFutureEvents(ctx context.Context) ([]*history.Event, error)
}
