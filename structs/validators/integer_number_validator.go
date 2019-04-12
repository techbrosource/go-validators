package validators

import (
	"fmt"
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// IntegerNumberValidator performs numerical value validation.
type IntegerNumberValidator struct {
	GreaterThanEqual int64            `json:"gte"`
	LessThanEqual    int64            `json:"lte"`
	GreaterThan      int64            `json:"gt"`
	LessThan         int64            `json:"lt"`
	Type             string           `json:"types"`
	DefaultProp      DefaultValidator `json:"default"`
}

// SetProperties to set properties required for validation
func (v *IntegerNumberValidator) SetProperties(tags reflect.StructTag, types reflect.Kind) {
	m, _ := strconv.ParseInt(tags.Get(constants.GTE), 10, 64)
	v.GreaterThanEqual = m
	m, _ = strconv.ParseInt(tags.Get(constants.LTE), 10, 64)
	v.LessThanEqual = m
	m, _ = strconv.ParseInt(tags.Get(constants.GT), 10, 64)
	v.GreaterThan = m
	m, _ = strconv.ParseInt(tags.Get(constants.LT), 10, 64)
	v.LessThan = m
	v.DefaultProp.SetProperties(tags)
	v.Type = types.String()
	return
}

// Validate performs validation on given int value
func (v *IntegerNumberValidator) Validate(val interface{}) (bool, error) {
	var status bool
	var err error
	switch v.Type {
	case "int":
		status, err = validateIntValue(v, val.(int))
	case "int8":
		status, err = validateInt8Value(v, val.(int8))
	case "int16":
		status, err = validateInt16Value(v, val.(int16))
	case "int32":
		status, err = validateInt32Value(v, val.(int32))
	case "int64":
		status, err = validateInt64Value(v, val.(int64))
	case "uint":
		status, err = validateUintValue(v, val.(uint))
	case "uint8":
		status, err = validateUint8Value(v, val.(uint8))
	case "uint16":
		status, err = validateUint16Value(v, val.(uint16))
	case "uint32":
		status, err = validateUint32Value(v, val.(uint32))
	case "uint64":
		status, err = validateUint64Value(v, val.(uint64))
	}
	return status, err
}

func validateIntValue(v *IntegerNumberValidator, val int) (bool, error) {
	gt := (int)(v.GreaterThan)
	lt := (int)(v.LessThan)
	gte := (int)(v.GreaterThanEqual)
	lte := (int)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateInt8Value(v *IntegerNumberValidator, val int8) (bool, error) {
	gt := (int8)(v.GreaterThan)
	lt := (int8)(v.LessThan)
	gte := (int8)(v.GreaterThanEqual)
	lte := (int8)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateInt16Value(v *IntegerNumberValidator, val int16) (bool, error) {
	gt := (int16)(v.GreaterThan)
	lt := (int16)(v.LessThan)
	gte := (int16)(v.GreaterThanEqual)
	lte := (int16)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateInt32Value(v *IntegerNumberValidator, val int32) (bool, error) {
	gt := (int32)(v.GreaterThan)
	lt := (int32)(v.LessThan)
	gte := (int32)(v.GreaterThanEqual)
	lte := (int32)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateInt64Value(v *IntegerNumberValidator, val int64) (bool, error) {
	gt := (int64)(v.GreaterThan)
	lt := (int64)(v.LessThan)
	gte := (int64)(v.GreaterThanEqual)
	lte := (int64)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateUintValue(v *IntegerNumberValidator, val uint) (bool, error) {
	gt := (uint)(v.GreaterThan)
	lt := (uint)(v.LessThan)
	gte := (uint)(v.GreaterThanEqual)
	lte := (uint)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateUint8Value(v *IntegerNumberValidator, val uint8) (bool, error) {
	gt := (uint8)(v.GreaterThan)
	lt := (uint8)(v.LessThan)
	gte := (uint8)(v.GreaterThanEqual)
	lte := (uint8)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateUint16Value(v *IntegerNumberValidator, val uint16) (bool, error) {
	gt := (uint16)(v.GreaterThan)
	lt := (uint16)(v.LessThan)
	gte := (uint16)(v.GreaterThanEqual)
	lte := (uint16)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateUint32Value(v *IntegerNumberValidator, val uint32) (bool, error) {
	gt := (uint32)(v.GreaterThan)
	lt := (uint32)(v.LessThan)
	gte := (uint32)(v.GreaterThanEqual)
	lte := (uint32)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}

func validateUint64Value(v *IntegerNumberValidator, val uint64) (bool, error) {
	gt := (uint64)(v.GreaterThan)
	lt := (uint64)(v.LessThan)
	gte := (uint64)(v.GreaterThanEqual)
	lte := (uint64)(v.LessThanEqual)
	if val == 0 && v.DefaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyNumber)
	} else if val != 0 {
		if val <= gt {
			return false, fmt.Errorf(constants.GreaterThanNumberError, gt)
		}
		if lt != 0 && val >= lt {
			return false, fmt.Errorf(constants.LessThanNumberError, lt)
		}
		if val < gte {
			return false, fmt.Errorf(constants.GreaterThanEqualNumberError, gte)
		}
		if lte != 0 && val > lte {
			return false, fmt.Errorf(constants.LessThanEqualNumberError, lte)
		}
	}
	return true, nil
}
