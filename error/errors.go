package error

type (
	Error interface {
		Error() string
		HTTPStatusCode() int
		GetErrors() *[]ErrorStruct
	}
)

type ErrorStruct struct {
	Field  string
	Reason string
}

type (
	errors struct {
		Err      error
		HTTPCode int
		Errors   *[]ErrorStruct
	}
)

// NewError - function for initializing error
func NewError(e error, httpCode int, es *[]ErrorStruct) Error {
	return &errors{
		Err:      e,
		HTTPCode: httpCode,
		Errors:   es,
	}
}

func (e *errors) Error() string {
	if e == nil {
		return ""
	}
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e *errors) HTTPStatusCode() int {
	if e == nil {
		return 0
	}

	return e.HTTPCode
}

func (e *errors) GetErrors() *[]ErrorStruct {
	if e == nil {
		return &[]ErrorStruct{}
	}

	return e.Errors
}
