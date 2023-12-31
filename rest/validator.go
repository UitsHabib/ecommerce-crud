package rest

import (
	"github.com/UitsHabib/ecommerce-crud/util"
	"github.com/go-playground/validator/v10"
)

var validStatusID validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if statudID, ok := fieldLevel.Field().Interface().(int); ok {
		return util.IsSupportedStatusID(statudID)
	}
	return false
}

var validPhone validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if phone, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedPhone(phone)
	}
	return false
}
