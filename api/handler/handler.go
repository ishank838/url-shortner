package handler

import "net/http"

func ShortenUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("shorten url called"))
	}
}

func ResolveUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("resolve url called"))
	}
}
