package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadsarimin/indocyber-api/models/cerr"
)

var validation *validator.Validate

func init() {
	validation = validator.New()
	validation.RegisterTagNameFunc(jsonName)
}

func Validate(T interface{}) error {

	e := validation.Struct(T)
	if e != nil {

		err := e.(validator.ValidationErrors)[0]

		if err.Tag() == "required" {
			return cerr.GetError("002", err.Field())
		}

		return cerr.GetError("001", err.Field())
	}

	return nil
}

func jsonName(fld reflect.StructField) string {

	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}
