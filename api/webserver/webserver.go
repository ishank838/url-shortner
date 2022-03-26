package webserver

import (
	"fmt"
	"log"
	"net/http"
	"url-shortner/config"
	"url-shortner/webserver/router"
	"url-shortner/webserver/routes"

	"github.com/gorilla/mux"
)

type WebServer struct {
	Port   string
	Router *mux.Router
	Routes []routes.Routes
}

func (server *WebServer) Serve() error {
	address := ":" + server.Port

	log.Println("serving at port: ", server.Port)

	err := http.ListenAndServe(address, server.Router)

	if err != nil {
		return fmt.Errorf("error at serve: %v", err)
	}
	return nil
}

func InitWebServer(webConfig *config.WebConfig) *WebServer {

	routes := routes.GetAllRoutes()

	router := router.NewRouter(routes)

	return &WebServer{
		Port:   webConfig.Port,
		Router: router,
		Routes: routes,
	}

}
