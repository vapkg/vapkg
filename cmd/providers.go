package main

import (
	"vapkg/internal/core"
	"vapkg/internal/providers"
)

var ProviderMap = map[core.ProviderType]core.IHttpProvider{
	core.HttpGitProvider: providers.NewGitHttpProvider(),
}
