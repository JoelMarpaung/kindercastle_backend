package custom

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type echoValidator struct {
	validator *validator.Validate
}

func (ev *echoValidator) Validate(i interface{}) error {
	if err := ev.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewValidator() echo.Validator {
	val := validator.New()

	// register function to get tag name from json tags.
	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		jsonSplitNum := 2
		name := strings.SplitN(fld.Tag.Get("json"), ",", jsonSplitNum)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	val.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		value, ok := field.Interface().(null.Bool)
		if !ok {
			// go-validator uses `reflect.Invalid` as
			// marker for treating data as nonexistent
			return reflect.Invalid
		}
		if !value.Valid {
			return reflect.Invalid
		}
		return reflect.Bool
	}, null.Bool{})

	return &echoValidator{validator: val}
}
