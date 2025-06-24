package core

type ProviderType string

type ProviderFactoryFn func(name string, data *VaPackageProvider) IProvider

type IProvider interface {
	GetName() string
	GetType() ProviderType
	GetPath(dep *VaPackageDependence) string
	GetFile(dep *VaPackageDependence) string
}

type ProviderRegistry struct {
	providers map[ProviderType]ProviderFactoryFn
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{make(map[ProviderType]ProviderFactoryFn)}
}

func CreateProviderRegistry() ProviderRegistry {
	return ProviderRegistry{make(map[ProviderType]ProviderFactoryFn)}
}

func (r *ProviderRegistry) Register(t ProviderType, inst ProviderFactoryFn) {
	r.providers[t] = inst
}

func (r *ProviderRegistry) Get(key string, data *VaPackageProvider) IProvider {
	if v, ok := r.providers[data.Type]; ok {
		return v(key, data)
	}

	return nil
}

func (r *ProviderRegistry) Exists(t ProviderType) bool {
	_, ok := r.providers[t]
	return ok
}
