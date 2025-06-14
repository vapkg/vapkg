package providers

import (
	"fmt"
	"os"
	"vapkg/internal/core"
)

// it must be something of https://<provider_domen>/<provider_uniq>/<pathPattern>
// <pathPattern>: /<repository>/v/<tag>/<attachment>/download
const pathPattern = "/%s/v/%s/%s/download"

func createPath(repo, tag, attachment string) string {
	return fmt.Sprintf(pathPattern, repo, tag, attachment)
}

func DownloadViaPtr(ctx *core.Context, idx int, outp string) *os.File {
	if idx < 0 || !ctx.Ws().Exists() || len(ctx.Ws().Pkg().Dependencies) <= idx {
		return nil
	}

	dep := ctx.Ws().Pkg().Dependencies[idx]

	return nil
}
