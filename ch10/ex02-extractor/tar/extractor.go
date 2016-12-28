package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ysohta/gopl-ex/ch10/ex02-extractor"
)

func init() {
	extractor.Register("tar", UnTar)
}

func UnTar(dst, src string) error {
	var file *os.File
	var err error

	if file, err = os.Open(src); err != nil {
		return fmt.Errorf("tar: %s", err)
	}
	defer file.Close()

	r := tar.NewReader(file)
	var header *tar.Header
	for {
		header, err = r.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("read failure: %s", err)
		}

		err = copyFile(dst, header, r)
		if err != nil {
			return err
		}
	}
}

func copyFile(dst string, h *tar.Header, r *tar.Reader) (err error) {
	path := filepath.Join(dst, h.Name)
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm); err != nil {
		return fmt.Errorf("mkdir failure:%s", err)
	}

	if !h.FileInfo().IsDir() {
		copied, err := os.Create(path)
		if err != nil {
			return err
		}
		defer copied.Close()

		if _, err := io.Copy(copied, r); err != nil {
			return fmt.Errorf("copy failure:%s", err)
		}
	}

	return nil
}
