package handler

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	HttpStatusCode int    `json:"http_status_code"`
	HttpStatusMsg  string `json:"http_status_msg"`
	ErrorMsg       string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {

	resp := apiError{
		HttpStatusCode: status,
		HttpStatusMsg:  http.StatusText(status),
		ErrorMsg:       err.Error(),
	}

	RespondWithJson(w, resp)
}

func RespondWithJson(w http.ResponseWriter, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
