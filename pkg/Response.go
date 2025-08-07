package responsePkg

import (
	"encoding/json"
	"net/http"
)

type ResponseLib struct{}

func MakeJsonResponse(w http.ResponseWriter, responseText any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseText)
}
