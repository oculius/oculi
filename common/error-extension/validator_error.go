package errext

import (
	"net/http"
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type (
	validatorError struct {
		source   error
		metadata []errorField
	}

	errorField struct {
		Field   string `json:"field"`
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}
)

func NewValidationError(err error, ut ut.Translator) HttpError {
	castedErr, ok := err.(validator.ValidationErrors)
	if !ok {
		panic("error is not validation error")
	}
	return &validatorError{source: err, metadata: buildMetadata(castedErr, ut)}
}

func buildMetadata(err validator.ValidationErrors, ut ut.Translator) []errorField {
	if len(err) == 0 {
		return []errorField{}
	}
	result := make([]errorField, len(err))
	for i, each := range err {
		msg := ""
		if ut != nil {
			msg = each.Translate(ut)
		}
		result[i] = errorField{
			Field:   each.Namespace(),
			Reason:  each.Tag(),
			Message: msg,
		}
	}
	return result
}

func (v *validatorError) Error() string {
	return v.source.Error()
}

func (v *validatorError) ResponseCode() int {
	return ValidatorErrorHttpStatus
}

func (v *validatorError) ResponseStatus() string {
	return http.StatusText(ValidatorErrorHttpStatus)
}

func (v *validatorError) Equal(err error) bool {
	casted, ok := err.(*validatorError)
	if !ok {
		return false
	}
	return v.Error() == err.Error() && reflect.DeepEqual(v.metadata, casted.metadata)
}

func (v *validatorError) Metadata() any {
	return v.metadata
}

func (v *validatorError) Source() error {
	return v.source
}

func (v *validatorError) Detail() string {
	return "validation error"
}
