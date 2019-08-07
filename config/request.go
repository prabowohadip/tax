package config

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

func ParsingParamter(ctx echo.Context, i interface{}) error {
	err := ctx.Bind(i)
	return err
}

func ValidateParamter(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	return err
}

func ParsingValidate(c echo.Context, i interface{}) (err error, result interface{}) {
	err = ParsingParamter(c, i)
	if err != nil {
		return err, i
	}
	err = ValidateParamter(i)
	if err != nil {
		return err, i
	}
	return nil, i
}

