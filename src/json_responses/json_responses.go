package jsonresponses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func JSONError(w http.ResponseWriter, statusCode int, err error) {
	errorReponse := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	JSONResponse(w, statusCode, errorReponse)
}
