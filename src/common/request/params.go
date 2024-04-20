package request

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetId retrieves the ID from the request parameters.
//
// It takes an http.Request as a parameter and returns a uint64 ID and an error.
func GetId(r *http.Request) (uint64, error) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
