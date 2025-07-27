package core

import "vapkg/internal/core/vapkg"

type ProviderFactoryFn func(name string, data *vapkg.Provider) IProvider

type IProvider interface {
	GetName() string
	GetType() vapkg.ProviderType
	GetPath(dep *vapkg.Dependency) string
	GetFile(dep *vapkg.Dependency) string
}

type ProviderRegistry struct {
	providers map[vapkg.ProviderType]ProviderFactoryFn
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{make(map[vapkg.ProviderType]ProviderFactoryFn)}
}

func CreateProviderRegistry() ProviderRegistry {
	return ProviderRegistry{make(map[vapkg.ProviderType]ProviderFactoryFn)}
}

func (r *ProviderRegistry) Register(t vapkg.ProviderType, inst ProviderFactoryFn) {
	r.providers[t] = inst
}

func (r *ProviderRegistry) Get(key string, data *vapkg.Provider) IProvider {
	if v, ok := r.providers[data.GetType()]; ok {
		return v(key, data)
	}

	return nil
}

func (r *ProviderRegistry) Exists(t vapkg.ProviderType) bool {
	_, ok := r.providers[t]
	return ok
}
