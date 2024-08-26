package request

import (
	"encoding/json"
	"io"
	"net/http"
	"taskmanager/internal/common/errors"
)

type entity interface{}

// ProcessBody reads the request body, unmarshals it into the provided entity, and returns an error if any.
//
// Parameters:
// - r: *http.Request - the HTTP request containing the body to process.
// - entity: entity - the entity to unmarshal the request body into.
//
// Returns:
// - error: an error if the request body could not be read or unmarshaled.
func ProcessBody(r *http.Request, entity entity) *errors.Error {
	defer r.Body.Close()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.NewError("invalid request body: "+err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(requestBody, entity); err != nil {
		return errors.NewError("invalid request body: "+err.Error(), http.StatusBadRequest)
	}

	return nil
}
