package cli

import (
	"fmt"
	"vapkg/internal/core"
)

var downloadCommand = core.Command{
	Usage:       "vapkg download [--<option>[, ...]]",
	Description: "",
	Handler:     downloadCommandHandleFn,
	Options: map[string]bool{
		"":           true,
		"provider":   true,
		"attachment": false,
	},
}

func downloadCommandHandleFn(ctx *core.Context, opts map[string]string) error {

	if !ctx.Ws().Exists() {
		return fmt.Errorf("vapkg must be inited")
	}

	var dep string
	if dep = getValue("", opts); dep == "" {
		return fmt.Errorf("...")
	}

	var attachment, pvKey = getValue("attachment", opts), ""

	if pvKey = getValue("provider", opts); pvKey == "" {
		return fmt.Errorf("--provider is required")
	}

	var provider *core.VaPackageProvider
	if provider = getProvider(ctx, pvKey); provider == nil {
		return fmt.Errorf("unknown provider '%s', it must be declared on vapkg", pvKey)
	}

	ctx.Logger().Infof("Download command called with: packet='%s'; attachment='%s'; prv='%s'", dep, attachment, pvKey)

	return nil
}

func getValue(key string, opts map[string]string) string {
	if val, ok := opts[key]; ok {
		return val
	}

	return ""
}

func getProvider(ctx *core.Context, key string) *core.VaPackageProvider {
	if provider, ok := ctx.Ws().Pkg().Providers[key]; ok {
		return &provider
	}

	return nil
}
