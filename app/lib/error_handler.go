package lib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorHandler struct{}

// NewErrorHandler creates a new error handler
func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}

func (e *ErrorHandler) IsValid(model any) error {
	validate := customValidator()
	if err := validate.Struct(model); err != nil {
		return err
	}
	return nil
}

func (e *ErrorHandler) ParseValidationErrors(err error) map[string]string {
	vErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}
	errs := make(map[string]string)

	for _, f := range vErr {
		currErr := f.ActualTag()
		if f.Param() != "" {
			currErr = fmt.Sprintf("%s=%s", currErr, f.Param())
		}
		errs[f.Field()] = currErr
	}

	return errs
}
