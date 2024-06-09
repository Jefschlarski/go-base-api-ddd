package responses

import (
	"api/src/application/common/errors"
	"encoding/json"
	"net/http"
)

// Json is a function that writes the provided body as JSON response with the given status code.
//
// Parameters:
// - w: http.ResponseWriter - the response writer to write the JSON response.
// - code: int - the HTTP status code to be set in the response.
// - body: interface{} - the body of the JSON response.
//
// Return type: None.
func Json(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			Error(w, errors.NewError(err.Error(), http.StatusInternalServerError))
		}
	}
}

// Error writes the provided error as JSON response with the given status code.
//
// Parameters:
// - w: http.ResponseWriter - the response writer to write the JSON response.
// - code: int - the HTTP status code to be set in the response.
// - err: error - the error to be included in the JSON response.
//
// Return type: None.
func Error(w http.ResponseWriter, err *errors.Error) {
	Json(w, err.Status, struct {
		Erro string `json:"error"`
	}{
		Erro: err.Message,
	})
}
