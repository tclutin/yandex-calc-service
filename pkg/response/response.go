package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, httpStatus int, message string) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(Error{Error: message})
}
