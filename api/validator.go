package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/machingclee/2023-11-04-go-gin/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	currency, ok := fieldLevel.Field().Interface().(string)
	if ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
