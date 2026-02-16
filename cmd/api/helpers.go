package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Define an envelope type
type envelope map[string]interface{}

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

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//encode json
	js, err := json.MarshalIndent(data, "", "\t")
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		//during the procces of decoding there may be an error
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("SyntaxError : body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contain incorrect JSON type for	field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contain incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		//there is notihing to read at all
		case errors.Is(err, io.EOF):
			return errors.New("Body must not be empty")

		//when decode the non-nil pointer value as target destination
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}
