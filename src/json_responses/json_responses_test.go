package jsonresponses

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONResponse(t *testing.T) {
	w := httptest.NewRecorder()

	data := struct {
		Message string `json:"message"`
	}{
		Message: "Test OK!",
	}

	JSONResponse(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, but got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', but got '%s'", contentType)
	}

	var response struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body json: %v", err)
	}

	if response.Message != data.Message {
		t.Errorf("Expected response message to be %s, got %s", data.Message, response.Message)
	}

}

func TestJSONError(t *testing.T) {
	w := httptest.NewRecorder()

	errMessage := "An error corred"
	err := errors.New(errMessage)

	JSONError(w, http.StatusNotFound, err)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status NotFound, but got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', but got '%s'", contentType)
	}

	var errorReponse struct {
		Error string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&errorReponse); err != nil {
		t.Errorf("Error decoding response body json: %v", err)
	}

	if errorReponse.Error != errMessage {
		t.Errorf("Expected error message to be %s, got %s", errMessage, errorReponse.Error)
	}
}
