package validators

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
)

// StringValidator validates string presence and/or its length.
type StringValidator struct {
	MinLen int    `json:"minlen"`
	MaxLen int    `json:"maxlen"`
	Regex  string `json:"regex"`
}

// SetProperties to set properties required for validation
func (v *StringValidator) SetProperties(t reflect.StructTag) {
	m, _ := strconv.Atoi(t.Get(constants.MinLen))
	v.MinLen = m
	m, _ = strconv.Atoi(t.Get(constants.MaxLen))
	v.MaxLen = m
	v.Regex = t.Get(constants.Regex)
}

// Validate performs validation on given string value
func (v *StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf(constants.EmptyString)
	}
	if v.MinLen != 0 && l < v.MinLen {
		return false, fmt.Errorf(constants.MinLengthError, v.MinLen)
	}
	if v.MaxLen >= v.MinLen && v.MaxLen != 0 && l > v.MaxLen {
		return false, fmt.Errorf(constants.MaxLengthError, v.MaxLen)
	}
	regex := regexp.MustCompile(v.Regex)
	r := regex.MatchString(val.(string))
	fmt.Print(r)
	if !regex.MatchString(val.(string)) {
		return false, fmt.Errorf(constants.InvalidString)
	}

	return true, nil
}
