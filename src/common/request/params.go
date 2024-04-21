package request

import (
	"api/src/common/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetId retrieves the ID from the request parameters.
func GetId(r *http.Request, p string) (uint64, *errors.Error) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params[p], 10, 64)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusBadRequest)
	}
	return id, nil
}
