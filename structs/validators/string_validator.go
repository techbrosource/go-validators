package validators

import (
	"fmt"
	"regexp"

	constants "github.com/techbrosource/go-validators/constants"
)

// StringValidator validates string presence and/or its length.
type StringValidator struct {
	Min   int    `json:"min"`
	Max   int    `json:"max"`
	Regex string `json:"regex"`
}

// Validate performs validation on given string value
func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf(constants.EmptyString)
	}
	if v.Min != 0 && l < v.Min {
		return false, fmt.Errorf(constants.MinLengthError, v.Min)
	}
	if v.Max >= v.Min && v.Max != 0 && l > v.Max {
		return false, fmt.Errorf(constants.MaxLengthError, v.Max)
	}
	regex := regexp.MustCompile(v.Regex)
	if !regex.MatchString(val.(string)) {
		return false, fmt.Errorf(constants.InvalidString)
	}

	return true, nil
}
