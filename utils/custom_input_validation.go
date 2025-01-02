package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Extract custom error messages from struct tags
func ExtractCustomErrorMessage(err error, obj interface{}) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			// Get the field name from the validation error
			fieldName := fieldError.Field()

			// Use reflection to get the `customError` tag value
			field, _ := reflect.TypeOf(obj).Elem().FieldByName(fieldName)
			if customMessage, ok := field.Tag.Lookup("customError"); ok {
				return customMessage
			}
		}
	}

	// Fallback to default error message
	return "Invalid input."
}
