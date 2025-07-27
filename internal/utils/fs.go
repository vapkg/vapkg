package utils

import (
	"os"
	"path/filepath"
)

var pwd = ""

func getPwd() string {
	return pwd
}

func GetDirectory() (string, error) {
	if pwd != "" {
		return pwd, nil
	}

	switch wd, err := os.Getwd(); err != nil {
	case err != nil:
		return "", err
	default:
		pwd = wd
	}

	return pwd, nil
}

func Exists(path string) bool {
	if getPwd() == "" {
		_, _ = GetDirectory()
	}
	_, err := os.Stat(filepath.Join(getPwd(), path))
	return err == nil
}

func IsDir(path string) bool {
	if getPwd() == "" {
		_, _ = GetDirectory()
	}

	s, err := os.Stat(filepath.Join(getPwd(), path))
	return err == nil && s.IsDir()
}

func IsFile(path string) bool {
	if getPwd() == "" {
		_, _ = GetDirectory()
	}
	s, err := os.Stat(filepath.Join(getPwd(), path))
	return err == nil && !s.IsDir()
}

func CreateDir(path string) error {
	if getPwd() == "" {
		_, _ = GetDirectory()
	}

	return os.Mkdir(filepath.Join(getPwd(), path), 0755)
}

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if getPwd() == "" {
		_, _ = GetDirectory()
	}

	return os.OpenFile(filepath.Join(getPwd(), name), flag, perm)
}

func Open(name string) (*os.File, error) {
	if getPwd() == "" {
		_, _ = GetDirectory()
	}
	return os.Open(filepath.Join(getPwd(), name))
}
