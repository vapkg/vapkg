package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(from, to string) error {
	var err error
	var resp *http.Response

	if resp, err = http.Head(from); err != nil {
		return err
	}

	size := resp.ContentLength

	if resp, err = http.Get(from); err != nil {
		return err
	}

	defer func(v io.ReadCloser) {
		_ = v.Close()
	}(resp.Body)

	var out *os.File
	if out, err = os.Create(to); err != nil {
		return err
	}

	defer func(v *os.File) {
		_ = v.Close()
	}(out)

	length := 0
	buf := make([]byte, 1024)
	n := int64(0)

	for {
		if length, err = resp.Body.Read(buf); err != nil && err != io.EOF {
			return err
		}

		n += int64(length)

		if length == 0 {
			break
		}

		if _, err = out.Write(buf[:length]); err != nil {
			return err
		}
	}

	if size > 0 && n != size {
		return fmt.Errorf("file %s has %d bytes instead of %d bytes", to, size, n)
	}

	return nil
}
