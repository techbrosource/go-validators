package validators

import (
	"fmt"
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// NumberValidator performs numerical value validation.
type NumberValidator struct {
	Min int32
	Max int32
}

// SetProperties to set properties required for validation
func (v *NumberValidator) SetProperties(t reflect.StructTag) {
	m, _ := strconv.ParseInt(t.Get(constants.MinNum), 10, 32)
	v.Min = int32(m)
	m, _ = strconv.ParseInt(t.Get(constants.MaxNum), 10, 32)
	v.Max = int32(m)
}

// Validate performs validation on given int value
func (v *NumberValidator) Validate(val interface{}) (bool, error) {
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
