package core

import (
	"encoding/json"
	"os"
)

type ProviderType string

const (
	GitProvider ProviderType = "git"
	PtrProvider ProviderType = "ptr"
)

// About author struct
type VaPackageAuthor struct {

	// any name
	Name string `json:"name"`

	// url (optional)
	Url string `json:"url,omitempty"`

	// email (optional)
	Email string `json:"email,omitempty"`
}

// About package providers struct
type VaPackageProvider struct {

	// supported types: git ( so far so -_-)
	Type string `json:"type"`

	// name
	Name string `json:"name"`

	// entry point
	Url string `json:"url"`
}

// VaPackageDependence About package deps struc
type VaPackageDependence struct {

	// provider uniq key (must be present as VaPackageProvider)
	Provider string `json:"provider"`

	// repository (just part of the path to access the source code project)
	Repository string `json:"repository"`

	// tag of revision (e.g. version)
	Tag string `json:"version"`

	// the name of the file to upload (optional)
	// It is used if you attach packages to the release
	// if not specified, the source code will be uploaded
	Attachment string `json:"attachment,omitempty"`

	// optional or not (optional :/)
	Optional bool `json:"optional,omitempty"`
}

type VaPackage struct {
	Name         string                       `json:"name"`
	Version      string                       `json:"version"`
	Url          string                       `json:"url,omitempty"`
	Description  string                       `json:"description,omitempty"`
	License      string                       `json:"license,omitempty"`
	Authors      []VaPackageAuthor            `json:"authors,omitempty"`
	Exports      []string                     `json:"exports"`
	Providers    map[string]VaPackageProvider `json:"repositories,omitempty"`
	Dependencies []VaPackageDependence        `json:"dependencies,omitempty"`
}

var vaPackagePattern = VaPackage{
	Name:        "awesome-proj",
	Description: "my awesome project",
	Version:     "0.0.1",
	Url:         "https://github.com/rejchev/vapkg",
	Authors: []VaPackageAuthor{
		{
			Name: "rejchev",
			Url:  "https://t.me/hevrej",
		},
	},
	Exports: []string{},
	Providers: map[string]VaPackageProvider{
		"github.com/rejchev": {
			Name: "Rejchev's GitHub",
			Type: "git",
			Url:  "https://github.com/rejchev",
		},
	},

	Dependencies: []VaPackageDependence{
		{
			Provider:   "github.com/rejchev",
			Repository: "vapkg",
			Tag:        "v0.0.1",
			Attachment: "vapkg-v001.tar.gz",
			Optional:   true,
		},
	},
}

func VaPackageExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func VaPackageInit(path string, name string) error {

	pkg := vaPackagePattern

	if len(name) > 0 {
		pkg.Name = name
	}

	return VaPackageWrite(path, &pkg)
}

func VaPackageWrite(path string, pkg *VaPackage) error {
	var err error
	var buf []byte
	if buf, err = json.MarshalIndent(pkg, "", "\t"); err != nil {
		return err
	}

	var file *os.File
	if file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666); err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.Write(buf)

	return err
}

func VaPackageRead(path string) (*VaPackage, error) {
	var err error
	var file *os.File
	if file, err = os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var pkg *VaPackage
	if err = json.NewDecoder(file).Decode(&pkg); err != nil {
		return nil, err
	}

	return pkg, nil
}
