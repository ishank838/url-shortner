package routes

import (
	"net/http"
	"url-shortner/handler"
)

type Routes struct {
	Name    string
	Method  string
	Url     string
	Handler http.HandlerFunc
}

func GetAllRoutes() []Routes {
	routes := []Routes{
		{
			Name:    "ShortenUrl",
			Method:  http.MethodPost,
			Url:     "/shorten",
			Handler: handler.ShortenUrl(),
		},
		{
			Name:    "Resolve Url",
			Method:  http.MethodGet,
			Url:     "/{shorten_key}",
			Handler: handler.ResolveUrl(),
		},
	}

	return routes
}
