package extractor

import (
	"fmt"
	"path/filepath"
)

type ExtractArchive func(dst, src string) error

var extractors map[string]ExtractArchive = map[string]ExtractArchive{}

func Register(suffix string, extractor ExtractArchive) {
	extractors[suffix] = extractor
}

func Extract(dst, src string) error {
	ext := filepath.Ext(src)
	suffix := ext[1:] // remove dot
	extractor, ok := extractors[suffix]
	if !ok {
		return fmt.Errorf("unsupported file type:%s", suffix)
	}
	return extractor(dst, src)
}
