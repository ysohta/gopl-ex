// Package strings implements functions to handle string.
package strings

import (
	"bytes"
	"strings"
)

const (
	interval  = 3
	separator = ','
	decimal   = '.'
)

func comma(s string) string {
	var buf bytes.Buffer
	var start, end int

	if end = strings.IndexByte(s, decimal); end < 0 {
		// decimal mark not found
		end = len(s)
	}
	start = end % interval

	buf.WriteString(s[:start])
	for i := start; i < end; i += interval {
		// avoid comma at top
		if buf.Len() > 0 {
			buf.WriteByte(separator)
		}
		buf.WriteString(s[i : i+interval])
	}
	buf.WriteString(s[end:])
	return buf.String()
}
