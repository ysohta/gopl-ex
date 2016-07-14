// Package strings implements functions to handle string.
package strings

import (
	"bytes"
)

const (
	interval  = 3
	separator = ','
)

func comma(s string) string {
	var buf bytes.Buffer
	start := len(s) % interval
	buf.WriteString(s[0:start])
	for i := start; i < len(s); i += interval {
		// avoid comma at top
		if buf.Len() > 0 {
			buf.WriteByte(separator)
		}
		buf.WriteString(s[i : i+interval])
	}
	return buf.String()
}
