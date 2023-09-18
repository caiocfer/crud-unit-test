package controller

import (
	"net/http"
	"strconv"

	"caio.com/test-user/repository"
	jsonresponses "caio.com/test-user/src/json_responses"
	"github.com/gorilla/mux"
)

func GetAllCostumers(w http.ResponseWriter, r *http.Request) {
	customers, err := repository.GetAllCostumers()

	if err != nil {
		jsonresponses.JSONResponse(w, http.StatusInternalServerError, err)
	}
	jsonresponses.JSONResponse(w, http.StatusOK, customers)

}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]

	id, err := strconv.Atoi(Id)
	customer, err := repository.GetCustomerByID(id)
	if err != nil {
		jsonresponses.JSONError(w, http.StatusInternalServerError, err)
	}
	jsonresponses.JSONResponse(w, http.StatusOK, customer)
}
