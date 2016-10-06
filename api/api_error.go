package api

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Msg        string `json:"message"`
	Httpstatus int    `json:"code"`
}

func (err *APIError) NewApiError(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(err)
}
