package response

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	encoder.SetIndent("", "")

	if err := encoder.Encode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_ = encoder.Encode(map[string]interface{}{"error": err.Error()})

		return
	}
}
