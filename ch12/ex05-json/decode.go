package json

import "encoding/json"

// Unmarshal parses json data and populates the variable
// whose address is in the non-nil pointer out.
func Unmarshal(data []byte, v interface{}) (err error) {
	return json.Unmarshal(data, v)
}
