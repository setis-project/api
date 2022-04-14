package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/setis-project/api/core/models"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	}
	return "Unknown error"
}

func GetBindErrors(err error) ([]models.ApiRequestError, bool) {
	var v validator.ValidationErrors
	if errors.As(err, &v) {
		out := make([]models.ApiRequestError, len(v))
		for i, fe := range v {
			out[i] = models.ApiRequestError{Field: fe.Field(), Message: GetErrorMsg(fe)}
		}
		return out, true
	}
	return []models.ApiRequestError{}, false
}
