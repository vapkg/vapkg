package workspace

import (
	"fmt"
	"vapkg/internal/core/vapkg"
	"vapkg/internal/utils"
)

type IWorkspaceManager interface {
	CreateWorkspace(name string, t vapkg.VaPackageType) error
	DeleteWorkspace()
	GetWorkspace() IWorkspace
	LoadWorkspace() error
	ReloadWorkspace() error
	Reset()
	IsExist() bool
}

var _ IWorkspaceManager = (*Manager)(nil)

type Manager struct {
	ws *Workspace
}

func NewManager() *Manager {
	return &Manager{nil}
}

func (m *Manager) CreateWorkspace(name string, t vapkg.VaPackageType) error {
	if m.IsExist() {
		return fmt.Errorf("workspace %s already exists", name)
	}

	switch ws, err := NewWorkspace(name, t); {
	case err != nil:
		return err

	default:
		m.ws = ws
		return nil
	}
}

func (m *Manager) LoadWorkspace() error {
	if m != nil && m.GetWorkspace() != nil {
		return fmt.Errorf("workspace already loaded")
	}

	return m.ReloadWorkspace()
}

func (m *Manager) ReloadWorkspace() error {
	if !m.IsExist() {
		return fmt.Errorf("nothing to reload, workspace does not exist")
	}

	switch ws, err := NewWorkspaceFromExisting(); {
	case err != nil:
		return err

	default:
		m.ws = ws
		return nil
	}
}

func (m *Manager) DeleteWorkspace() {

}

func (m *Manager) Reset() {
	m.ws = nil
}

func (m *Manager) GetWorkspace() IWorkspace {
	if m.ws == nil {
		return nil
	}

	return m.ws
}

func (m *Manager) IsExist() bool {
	return utils.Exists(vapkg.PackageFile) && utils.IsFile(vapkg.PackageFile)
}
