package cli

import (
	"archive/zip"
	"fmt"
	"path"
	"strings"
	"time"
	"vapkg/internal/core"
	"vapkg/internal/core/vapkg"

	"vapkg/internal/utils"
)

const delimiter = "@"

var installCommand = core.Command{
	Usage:       "vapkg install [--<option>[, ...]]",
	Description: "",
	Handler:     installCommandHandleFn,
	Options: map[string]bool{
		"":         false,
		"silent":   false,
		"provider": false,
		"os":       false,
	},
}

func installCommandHandleFn(ctx core.IContext, opts map[string]string) error {

	if !ctx.Workspace().IsExist() {
		return fmt.Errorf("{FRD}vapkg isn't inited, use: {R}vapkg init\n")
	}

	if ctx.Workspace().GetWorkspace().GetType() != vapkg.VaPackageTypeServer {
		return fmt.Errorf("{FRD}vapkg type must be 'server' to install deps")
	}

	if _, ok := opts[""]; ok {
		return fmt.Errorf("")
	}

	return install(ctx, opts)
}

func install(ctx core.IContext, opts map[string]string) error {

	p := ""
	err := (error)(nil)
	for _, dep := range ctx.Workspace().GetWorkspace().GetPackage().GetDependencies() {
		if p, err = downloadDependency(ctx, &dep); err != nil {
			ctx.Logger().Infof("Downloading failed (%v)", err)
			continue
		}

		if err = tryExportZippedData(ctx, &dep, p); err != nil {
			ctx.Logger().Infof("Unzipping failed (%v)", err)
			//return fmt.Errorf("{FRD}zipping failed (%v)\n", err.Error())
			continue
		}

	}

	return nil
}

func downloadDependency(ctx core.IContext, dep *vapkg.Dependency) (string, error) {
	if provider := ctx.Workspace().GetWorkspace().GetPackage().GetProvider(dep.Provider); provider != nil {
		if ctx.Providers().Exists(provider.Type) {
			depCachedPath := path.Join(ctx.Config().PackagePath(), dep.GetProviderSignature(), dep.GetRepository(), dep.GetTag())

			switch provInstance := ctx.Providers().Get(string(provider.Type), provider); provInstance {
			case nil:
				return "", fmt.Errorf("provider '%s' does not support", provider.Type)

			default:
				depCachedFile := path.Join(depCachedPath, provInstance.GetFile(dep))
				spinner.Start(uint32(time.Second), utils.VaSprintf("{R}downloading {FGR}%s{R}", vapkg.GetVaPackageDependencyDisplay(dep)))
				err := utils.DownloadFile(provInstance.GetPath(dep), depCachedFile)
				duration := spinner.Stop()

				if err == nil {
					_, _ = utils.VaPrintfWithPrefix("pkg {FGR}%s {R}as {FGR}%s {R}download completed in {FGR}%s{R}\n",
						provInstance.GetFile(dep), vapkg.GetVaPackageDependencyDisplay(dep), duration)
				}

				return depCachedFile, err
			}
		}
	}

	return "", fmt.Errorf("provider '%s' does not present or support", dep.Provider)
}

func tryExportZippedData(ctx core.IContext, dep *vapkg.Dependency, source string) error {
	switch z, err := zip.OpenReader(source); {
	case err != nil:
		return err

	default:
		defer func(z *zip.ReadCloser) {
			_ = z.Close()
		}(z)

		idx := findVaPackageFile(z.File)

		if idx == -1 {
			return fmt.Errorf("'%s' not found in zip", vapkg.PackageFile)
		}

		fmt.Printf("file: %s", z.File[idx].Name)

		return nil
	}
}

func findVaPackageFile(files []*zip.File) int {

	for i, f := range files {
		if strings.HasSuffix(f.Name, vapkg.PackageFile) {
			return i
		}
	}

	return -1
}
