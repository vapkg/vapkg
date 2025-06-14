package core

import (
	"os"
)

type ProviderHandlerFn func(ctx *Context, depIdx int, outp string) *os.File

type ProviderRegistry struct {
	providers map[string]ProviderHandlerFn
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{make(map[string]ProviderHandlerFn)}
}

func CreateProviderRegistry() ProviderRegistry {
	return ProviderRegistry{make(map[string]ProviderHandlerFn)}
}

func (r *ProviderRegistry) Register(name string, fn ProviderHandlerFn) {
	r.providers[name] = fn
}
