package controller

import (
	"net/http"

	"caio.com/test-user/repository"
	jsonresponses "caio.com/test-user/src/json_responses"
)

func GetAllCostumers(w http.ResponseWriter, r *http.Request) {
	customers, err := repository.GetAllCostumers()

	if err != nil {
		jsonresponses.JSONResponse(w, http.StatusInternalServerError, err)
	}
	jsonresponses.JSONResponse(w, http.StatusOK, customers)

}
