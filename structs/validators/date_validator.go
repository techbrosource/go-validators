package validators

import (
	"reflect"
)

// DateValidator does not perform any validations.
type DateValidator struct {
}

// SetProperties to set properties required for validation
func (v *DateValidator) SetProperties(t reflect.StructTag) {
}

// Validate performs validation on given date value
func (v *DateValidator) Validate(val interface{}) (bool, error) {
	return false, nil
}
