package sopa

type ErrorType int

const (
	// ErrUnableToParse will be returned when the HTML could not be parsed
	ErrUnableToParse ErrorType = iota
	// ErrElementNotFound will be returned when element was not found
	ErrElementNotFound
	// ErrNoNextSibling will be returned when no next sibling can be found
	ErrNoNextSibling
	// ErrNoPreviousSibling will be returned when no previous sibling can be found
	ErrNoPreviousSibling
	// ErrNoNextElementSibling will be returned when no next element sibling can be found
	ErrNoNextElementSibling
	// ErrNoPreviousElementSibling will be returned when no previous element sibling can be found
	ErrNoPreviousElementSibling
	// ErrCreatingGetRequest will be returned when the get request couldn't be created
	ErrCreatingGetRequest
	// ErrInGetRequest will be returned when there was an error during the get request
	ErrInGetRequest
	// ErrCreatingPostRequest will be returned when the post request couldn't be created
	ErrCreatingPostRequest
	// ErrMarshallingPostRequest will be returned when the body of a post request couldn't be serialized
	ErrMarshallingPostRequest
	// ErrReadingResponse will be returned if there was an error reading the response to our get request
	ErrReadingResponse
)

type Error struct {
	Type ErrorType
	msg  string
}

func (se Error) Error() string {
	return se.msg
}

func newError(t ErrorType, msg string) Error {
	return Error{Type: t, msg: msg}
}
