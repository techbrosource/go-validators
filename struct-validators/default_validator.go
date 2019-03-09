package validator

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

// Validate performs validation on given string value
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return false, nil
}
