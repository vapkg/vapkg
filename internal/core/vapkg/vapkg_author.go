package vapkg

type Author struct {

	// any name
	Name string `json:"name"`

	// url (optional)
	URL *string `json:"url,omitempty"`

	// email (optional)
	Email *string `json:"email,omitempty"`
}

func NewAuthor(name string) *Author {
	return &Author{Name: name}
}

func (a *Author) SetName(n string) *Author {
	a.Name = n
	return a
}

func (a *Author) SetURL(url string) *Author {
	a.URL = &url
	return a
}

func (a *Author) SetEmail(email string) *Author {
	a.Email = &email
	return a
}

func (a *Author) GetName() string {
	return a.Name
}

func (a *Author) GetURL() string {
	if a.URL == nil {
		return ""
	}

	return *a.URL
}

func (a *Author) GetEmail() string {
	if a.Email == nil {
		return ""
	}
	return *a.Email
}
