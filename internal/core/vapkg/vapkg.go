package vapkg

const VersionSeparator = "@"

const (
	PackageFileName      = "vapkg"
	PackageFileExtension = "json"
	PackageFile          = PackageFileName + "." + PackageFileExtension
)

const (
	PackageLockFileName      = "vapkg.lock"
	PackageLockFileExtension = "json"
	PackageLockFile          = PackageLockFileName + "." + PackageLockFileExtension
)

type VaPackage struct {
	Name         string              `json:"name"`
	Version      string              `json:"version"`
	Type         string              `json:"type"`
	URL          *string             `json:"url,omitempty"`
	Description  *string             `json:"description,omitempty"`
	License      *string             `json:"license,omitempty"`
	Authors      []Author            `json:"authors,omitempty"`
	Exports      []string            `json:"exports,omitempty"`
	Providers    map[string]Provider `json:"repositories,omitempty"`
	Dependencies []Dependency        `json:"dependencies,omitempty"`
}

func NewVaPackage() *VaPackage {
	return &VaPackage{}
}

func (v *VaPackage) GetName() string {
	return v.Name
}

func (v *VaPackage) GetVersion() string {
	return v.Version
}

func (v *VaPackage) GetType() VaPackageType {
	if GetType() == VaPackageTypeUnknown {
		ParseVapkgType(v.Type)
	}

	return GetType()
}

func (v *VaPackage) GetLicense() string {
	if v.License != nil {
		return *v.License
	}

	return ""
}

func (v *VaPackage) GetURL() string {
	if v.URL != nil {
		return *v.URL
	}

	return ""
}

func (v *VaPackage) GetDescription() string {
	if v.Description != nil {
		return *v.Description
	}
	return ""
}

func (v *VaPackage) GetAuthors() []Author {
	return v.Authors
}

func (v *VaPackage) GetExports() []string {
	return v.Exports
}

func (v *VaPackage) GetProvider(name string) *Provider {
	if p, ok := v.Providers[name]; ok {
		return &p
	}
	return nil
}

func (v *VaPackage) GetDependencies() []Dependency {
	return v.Dependencies
}

func (v *VaPackage) AddDependency(dependency Dependency) bool {
	if d := v.FindDependency(dependency.GetAttachment()); d != -1 {
		return false
	}

	v.Dependencies = append(v.Dependencies, dependency)
	return true
}

func (v *VaPackage) FindDependency(attachment string) int {
	for idx, d := range v.Dependencies {
		if d.GetAttachment() == attachment {
			return idx
		}
	}
	return -1
}

func (v *VaPackage) RemoveDependency(attachment string) {
	if d := v.FindDependency(attachment); d != -1 {
		v.Dependencies = append(v.Dependencies[:d], v.Dependencies[d+1:]...)
	}
}

func (v *VaPackage) ToLock() Lock {
	return Lock{}
}

func (v *VaPackage) SetName(name string) *VaPackage {
	v.Name = name
	return v
}

func (v *VaPackage) SetDescription(description string) *VaPackage {
	v.Description = &description
	return v
}

func (v *VaPackage) SetVersion(version string) *VaPackage {
	v.Version = version
	return v
}

func (v *VaPackage) SetType(t VaPackageType) *VaPackage {
	v.Type = t.String()
	return v
}

func (v *VaPackage) SetURL(url string) *VaPackage {
	v.URL = &url
	return v
}

func (v *VaPackage) SetAuthor(a *Author) *VaPackage {
	if a == nil {
		return v
	}

	v.Authors = append(v.Authors, *a)
	return v
}

func (v *VaPackage) SetAuthors(a []Author) *VaPackage {
	if a == nil {
		return v
	}

	v.Authors = append(v.Authors, a...)
	return v
}

func (v *VaPackage) SetProvider(n string, provider *Provider) *VaPackage {
	if provider == nil {
		return v
	}

	if v.Providers == nil {
		v.Providers = make(map[string]Provider)
	}

	v.Providers[n] = *provider
	return v
}

func (v *VaPackage) SetDependency(d *Dependency) *VaPackage {
	if d == nil {
		return v
	}

	v.Dependencies = append(v.Dependencies, *d)
	return v
}

func (v *VaPackage) SetDependencies(d []Dependency) *VaPackage {
	if d == nil {
		return v
	}

	v.Dependencies = append(v.Dependencies, d...)
	return v
}
