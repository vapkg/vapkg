package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	var err error
	var resp *http.Response

	if resp, err = http.Head(url); err != nil {
		return err
	}

	size := resp.ContentLength

	if resp, err = http.Get(url); err != nil {
		return err
	}

	defer func(v io.ReadCloser) {
		_ = v.Close()
	}(resp.Body)

	var out *os.File
	if out, err = os.Create(filepath); err != nil {
		return err
	}

	defer func(v *os.File) {
		_ = v.Close()
	}(out)

	text := ""
	length := 0
	buf := make([]byte, size)
	m, n := int64(0), int64(0)

	for {
		if length, err = resp.Body.Read(buf); err != nil {
			return err
		}

		n += int64(length)

		if n-m > 1024*1024 {
			m = n
			_ = fmt.Sprintf("%s (%d %%)", text, 100*n/size) // spinner
		}

		if length == 0 {
			break
		}

		if _, err = out.Write(buf[:length]); err != nil {
			return err
		}
	}

	if size > 0 && n != size {
		return fmt.Errorf("file %s has %d bytes instead of %d bytes", filepath, size, n)
	}

	return nil
}
