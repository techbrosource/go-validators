package validators

import (
	"fmt"
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// FloatNumberValidator performs numerical value validation.
type FloatNumberValidator struct {
	Min         float64            `json:"min"`
	Max         float64            `json:"max"`
	Type        string           `json:"types"`
	DefaultProp DefaultValidator `json:"default"`
}

// SetProperties to set properties required for validation
func (v *FloatNumberValidator) SetProperties(tags reflect.StructTag, types reflect.Kind) {
	m, _ := strconv.ParseFloat(tags.Get(constants.MinNum), 64)
	v.Min = m
	m, _ = strconv.ParseFloat(tags.Get(constants.MaxNum), 64)
	v.Max = m
	v.DefaultProp.SetProperties(tags)
	v.Type = types.String()
	return
}

// Validate performs validation on given int value
func (v *FloatNumberValidator) Validate(val interface{}) (bool, error) {
	defaultProp := v.DefaultProp
	var status bool
	var err error
	switch v.Type {
	case "float32":
		status, err = validateFloat32Value(val.(float32), (float32)(v.Min), (float32)(v.Max), defaultProp)
	case "float64":
		status, err = validateFloat64Value(val.(float64), (float64)(v.Min), (float64)(v.Max), defaultProp)
	}
	return status, err
}

func validateFloat32Value(val, min, max float32, defaultProp DefaultValidator) (bool, error) {
	if val == 0 && defaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= min {
			return false, fmt.Errorf(constants.MinNumberError, min)
		}
		if max != 0 && val > max {
			return false, fmt.Errorf(constants.MaxNumberError, max)
		}
	}
	return true, nil
}

func validateFloat64Value(val, min, max float64, defaultProp DefaultValidator) (bool, error) {
	if val == 0 && defaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= min {
			return false, fmt.Errorf(constants.MinNumberError, min)
		}
		if max != 0 && val > max {
			return false, fmt.Errorf(constants.MaxNumberError, max)
		}
	}
	return true, nil
}