package router

import (
	"log"
	"url-shortner/webserver/routes"

	"github.com/gorilla/mux"
)

func NewRouter(routes []routes.Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	setRoutes(router, routes)
	return router
}

func setRoutes(router *mux.Router, routes []routes.Routes) {

	for _, v := range routes {
		router.Methods(v.Method).Path(v.Url).Name(v.Name).HandlerFunc(v.Handler)

		log.Println("Route setup", v.Name, v.Method, v.Url)
	}
}
