package lib

import "encoding/json"

func MustJSON(j map[string]interface{}) []byte {
	if j == nil {
		return []byte{}
	}

	b, _ := json.Marshal(j)

	return b
}
