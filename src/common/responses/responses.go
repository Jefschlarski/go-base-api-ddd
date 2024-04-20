package responses

import (
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
			Error(w, http.StatusInternalServerError, err)
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
func Error(w http.ResponseWriter, code int, err error) {
	Json(w, code, struct {
		Erro string `json:"error"`
	}{
		Erro: err.Error(),
	})
}
