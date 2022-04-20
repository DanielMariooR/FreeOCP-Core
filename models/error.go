package models

import (
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (customError CustomError) Error() string {
	return customError.Message
}

var ErrNotFound = &CustomError{
	Code:    http.StatusNotFound,
	Message: "Data Not Found!",
}
