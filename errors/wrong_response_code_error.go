package errors

import (
	"fmt"
	"net/http"
)

type WrongResponseCodeError struct {
	Response *http.Response
}

func (error *WrongResponseCodeError) Error() string {
	return fmt.Sprintf("Reciverd response code :  %v", error.Response.StatusCode)
}
