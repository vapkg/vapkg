package providers

import (
	"fmt"
	"vapkg/internal/core"
	"vapkg/internal/core/vapkg"
)

const GitHttpProviderType vapkg.ProviderType = "http-git"

const gitAttachmentPathPattern = "%s/%s/releases/download/%s/%s"
const gitSourcePathPattern = "%s/%s/archive/refs/tags/%s.zip"

type GitHttpProvider struct {
	name string
	data *vapkg.Provider
}

func NewGitHttpProvider(name string, data *vapkg.Provider) *GitHttpProvider {
	return &GitHttpProvider{
		name: name,
		data: data,
	}
}

func NewGitHttpProviderA(name string, data *vapkg.Provider) core.IProvider {
	return NewGitHttpProvider(name, data)
}

func (p *GitHttpProvider) GetType() vapkg.ProviderType {
	return GitHttpProviderType
}

func (p *GitHttpProvider) GetName() string {
	return p.name
}

func (p *GitHttpProvider) GetPath(pkg *vapkg.Dependency) string {

	if pkg.GetAttachment() == "" {
		return fmt.Sprintf(gitSourcePathPattern, p.data.GetURL(), pkg.GetRepository(), pkg.GetTag())
	}

	return fmt.Sprintf(gitAttachmentPathPattern, p.data.GetURL(), pkg.GetRepository(), pkg.GetTag(), pkg.GetAttachment())
}

func (p *GitHttpProvider) GetFile(dep *vapkg.Dependency) string {
	if dep.GetAttachment() == "" {
		return dep.GetTag() + ".zip"
	}

	return dep.GetAttachment()
}
