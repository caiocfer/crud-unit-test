package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Function func(http.ResponseWriter, *http.Request)
	Method   string
}

func ConfigRoutes(router *mux.Router) *mux.Router {
	routes := customerRoute

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router

}
