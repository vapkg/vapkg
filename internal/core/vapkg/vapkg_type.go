package vapkg

import (
	"strings"
)

type VaPackageType uint8

const (
	VaPackageTypeUnknown VaPackageType = iota
	VaPackageTypeServer
	VaPackageTypeComponent
)

var vaTypes = map[VaPackageType]string{
	VaPackageTypeUnknown:   "",
	VaPackageTypeServer:    "server",
	VaPackageTypeComponent: "component",
}

func (v VaPackageType) String() string {
	return vaTypes[v]
}

var t = VaPackageTypeUnknown

func isAcronym(t, buf string) bool {
	pos := strings.Index(t, buf)

	return pos == 0 || (pos > 0 && t[pos-1] == ';')
}

func ParseVapkgType(s string) VaPackageType {
	aliases := map[VaPackageType]string{
		VaPackageTypeUnknown:   "",
		VaPackageTypeServer:    "srv;server",
		VaPackageTypeComponent: "cmp;component",
	}

	for k, v := range aliases {
		if isAcronym(v, s) {
			t = k
			break
		}
	}

	return t
}

func GetType() VaPackageType {
	return t
}
