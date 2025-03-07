package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var ValidateDrinkVariant validator.Func = func(fl validator.FieldLevel) bool {
	curr, ok := fl.Field().Interface().(DrinkVariant)
	if ok {
		return curr == ICED || curr == HOT
	}
	return false
}

var ValidateDrinkCategory validator.Func = func(fl validator.FieldLevel) bool {
	curr, ok := fl.Field().Interface().(DrinkCategory)
	if ok {
		return curr == COFFEE || curr == NON_COFFEE
	}
	return false
}

func IsValidCollectionDate(collectionTime time.Time) bool {
	if collectionTime.Before(time.Now()) {
		return false
	}
	return true
}
