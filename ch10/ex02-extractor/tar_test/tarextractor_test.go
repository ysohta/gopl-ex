package tar_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ysohta/gopl-ex/ch10/ex02-extractor"
	_ "github.com/ysohta/gopl-ex/ch10/ex02-extractor/tar"
	_ "github.com/ysohta/gopl-ex/ch10/ex02-extractor/zip"
)

func TestExtractMultiRegistered(t *testing.T) {
	var workspace = "./work"

	tests := []struct {
		dst string
		src string
		err error
	}{
		{
			workspace,
			"../testdata/sample.zip",
			nil,
		},
		{
			workspace,
			"../testdata/sample.tar",
			nil,
		},
	}

	for _, test := range tests {
		defer os.RemoveAll(workspace)

		err := extractor.Extract(test.dst, test.src)
		if fmt.Sprint(err) != fmt.Sprint(test.err) {
			t.Errorf("want:%v got:%v", test.err, err)
		}
	}
}
