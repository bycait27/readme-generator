package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates any struct with validation tags
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		return makeErrorReadable(err)
	}
	return nil
}

// makeErrorReadable converts validator errors to simple messages
func makeErrorReadable(err error) error {
	var validationErrors validator.ValidationErrors

	// check if it's a validation error
	if errors.As(err, &validationErrors) {
		var messages []string

		// convert each error to a readable message
		for _, e := range validationErrors {
			message := makeFieldErrorReadable(e)
			messages = append(messages, message)
		}

		// join all messages
		return fmt.Errorf("validation failed: %s", strings.Join(messages, ", "))
	}

	// if it's not a validation error, return as-is
	return err
}

// makeFieldErrorReadable converts a single field error to readable text
func makeFieldErrorReadable(e validator.FieldError) string {
	field := e.Field()

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, e.Param())
	case "oneof":
		options := strings.ReplaceAll(e.Param(), " ", ", ")
		return fmt.Sprintf("%s must be one of: %s", field, options)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
