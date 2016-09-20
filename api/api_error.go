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
	json.NewEncoder(*w).Encode(err)
}
