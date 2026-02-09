package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(w http.ResponseWriter, r *http.Request) (int64, error) {
	//getting params from the context
	params := httprouter.ParamsFromContext(r.Context())

	//we get params as string so we need to convert into int64 base 10
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid Id")
	}
	return id, nil

}
