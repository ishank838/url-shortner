package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortner/model"
	"url-shortner/service/url"

	validator "gopkg.in/go-playground/validator.v9"
)

func ShortenUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var urlReq model.ShortenRequest

		err := json.NewDecoder(r.Body).Decode(&urlReq)

		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid request"))
			return
		}

		v := validator.New()
		err = v.Struct(urlReq)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid data: %v", err))
			return
		}

		resp, err := url.ShortenUrl(urlReq)

		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		RespondWithJson(w, resp)
	}
}

func ResolveUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("resolve url called"))
	}
}
