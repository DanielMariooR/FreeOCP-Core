package helper

import (
	"net/http"
	"net/url"
)

type RequestParams struct {
	URL     string
	Method  string
	Header  *http.Header
	Payload interface{}
}

type ErrorStruct struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type FormRequestParams struct {
	URL     string
	Method  string
	Header  *http.Header
	Payload url.Values
}

type RequestError struct {
	Error interface{} `json:"error"`
}

type RequestSingleError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type RequestErrors struct {
	StatusCode int           `json:"code"`
	Message    string        `json:"message"`
	Errors     []ErrorStruct `json:"errors"`
}
