package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	//in order to be valid json string,we need to quote the jsonValue
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}
