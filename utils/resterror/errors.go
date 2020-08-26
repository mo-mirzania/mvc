package resterror

import "net/http"

// RestErr struct
type RestErr struct {
	Message string
	Status  int
	Error   string
}

// NotFoundError func
func NotFoundError(msg string) *RestErr {
	restErr := &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found_error",
	}
	return restErr
}
