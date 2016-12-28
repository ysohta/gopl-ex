package testutils

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func CompareFiles(want, got string, t *testing.T) {
	f1 := FileTree(want)
	f2 := FileTree(got)

	if !reflect.DeepEqual(f1, f2) {
		t.Errorf("want:%v got:%v", f1, f2)
	}
}

func FileTree(root string) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		p, _ := filepath.Rel(root, path)
		files = append(files, p)
		return nil
	})
	return files
}
