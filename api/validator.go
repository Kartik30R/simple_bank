package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/Kartik30R/simple_bank/utils"
)

var currencyValidator validator.Func = func(fl validator.FieldLevel) bool {

	if currency , ok :=fl.Field().Interface().(string); ok{
return utils.IsSupportedCurrency(currency)
	}
	return false
}
