package validators

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	constants "github.com/techbrosource/go-validators/constants"
)

// StringValidator validates string presence and/or its length.
type StringValidator struct {
	MinLen      int              `json:"minlen"`
	MaxLen      int              `json:"maxlen"`
	Regex       string           `json:"regex"`
	DefaultProp DefaultValidator `json:"default"`
}

// SetProperties to set properties required for validation
func (v *StringValidator) SetProperties(t reflect.StructTag) {
	m, _ := strconv.Atoi(t.Get(constants.MinLen))
	v.MinLen = m
	m, _ = strconv.Atoi(t.Get(constants.MaxLen))
	v.MaxLen = m
	v.Regex = t.Get(constants.Regex)
	v.DefaultProp.SetProperties(t)
}

// Validate performs validation on given string value
func (v *StringValidator) Validate(val interface{}) (bool, error) {
	defaultProp := v.DefaultProp
	l := len(strings.TrimSpace(val.(string)))
	if l == 0 && defaultProp.Required == "true" {
		return false, fmt.Errorf(constants.EmptyString)
	} else if l != 0 {
		if v.MinLen != 0 && l < v.MinLen {
			return false, fmt.Errorf(constants.MinLengthError, v.MinLen)
		}
		if v.MaxLen >= v.MinLen && v.MaxLen != 0 && l > v.MaxLen {
			return false, fmt.Errorf(constants.MaxLengthError, v.MaxLen)
		}
		regex := regexp.MustCompile(v.Regex)

		if !regex.MatchString(val.(string)) {
			if len(defaultProp.Expects) == 0 {
				return false, fmt.Errorf(constants.InvalidDataString)
			}
			return false, fmt.Errorf(constants.InvalidPrefixString + defaultProp.Expects)
		}
	}
	return true, nil
}
