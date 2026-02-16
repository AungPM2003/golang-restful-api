package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// this is the error format if unmarshal cannot parse
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	//in order to be valid json string,we need to quote the jsonValue
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}

// we have to do it with pointer reciever
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	//we expected this json value to be string.Eg - "102 mins"
	//dequote first
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
