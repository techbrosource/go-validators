package validators

import (
	"reflect"

	constants "github.com/techbrosource/go-validators/constants"
)

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
	Example  string `json:"example"`
	Expects  string `json:"expects"`
	Required string `json:required`
}

// SetProperties to set properties required for validation
func (v *DefaultValidator) SetProperties(t reflect.StructTag) {
	v.Expects = t.Get(constants.Expects)
	v.Example = t.Get(constants.Example)
	v.Required = t.Get(constants.Required)
}

// Validate performs validation on given string value
func (v *DefaultValidator) Validate(val interface{}) (bool, error) {
	return false, nil
}
