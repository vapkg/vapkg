package core

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	PackageName      = "vapkg"
	PackageExtension = "json"
)

var gameIRequired = []string{
	"gameinfo.gi",
	"gameinfo.txt",
	"gameinfo.ini",
}

type Workspace struct {
	path string
	pkg  *VaPackage
}

func CreateWorkspace(path string) Workspace {
	return *NewWorkspace(path)
}

func NewWorkspace(path string) *Workspace {
	pkgPath := pathWithPackageName(path)

	var pkg *VaPackage
	if VaPackageExists(pkgPath) {
		pkg, _ = VaPackageRead(pkgPath)
	}

	return &Workspace{path, pkg}
}

func (ws *Workspace) Init(name string) error {

	var exists bool
	for _, v := range gameIRequired {
		if _, err := os.Stat(filepath.Join(ws.path, v)); !os.IsNotExist(err) {
			exists = true
			break
		}
	}

	if !exists {
		return fmt.Errorf("required gameinfo.<ext> file isn't found")
	}

	return VaPackageInit(pathWithPackageName(ws.path), name)
}

func (ws *Workspace) Exists() bool {
	return ws.pkg != nil
}

func (ws *Workspace) Path() string {
	return ws.path
}

func pathWithPackageName(path string) string {
	return filepath.Join(path, PackageName+"."+PackageExtension)
}
