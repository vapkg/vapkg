package core

import (
	"os"
)

type IHttpProvider interface {
	Type() ProviderType
	GetPath(dep *VaPackageDependence) (string, error)
}

type ProviderHandlerFn func(ctx *Context, depIdx int, outp string) *os.File

type ProviderRegistry struct {
	providers map[ProviderType]IHttpProvider
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{make(map[ProviderType]IHttpProvider)}
}

func CreateProviderRegistry() ProviderRegistry {
	return ProviderRegistry{make(map[ProviderType]IHttpProvider)}
}

func (r *ProviderRegistry) Register(t ProviderType, inst IHttpProvider) {
	r.providers[t] = inst
}
