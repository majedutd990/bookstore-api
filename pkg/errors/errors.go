package errors

import (
	"net/http"
)

// our kind here is private because it is lower case
type (
	kind        uint
	serverError struct {
		kind    kind
		message string
	}
)

// _means we do not need 0 value of iota
const (
	_ kind = iota
	KindInvalid
	KindNotFound
	KindUnAuthorized
	KindUnexpected
	KindNotAllowed
)

// these above constants are the only way to use errors
//  cause we use them in below var

// mapping these kind values to their error's message
// using a map var like below
var (
	httpErrors = map[kind]int{
		KindInvalid:      http.StatusBadRequest,
		KindNotFound:     http.StatusNotFound,
		KindUnAuthorized: http.StatusUnauthorized,
		KindUnexpected:   http.StatusInternalServerError,
		KindNotAllowed:   http.StatusMethodNotAllowed,
	}
)

func New(kind kind, msg string) error {
	return serverError{
		kind:    kind,
		message: msg,
	}
}
func (e serverError) Error() string {
	return e.message
}

//  now we need a function that takes a server error and returns an http err
func HttpError(e error) (string, int) {
	// the below structure e.(serverError) checks type of e if it is over serverError ok will be true
	serverErr, ok := e.(serverError)
	if !ok {
		return e.Error(), http.StatusBadRequest
	}
	code, ok := httpErrors[serverErr.kind]
	if !ok {
		return serverErr.message, http.StatusBadRequest
	}
	return serverErr.message, code
}
