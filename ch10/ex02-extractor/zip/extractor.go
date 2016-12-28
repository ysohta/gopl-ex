package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ysohta/gopl-ex/ch10/ex02-extractor"
)

func init() {
	extractor.Register("zip", Unzip)
}

func Unzip(dst, src string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if err := copyFile(dst, f); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(dst string, f *zip.File) (err error) {
	path := filepath.Join(dst, f.Name)
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm); err != nil {
		return fmt.Errorf("mkdir failure:%s", err)
	}

	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("open failure:%s", err)
	}

	defer rc.Close()

	if !f.FileInfo().IsDir() {
		copied, err := os.Create(path)
		if err != nil {
			return err
		}
		defer copied.Close()

		if _, err := io.Copy(copied, rc); err != nil {
			return fmt.Errorf("copy failure:%s", err)
		}
	}

	return nil
}
