package tar

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ysohta/gopl-ex/ch10/ex02-extractor/testutils"
)

var workspace = "./work"

func TestUnTar(t *testing.T) {
	tests := []struct {
		dst  string
		src  string
		want string
		err  error
	}{
		{
			workspace,
			"../testdata/sample.tar",
			"../testdata/sample",
			nil,
		},
		{
			workspace,
			"../testdata/file.txt",
			"",
			fmt.Errorf("read failure: unexpected EOF"),
		},
	}

	for _, test := range tests {
		defer os.RemoveAll(workspace)

		err := UnTar(test.dst, test.src)
		if err != nil {
			if fmt.Sprint(err) != fmt.Sprint(test.err) {
				t.Errorf("want:%v got:%v", test.err, err)
			}
			continue
		}

		base := filepath.Base(test.want)
		ext := filepath.Ext(base)
		basename := base[:len(base)-len(ext)] // without ext

		path := filepath.Join(test.dst, basename)
		testutils.CompareFiles(test.want, path, t)
	}
}
