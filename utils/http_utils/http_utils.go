package http_utils

import (
	"encoding/json"
	restError "items_api/utils/errors"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, resErr restError.RestError) {
	ResponseJson(w, resErr.Status, resErr)
}