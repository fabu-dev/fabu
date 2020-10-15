package api

type Error struct {
	Type    string
	Code    int
	Message string
}

func NewError(Code int, Message string) *Error {
	return &Error{
		Code:    Code,
		Message: Message,
	}
}
