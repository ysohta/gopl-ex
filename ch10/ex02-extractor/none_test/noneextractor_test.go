package none_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ysohta/gopl-ex/ch10/ex02-extractor"
)

func TestExtractNoneRegistered(t *testing.T) {
	var workspace = "./work"

	tests := []struct {
		dst string
		src string
		err error
	}{
		{
			workspace,
			"../testdata/sample.zip",
			fmt.Errorf("unsupported file type:zip"),
		},
		{
			workspace,
			"../testdata/sample.tar",
			fmt.Errorf("unsupported file type:tar"),
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
