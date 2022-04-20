package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/constant"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases"
)

type (
	CustomError struct {
		Field  string `json:"field"`
		Reason string `json:"reason"`
	}
	customValidator struct {
		validator *validator.Validate
	}
)

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Init(e *echo.Echo, dbManager *databases.Manager) {
	validate := validator.New()
	registerTagNameWithLabel(validate)
	e.Validator = &customValidator{validator: validate}
}

func BuildCustomErrors(err error) []CustomError {
	errors := []CustomError{}

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, buildCustomError(err))
	}

	return errors
}

func registerTagNameWithJson(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

func registerTagNameWithLabel(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")

		if name == "-" {
			return ""
		}

		return name
	})
}

// Add Extra case in switch to add new generic error
func buildCustomError(err validator.FieldError) CustomError {
	var reason string
	customField := buildCustomField(err)
	errField := err.Field()

	switch err.Tag() {
	case "required":
		reason = fmt.Sprintf("%s tidak boleh kosong!", constant.ValidationFields[errField])
	}

	return CustomError{Field: customField, Reason: reason}
}

func buildCustomField(err validator.FieldError) string {
	splits := strings.Split(err.Namespace(), ".")
	splitsCount := len(splits)
	field := splits[1]

	for i := 2; i < splitsCount; i++ {
		field += "." + splits[i]
	}

	return field
}
