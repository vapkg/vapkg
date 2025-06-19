package providers

import (
	"fmt"
	"vapkg/internal/core"
)

const gitAttachmentPathPattern = "/%s/releases/download/%s/%s"
const gitSourcePathPattern = "/%s/archive/refs/tags/%s/source.zip"

type GitHttpProvider struct {
}

func NewGitHttpProvider() core.IHttpProvider {
	return &GitHttpProvider{}
}

func (p *GitHttpProvider) Type() core.ProviderType {
	return core.HttpGitProvider
}

func (p *GitHttpProvider) GetPath(pkg *core.VaPackageDependence) (string, error) {

	if pkg.Attachment == "" {
		return fmt.Sprintf(gitSourcePathPattern, pkg.Repository, pkg.Tag), nil
	}

	return fmt.Sprintf(gitAttachmentPathPattern, pkg.Repository, pkg.Tag, pkg.Attachment), nil
}
