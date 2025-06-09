package core

const (
	Name    = "vapkg"
	License = "Apache License, Version 3.0"
	Author  = "rejchev"
	URL     = "https://github.com/rejchev/vapkg"
	Version = "0.0.1"
)

type ICore interface {
	Name() string
	Version() string
	Author() string
	URL() string
	License() string
}

type Core struct {
	name    string
	version string
	author  string
	url     string
	license string
}

func NewCore() *Core {
	return &Core{
		Name,
		Version,
		Author,
		URL,
		License,
	}
}

func CreateCore() Core {
	return *NewCore()
}

func (c *Core) Name() string {
	return c.name
}

func (c *Core) Version() string {
	return c.version
}

func (c *Core) Author() string {
	return c.author
}

func (c *Core) URL() string {
	return c.url
}

func (c *Core) License() string {
	return c.license
}
