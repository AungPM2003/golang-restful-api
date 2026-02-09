package main

import (
	"encoding/json"
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

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	//encode json
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//append a new line in json  just for view
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
