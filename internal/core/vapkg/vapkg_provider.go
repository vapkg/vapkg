package vapkg

type ProviderType string

type Provider struct {

	// supported types: git ( so far so -_-)
	Type ProviderType `json:"type"`

	// name
	Name string `json:"name"`

	// web point
	URL string `json:"url"`
}

func NewProviderE(t, name, url string) *Provider {
	return &Provider{ProviderType(t), name, url}
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) SetType(t ProviderType) *Provider {
	p.Type = t
	return p
}

func (p *Provider) SetName(name string) *Provider {
	p.Name = name
	return p
}

func (p *Provider) SetURL(url string) *Provider {
	p.URL = url
	return p
}

func (p *Provider) GetType() ProviderType {
	return p.Type
}

func (p *Provider) GetName() string {
	return p.Name
}

func (p *Provider) GetURL() string {
	return p.URL
}
