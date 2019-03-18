package validators

import (
	"fmt"
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// NumberValidator performs numerical value validation.
type NumberValidator struct {
	Min         int32            `json:"min"`
	Max         int32            `json:"max"`
	DefaultProp DefaultValidator `json:"default"`
}

// SetProperties to set properties required for validation
func (v *NumberValidator) SetProperties(t reflect.StructTag) {
	m, _ := strconv.ParseInt(t.Get(constants.MinNum), 10, 32)
	v.Min = int32(m)
	m, _ = strconv.ParseInt(t.Get(constants.MaxNum), 10, 32)
	v.Max = int32(m)
	v.DefaultProp.SetProperties(t)
}

// Validate performs validation on given int value
func (v *NumberValidator) Validate(val interface{}) (bool, error) {
	defaultProp := v.DefaultProp
	num := val.(int32)
	if num == 0 && defaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if num != 0 {
		if num <= v.Min {
			return false, fmt.Errorf(constants.MinNumberError, v.Min)
		}
		if v.Max != 0 && num > v.Max {
			return false, fmt.Errorf(constants.MaxNumberError, v.Max)
		}
	}
	return true, nil
}
