package providers

import (
	"fmt"
	"vapkg/internal/core"
)

const GitHttpProviderType core.ProviderType = "http-git"

const gitAttachmentPathPattern = "%s/%s/releases/download/%s/%s"
const gitSourcePathPattern = "%s/%s/archive/refs/tags/%s.zip"

type GitHttpProvider struct {
	name string
	data *core.VaPackageProvider
}

func NewGitHttpProvider(name string, data *core.VaPackageProvider) *GitHttpProvider {
	return &GitHttpProvider{
		name: name,
		data: data,
	}
}

func NewGitHttpProviderA(name string, data *core.VaPackageProvider) core.IProvider {
	return NewGitHttpProvider(name, data)
}

func (p *GitHttpProvider) GetType() core.ProviderType {
	return GitHttpProviderType
}

func (p *GitHttpProvider) GetName() string {
	return p.name
}

func (p *GitHttpProvider) GetPath(pkg *core.VaPackageDependence) string {

	if pkg.Attachment == "" {
		return fmt.Sprintf(gitSourcePathPattern, p.data.Url, pkg.Repository, pkg.Tag)
	}

	return fmt.Sprintf(gitAttachmentPathPattern, p.data.Url, pkg.Repository, pkg.Tag, pkg.Attachment)
}

func (p *GitHttpProvider) GetFile(dep *core.VaPackageDependence) string {
	if dep.Attachment == "" {
		return dep.Tag + ".zip"
	}

	return dep.Attachment
}
