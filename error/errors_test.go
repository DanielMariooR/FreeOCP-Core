package error

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors_Error(t *testing.T) {
	tableTest := []struct {
		Name        string
		Err         *errors
		ExpectedMsg string
	}{
		{
			Name:        "should return empty string error message",
			Err:         nil,
			ExpectedMsg: "",
		},
		{
			Name:        "should return error message",
			Err:         &errors{Err: fmt.Errorf("something went wrong")},
			ExpectedMsg: "something went wrong",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.Name, func(t *testing.T) {
			msg := tt.Err.Error()
			assert.Equal(t, tt.ExpectedMsg, msg)
		})
	}
}

func TestErrors_HTTPStatusCode(t *testing.T) {
	tableTest := []struct {
		Name         string
		Err          *errors
		ExpectedCode int
	}{
		{
			Name:         "no error, should return code: 0 ",
			Err:          nil,
			ExpectedCode: 0,
		},
		{
			Name:         "should return code: 201 ",
			Err:          &errors{HTTPCode: http.StatusCreated},
			ExpectedCode: 201,
		},
		{
			Name:         "should return code: 200",
			Err:          &errors{HTTPCode: http.StatusOK},
			ExpectedCode: 200,
		},
		{
			Name:         "should return code: 500",
			Err:          &errors{HTTPCode: http.StatusInternalServerError},
			ExpectedCode: 500,
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.Name, func(t *testing.T) {
			code := tt.Err.HTTPStatusCode()
			assert.Equal(t, tt.ExpectedCode, code)
		})
	}
}
