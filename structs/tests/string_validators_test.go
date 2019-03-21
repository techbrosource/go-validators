package tests

import (
	"testing"

	constants "github.com/techbrosource/go-validators/constants"
	validate "github.com/techbrosource/go-validators/utils"
	. "gopkg.in/go-playground/assert.v1"
)

func TestRequired(t *testing.T) {
	testData := struct {
		TestField string `json:"test_field" required:"true"`
	}{
		TestField: "",
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), constants.EmptyString)
}

func TestMinLength(t *testing.T) {
	testData := struct {
		TestField string `json:"test_field" minlen:"2" required:"true"`
	}{
		TestField: "1",
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be at least 2 chars long")
}

func TestMaxLength(t *testing.T) {
	testData := struct {
		TestField string `json:"test_field" maxlen:"5" required:"true"`
	}{
		TestField: "123456",
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 5 chars long")
}

func TestLengthRange(t *testing.T) {
	testData := struct {
		MaxTestField   string `json:"min_test_field" minlen:"2" maxlen:"5" example:"str"`
		MinTestField   string `json:"max_test_field" minlen:"2" maxlen:"5"`
		ValidTestField string `json:"valid_test_field" minlen:"2" maxlen:"5"`
	}{
		MaxTestField:   "123456",
		MinTestField:   "1",
		ValidTestField: "1234",
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 2)
	Equal(t, fe[0].GetName(), "MaxTestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 5 chars long")
	Equal(t, fe[0].GetExample(), "str")
	Equal(t, fe[1].GetName(), "MinTestField")
	Equal(t, fe[1].GetError().Error(), "must be at least 2 chars long")
}

func TestRegexMatch(t *testing.T) {
	testData := struct {
		AlphaNumTestField   string `json:"alpha_num_test_field" regex:"^(0|[1-9][0-9]*)$" expects:"number"`
		ZeroPrefixTestField string `json:"zero_prefix_test_field" regex:"^(0|[1-9][0-9]*)$" minlen:"3"`
		MaxLengthTestField  string `json:"max_length_test_field" regex:"^(0|[1-9][0-9]*)$" minlen:"2" maxlen:"5"`
		ValidTestField      string `json:"valid_test_field" regex:"^(0|[1-9][0-9]*)$" minlen:"2" maxlen:"5"`
	}{
		AlphaNumTestField:   "12345s",
		ZeroPrefixTestField: "01",
		MaxLengthTestField:  "1234554s",
		ValidTestField:      "123",
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 3)
	Equal(t, fe[0].GetName(), "AlphaNumTestField")
	Equal(t, fe[0].GetError().Error(), constants.InvalidPrefixString+"number")
	Equal(t, fe[1].GetName(), "ZeroPrefixTestField")
	Equal(t, fe[1].GetError().Error(), "must be at least 3 chars long")
	Equal(t, fe[2].GetName(), "MaxLengthTestField")
	Equal(t, fe[2].GetError().Error(), "must be less than 5 chars long")
}
