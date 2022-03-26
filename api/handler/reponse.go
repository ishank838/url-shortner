package handler

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), status)
}

func RespondWithJson(w http.ResponseWriter, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
