package vapkg

type LockEntrySignature string

type LockEntryDependency struct {
	Dependency string                      `json:"dependency"`
	Exported   []LockEntryDependencyExport `json:"exported"`
}

type LockEntryDependencyExport struct {
	Hash string `json:"hash"`
	Path string `json:"path"`
}

type LockEntry struct {
	Name         string                         `json:"name"`
	Version      string                         `json:"version"`
	License      *string                        `json:"license,omitempty"`
	Source       string                         `json:"source"`
	Origin       *string                        `json:"origin,omitempty"`
	Hash         string                         `json:"hash"`
	Provider     string                         `json:"provider"`
	Dependencies map[string]LockEntryDependency `json:"dependencies,omitempty"`
}

type Lock struct {
	Name         string                           `json:"name"`
	Version      string                           `json:"version"`
	Type         string                           `json:"type"`
	CacheVersion string                           `json:"cache_version"`
	Packages     map[LockEntrySignature]LockEntry `json:"packages,omitempty"`
}

func (l Lock) GetName() string {
	return l.Name
}

func (l Lock) GetVersion() string {
	return l.Version
}

func (l Lock) GetType() string {
	return l.Type
}

func (l Lock) GetCacheVersion() string {
	return l.CacheVersion
}

func (l Lock) GetPackages() map[LockEntrySignature]LockEntry {
	return l.Packages
}

func (l LockEntry) GetName() string {
	return l.Name
}

func (l LockEntry) GetVersion() string {
	return l.Version
}

func (l LockEntry) GetLicense() string {
	if l.License != nil {
		return *l.License
	}
	return ""
}

func (l LockEntry) GetSource() string {
	return l.Source
}

func (l LockEntry) GetOrigin() string {
	if l.Origin != nil {
		return *l.Origin
	}

	return ""
}

func (l LockEntry) GetHash() string {
	return l.Hash
}

func (l LockEntry) GetProvider() string {
	return l.Provider
}

func (l LockEntry) GetDependencies() map[string]LockEntryDependency {
	return l.Dependencies
}

func (l LockEntryDependency) GetDependency() string {
	return l.Dependency
}

func (l LockEntryDependency) GetExported() []LockEntryDependencyExport {
	return l.Exported
}

func (el LockEntryDependencyExport) GetHash() string {
	return el.Hash
}

func (el LockEntryDependencyExport) GetPath() string {
	return el.Path
}
