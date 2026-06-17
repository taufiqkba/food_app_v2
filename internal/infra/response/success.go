package response

import "net/http"

func newSuccess(httpStatusCode int, msg string, opts ...OptionResponse) Response {
	var resp = Response{
		HttpStatus: httpStatusCode,
		Message:    msg,
		Success:    true,
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}

func NewSuccessCreated(msg string, opts ...OptionResponse) Response {
	return newSuccess(http.StatusCreated, msg, opts...)
}

func NewSuccessOK(msg string, opts ...OptionResponse) Response {
	return newSuccess(http.StatusOK, msg, opts...)
}
