package helper

import "github.com/go-playground/validator/v10"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
