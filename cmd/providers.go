package main

import (
	"vapkg/internal/core"
	"vapkg/internal/core/vapkg"
	"vapkg/internal/providers"
)

var ProviderMap = map[vapkg.ProviderType]core.ProviderFactoryFn{
	providers.GitHttpProviderType: providers.NewGitHttpProviderA,
}
