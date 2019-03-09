package validator

import (
	"fmt"
	"regexp"

	constants "github.com/techbrosource/go-validators/constants"
)

// StringValidator validates string presence and/or its length.
type StringValidator struct {
	Min int
	Max int
}

// Validate performs validation on given string value
func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf(constants.EmptyString)
	}

	return true, nil
}

// StringLengthValidator validates string presence and/or its length.
type StringLengthValidator struct {
	Min int
	Max int
}

// Validate performs validation on given string value
func (v StringLengthValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf(constants.EmptyString)
	}
	if v.Min == 0 && v.Max == 0 {
		return true, nil
	}
	if l < v.Min {
		return false, fmt.Errorf(constants.MinLengthError, v.Min)
	}
	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf(constants.MaxLengthError, v.Max)
	}

	return true, nil
}

var intRegex = regexp.MustCompile(constants.IntRegex)

// StringRegexValidator validates string presence and/or its length.
type StringRegexValidator struct {
	Regex string
}

// Validate performs validation on given string value
func (v StringRegexValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
		return false, fmt.Errorf(constants.EmptyString)
	}
	if !intRegex.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid email address")
	}

	return true, nil
}
