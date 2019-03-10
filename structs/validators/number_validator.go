package validators

import (
	"fmt"

	constants "github.com/techbrosource/go-validators/constants"
)

// NumberValidator performs numerical value validation.
type NumberValidator struct {
	Min int32
	Max int32
}

// Validate performs validation on given int value
func (v NumberValidator) Validate(val interface{}) (bool, error) {
	num := val.(int32)
	if num == 0 && v.Min == 0 && v.Max == 0 {
		return false, fmt.Errorf(constants.EmptyNumber)
	}
	if num <= v.Min {
		return false, fmt.Errorf(constants.MinNumberError, v.Min)
	}
	if v.Max != 0 && num > v.Max {
		return false, fmt.Errorf(constants.MaxNumberError, v.Max)
	}
	return true, nil
}
