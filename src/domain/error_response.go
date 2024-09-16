package domain

type (
	ErrorCode    int
	ErrorMessage string
)

const (
	InvalidInput       ErrorCode = 1001
	UnauthorizedAccess ErrorCode = 1002
	NotFound           ErrorCode = 1003
	DatabaseError      ErrorCode = 1004
	ValidateionError   ErrorCode = 1005
	InvalidError       ErrorCode = 1006
)

var ErrorMessages = map[ErrorCode]ErrorMessage{
	InvalidInput:       "Invalid input provided.",
	UnauthorizedAccess: "Unauthorized access. Please log in.",
	NotFound:           "Resource not found.",
	DatabaseError:      "Database error occurred.",
	ValidateionError:   "Validation error occured.",
	InvalidError:       "Invalid error occured.",
}

type ErrorResponse struct {
	Status string  `json:"status"`
	Errors []error `json:"errors"`
}

type MyError struct {
	StatusCode ErrorCode    `json:"status_code"`
	Message    ErrorMessage `json:"message"`
	Param      string       `json:"param,omitempty"`
	Detail     interface{}  `json:"error_detail,omitempty"`
}

func (e MyError) Error() string {
	return string(e.Message)
}

func NewMyError(code ErrorCode, param string) error {
	return MyError{StatusCode: code, Message: ErrorMessages[code], Param: param}
}

func NewErrorResponse(errors ...error) ErrorResponse {
	return ErrorResponse{Status: "NG", Errors: errors}
}
