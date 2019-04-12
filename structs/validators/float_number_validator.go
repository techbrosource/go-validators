package validators

import (
	"fmt"
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// FloatNumberValidator performs numerical value validation.
type FloatNumberValidator struct {
	GreaterThanEqual float64          `json:"gte"`
	LessThanEqual    float64          `json:"lte"`
	GreaterThan      float64          `json:"gt"`
	LessThan         float64          `json:"lt"`
	Type             string           `json:"types"`
	DefaultProp      DefaultValidator `json:"default"`
}

// SetProperties to set properties required for validation
func (v *FloatNumberValidator) SetProperties(tags reflect.StructTag, types reflect.Kind) {
	m, _ := strconv.ParseFloat(tags.Get(constants.GTE), 64)
	v.GreaterThanEqual = m
	m, _ = strconv.ParseFloat(tags.Get(constants.LTE), 64)
	v.LessThanEqual = m
	m, _ = strconv.ParseFloat(tags.Get(constants.GT), 64)
	v.GreaterThan = m
	m, _ = strconv.ParseFloat(tags.Get(constants.LT), 64)
	v.LessThan = m
	v.DefaultProp.SetProperties(tags)
	v.Type = types.String()
	return
}

// Validate performs validation on given int value
func (v *FloatNumberValidator) Validate(val interface{}) (bool, error) {
	var status bool
	var err error
	switch v.Type {
	case "float32":
		status, err = validateFloat32Value(v, val.(float32))
	case "float64":
		status, err = validateFloat64Value(v, val.(float64))
	}
	return status, err
}

func validateFloat32Value(v *FloatNumberValidator, val float32) (bool, error) {
	gt := (float32)(v.GreaterThan)
	lt := (float32)(v.LessThan)
	gte := (float32)(v.GreaterThanEqual)
	lte := (float32)(v.LessThanEqual)
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

func validateFloat64Value(v *FloatNumberValidator, val float64) (bool, error) {
	gt := (float64)(v.GreaterThan)
	lt := (float64)(v.LessThan)
	gte := (float64)(v.GreaterThanEqual)
	lte := (float64)(v.LessThanEqual)
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
