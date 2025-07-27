package vapkg

import (
	"strings"
)

func GetVaPackagePattern(name string, t VaPackageType) *VaPackage {
	if t == VaPackageTypeUnknown {
		return nil
	}

	if name == "" {
		name = "my-pkg"
	}

	return NewVaPackage().
		SetName(name).
		SetVersion("0.0.1").
		SetDescription("My awesome package description").
		SetType(t).
		SetURL("https://github.com/rejchev/vapkg").
		SetAuthor(NewAuthor("rejchev").SetURL("https://t.me/hevrej")).
		SetProvider("github.com/rejchev", NewProviderE("http-git", "Rejchev's GitHub", "https://github.com/rejchev")).
		SetDependency(NewDependency().
			SetProvider("github.com/rejchev").
			SetRepository("vapkg").
			SetTag("v0.0.1").
			SetAttachment("jansson-1.7.3.125-windows-latest.zip").
			SetOptional(true))
}

func BreakShortenVaPackage(v string) (string, string) {
	if idx := strings.Index(v, VersionSeparator); idx > 0 {

		return v[:idx], v[idx:]
	}
	return "", ""
}

func GetVaPackageDependencyDisplay(d *Dependency) string {
	return d.Repository + "@" + d.Tag
}
