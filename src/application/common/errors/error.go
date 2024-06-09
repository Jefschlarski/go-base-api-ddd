package errors

type Error struct {
	Message string
	Status  int
}

func NewError(message string, status int) *Error {
	return &Error{
		Message: message,
		Status:  status,
	}
}
