package workspace

import (
	"fmt"
	"vapkg/internal/core/vapkg"
	"vapkg/internal/utils"
)

type IWorkspace interface {
	GetType() vapkg.VaPackageType
	GetPackage() *vapkg.VaPackage
}

type Workspace struct {
	pkg     *vapkg.VaPackage
	pkgLock *vapkg.Lock
}

var _ IWorkspace = (*Workspace)(nil)

func NewWorkspace(name string, t vapkg.VaPackageType) (*Workspace, error) {
	if t == vapkg.VaPackageTypeUnknown {
		t = vapkg.VaPackageTypeServer
	}

	switch data := vapkg.GetVaPackagePattern(name, t); data {
	case nil:
		return nil, fmt.Errorf("invalid vapkg type %s", t)
	default:
		if err := utils.JsonToPath(vapkg.PackageFile, data); err != nil {
			return nil, err
		}

		return &Workspace{
			pkg: data,
		}, nil
	}
}

func NewWorkspaceFromExisting() (*Workspace, error) {

	if !utils.Exists(vapkg.PackageFile) || !utils.IsFile(vapkg.PackageFile) {
		return nil, fmt.Errorf("%s file not found", vapkg.PackageFile)
	}

	pkg := new(vapkg.VaPackage)

	if err := utils.JsonFromPath(vapkg.PackageFile, pkg); err != nil {
		return nil, err
	}

	return &Workspace{pkg: pkg}, nil
}

func (w *Workspace) GetPackage() *vapkg.VaPackage {
	return w.pkg
}

func (w *Workspace) GetType() vapkg.VaPackageType {
	if w.pkg == nil {
		return vapkg.VaPackageTypeUnknown
	}

	if vapkg.GetType() == vapkg.VaPackageTypeUnknown {
		vapkg.ParseVapkgType(w.pkg.Type)
	}

	return vapkg.GetType()
}

func (w *Workspace) Dump() error {
	return nil
}
