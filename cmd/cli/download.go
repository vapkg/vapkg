package cli

import (
	"fmt"
	"os"
	"path"
	"time"
	"vapkg/internal/core"
	"vapkg/internal/utils"
)

var spinner = utils.NewSpinnerPrinter([]string{"| ", "/ ", "- ", "\\ "})

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

	var repo, ver, attachment = "", "", getValue("attachment", opts)
	if repo, ver = core.GetVaPackageDepFromShorten(getValue("", opts)); repo == "" {
		return fmt.Errorf("expected dependency shorten format: <dep>@<ver> | <dep>")
	}

	dep := &core.VaPackageDependence{Repository: repo, Tag: ver, Attachment: attachment}

	var pvKey string
	if pvKey = getValue("provider", opts); pvKey == "" {
		return fmt.Errorf("--provider is required")
	}

	var err error
	var provider core.IProvider
	if provider, err = getProvider(ctx, pvKey); err != nil {
		return fmt.Errorf("%s", err)
	}

	ctx.Logger().Infof("Download command called with: packet='%s@%s'; attachment='%s'; prv='%s'", repo, ver, attachment, pvKey)

	out := getCachePath(ctx, pvKey, dep, provider.GetFile(dep))

	ctx.Logger().Infof("Download command called with: '%s'; attachment='%s'", out, attachment)

	return downloadWithSpinner(out, provider, dep)
}

func downloadWithSpinner(dest string, provider core.IProvider, dep *core.VaPackageDependence) error {
	spinner.Start(uint32(time.Second), utils.VaSprintf("downloading {FGR}%s{R}", core.GetVaPackageDepShorten(dep)))
	err := utils.DownloadFile(provider.GetPath(dep), dest)
	duration := spinner.Stop()

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	_, _ = utils.VaPrintfWithPrefix("package {FGR}%s {R}as {FGR}%s {R}download completed in {FGR}%s{R}",
		provider.GetFile(dep), core.GetVaPackageDepShorten(dep), duration)

	return err
}

func getValue(key string, opts map[string]string) string {
	if val, ok := opts[key]; ok {
		return val
	}

	return ""
}

func getCachePath(ctx *core.Context, vaKey string, dep *core.VaPackageDependence, file string) string {
	dirs := path.Join(ctx.Config().CacheFolder(), vaKey, dep.Repository, dep.Tag)
	_ = os.MkdirAll(dirs, os.ModePerm)
	return path.Join(dirs, file)
}

func getProvider(ctx *core.Context, key string) (core.IProvider, error) {
	if vaProvider, ok := ctx.Ws().Pkg().Providers[key]; ok {

		if !ctx.Providers().Exists(vaProvider.Type) {
			return nil, fmt.Errorf("provider '%s' does not support", vaProvider.Type)
		}

		return ctx.Providers().Get(key, &vaProvider), nil
	}

	return nil, fmt.Errorf("provider key '%s' must be present onto vapkg.json providers", key)
}
