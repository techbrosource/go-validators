package structvalidators

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

// Validate performs validation on given string value
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return false, nil
}

// Validate1 performs validation on given string value
func (v DefaultValidator) Validate1(val interface{}) (bool, interface{}) {
	return false, nil
}
