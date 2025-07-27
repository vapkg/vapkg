package vapkg

type Dependency struct {

	// provider uniq key (must be present as VaPackageProvider)
	Provider string `json:"provider"`

	// repository (just part of the path to access the source code project)
	Repository string `json:"repository"`

	// tag of revision (e.g. version)
	Tag string `json:"version"`

	// the name of the file to upload (optional)
	// It is used if you attach packages to the release
	// if not specified, the source code will be uploaded
	Attachment string `json:"attachment"`

	// optional or not (optional :/)
	Optional bool `json:"optional,omitempty"`
}

func NewDependency() *Dependency {
	return &Dependency{}
}

func (d *Dependency) SetProvider(provider string) *Dependency {
	d.Provider = provider
	return d
}

func (d *Dependency) SetRepository(repository string) *Dependency {
	d.Repository = repository
	return d
}

func (d *Dependency) SetTag(tag string) *Dependency {
	d.Tag = tag
	return d
}

func (d *Dependency) SetAttachment(attachment string) *Dependency {
	d.Attachment = attachment
	return d
}

func (d *Dependency) SetOptional(optional bool) *Dependency {
	d.Optional = optional
	return d
}

func (d *Dependency) GetProviderSignature() string {
	return d.Provider
}

func (d *Dependency) GetRepository() string {
	return d.Repository
}

func (d *Dependency) GetTag() string {
	return d.Tag
}

func (d *Dependency) GetAttachment() string {
	return d.Attachment
}

func (d *Dependency) IsOptional() bool {
	return d.Optional
}
