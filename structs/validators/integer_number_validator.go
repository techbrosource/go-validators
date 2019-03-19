package validators

import (
	"fmt"
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// IntegerNumberValidator performs numerical value validation.
type IntegerNumberValidator struct {
	Min         int64            `json:"min"`
	Max         int64            `json:"max"`
	Type        string           `json:"types"`
	DefaultProp DefaultValidator `json:"default"`
}

// SetProperties to set properties required for validation
func (v *IntegerNumberValidator) SetProperties(tags reflect.StructTag, types reflect.Kind) {
	m, _ := strconv.ParseInt(tags.Get(constants.MinNum), 10, 64)
	v.Min = m
	m, _ = strconv.ParseInt(tags.Get(constants.MaxNum), 10, 64)
	v.Max = m
	v.DefaultProp.SetProperties(tags)
	v.Type = types.String()
	return
}

// Validate performs validation on given int value
func (v *IntegerNumberValidator) Validate(val interface{}) (bool, error) {
	defaultProp := v.DefaultProp
	var status bool
	var err error
	switch v.Type {
	case "int":
		status, err = validateIntValue(val.(int), (int)(v.Min), (int)(v.Max), defaultProp)
	case "int8":
		status, err = validateInt8Value(val.(int8), (int8)(v.Min), (int8)(v.Max), defaultProp)
	case "int16":
		status, err = validateInt16Value(val.(int16), (int16)(v.Min), (int16)(v.Max), defaultProp)
	case "int32":
		status, err = validateInt32Value(val.(int32), (int32)(v.Min), (int32)(v.Max), defaultProp)
	case "int64":
		status, err = validateInt64Value(val.(int64), (int64)(v.Min), (int64)(v.Max), defaultProp)
	case "uint":
		status, err = validateUintValue(val.(uint), (uint)(v.Min), (uint)(v.Max), defaultProp)
	case "uint8":
		status, err = validateUint8Value(val.(uint8), (uint8)(v.Min), (uint8)(v.Max), defaultProp)
	case "uint16":
		status, err = validateUint16Value(val.(uint16), (uint16)(v.Min), (uint16)(v.Max), defaultProp)
	case "uint32":
		status, err = validateUint32Value(val.(uint32), (uint32)(v.Min), (uint32)(v.Max), defaultProp)
	case "uint64":
		status, err = validateUint64Value(val.(uint64), (uint64)(v.Min), (uint64)(v.Max), defaultProp)
	}
	return status, err
}

func validateIntValue(val, min, max int, defaultProp DefaultValidator) (bool, error) {
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

func validateInt8Value(val, min, max int8, defaultProp DefaultValidator) (bool, error) {
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

func validateInt16Value(val, min, max int16, defaultProp DefaultValidator) (bool, error) {
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

func validateInt32Value(val, min, max int32, defaultProp DefaultValidator) (bool, error) {
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

func validateInt64Value(val, min, max int64, defaultProp DefaultValidator) (bool, error) {
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

func validateUintValue(val, min, max uint, defaultProp DefaultValidator) (bool, error) {
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

func validateUint8Value(val, min, max uint8, defaultProp DefaultValidator) (bool, error) {
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

func validateUint16Value(val, min, max uint16, defaultProp DefaultValidator) (bool, error) {
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

func validateUint32Value(val, min, max uint32, defaultProp DefaultValidator) (bool, error) {
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

func validateUint64Value(val, min, max uint64, defaultProp DefaultValidator) (bool, error) {
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
