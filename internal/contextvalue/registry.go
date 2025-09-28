package contextvalue

import (
	"github.com/mohan2020coder/go-workflows/internal/sync"
	"github.com/mohan2020coder/go-workflows/registry"
)

type registryKey struct{}

func WithRegistry(ctx sync.Context, r *registry.Registry) sync.Context {
	return sync.WithValue(ctx, registryKey{}, r)
}

func GetRegistry(ctx sync.Context) *registry.Registry {
	if v := ctx.Value(registryKey{}); v != nil {
		return v.(*registry.Registry)
	}

	return nil
}
