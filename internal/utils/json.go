package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func JsonFromFile(f *os.File, t any) error {
	if f == nil {
		return errors.New("file is nil")
	}

	return json.NewDecoder(f).Decode(t)
}

func JsonToFile(f *os.File, v any) error {
	if f == nil {
		return errors.New("file is nil")
	}

	if buf, err := json.MarshalIndent(v, "", "\t"); true {

		if err != nil {
			return err
		}

		_, err = f.Write(buf)
		return err
	}

	return nil
}

func JsonFromPath(path string, t any) error {
	switch file, err := OpenFile(path, os.O_RDONLY, 0666); {
	case err != nil:
		return err

	default:
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		return JsonFromFile(file, t)
	}
}

func JsonToPath(path string, v any) error {
	switch file, err := OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666); {
	case err != nil:
		return err

	default:
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		return JsonToFile(file, v)
	}
}
