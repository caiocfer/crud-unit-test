package routes

import (
	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	router := mux.NewRouter()

	return ConfigRoutes(router)

}
