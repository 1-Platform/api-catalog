package response

import "github.com/go-playground/validator/v10"

type Response struct {
	Err            error       `json:"-"`               // low-level runtime error
	HTTPStatusCode int         `json:"-"`               // http response status code
	Success        bool        `json:"success"`         // flag to indicated failed to success
	Message        string      `json:"message"`         // user-level status message
	AppCode        int64       `json:"code,omitempty"`  // application-specific error code
	ErrorText      interface{} `json:"error,omitempty"` // application-level error message, for debugging
	Data           interface{} `json:"data,omitempty"`
}

func (res *Response) Error() string {
	return res.Message
}

func Success(data interface{}, message string) *Response {
	return &Response{
		Message: message,
		Success: true,
		Data:    data,
	}
}

func ErrValidationFailed(err validator.ValidationErrors) *Response {
	errMap := make(map[string]interface{})

	for _, val := range err {
		errMap[val.Field()] = val.Error()
	}

	return &Response{
		Err:            err,
		HTTPStatusCode: 403,
		Message:        "Validation failed",
		ErrorText:      errMap,
		Success:        false,
	}
}

func ErrInvalidRequest(err error) *Response {
	return &Response{
		Err:            err,
		HTTPStatusCode: 400,
		Message:        "Invalid request",
		ErrorText:      err.Error(),
		Success:        false,
	}
}

func ErrUnauthorizedAccess(err error) *Response {
	return &Response{
		Err:            err,
		HTTPStatusCode: 401,
		Message:        "Unauthorized access",
		ErrorText:      err.Error(),
		Success:        false,
	}
}
