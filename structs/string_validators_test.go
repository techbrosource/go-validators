package structs

import (
	"testing"

	constants "github.com/techbrosource/go-validators/constants"
	validate "github.com/techbrosource/go-validators/utils"
	. "gopkg.in/go-playground/assert.v1"
)

func TestRequired(t *testing.T) {
	testData := struct {
		TestField string `json:"test_field" validate:"string"`
	}{
		TestField: "",
	}

	m := validate.Validate(testData)
	NotEqual(t, m, nil)

	fe := m.GetFields()
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), constants.EmptyString)
}

func TestMinLength(t *testing.T) {
	testData := struct {
		TestField string `json:"test_field" validate:"string,min=2"`
	}{
		TestField: "1",
	}

	m := validate.Validate(testData)
	NotEqual(t, m, nil)

	fe := m.GetFields()
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be at least 2 chars long")
}

func TestMaxLength(t *testing.T) {
	testData := struct {
		TestField string `json:"test_field" validate:"string,max=5"`
	}{
		TestField: "123456",
	}

	m := validate.Validate(testData)
	NotEqual(t, m, nil)

	fe := m.GetFields()
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 5 chars long")
}

func TestLengthRange(t *testing.T) {
	testData := struct {
		MaxTestField   string `json:"min_test_field" validate:"string,min=2,max=5"`
		MinTestField   string `json:"max_test_field" validate:"string,min=2,max=5"`
		ValidTestField string `json:"valid_test_field" validate:"string,min=2,max=5"`
	}{
		MaxTestField:   "123456",
		MinTestField:   "1",
		ValidTestField: "1234",
	}

	m := validate.Validate(testData)
	NotEqual(t, m, nil)

	fe := m.GetFields()
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 2)
	Equal(t, fe[0].GetName(), "MaxTestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 5 chars long")
	Equal(t, fe[1].GetName(), "MinTestField")
	Equal(t, fe[1].GetError().Error(), "must be at least 2 chars long")
}

func TestRegexMatch(t *testing.T) {
	testData := struct {
		AlphaNumTestField   string `json:"alpha_num_test_field" validate:"string,regex=^(0|[1-9][0-9]*)$"`
		ZeroPrefixTestField string `json:"zero_prefix_test_field" validate:"string,regex=^(0|[1-9][0-9]*)$"`
		ValidTestField      string `json:"valid_test_field" validate:"string,regex=^(0|[1-9][0-9]*)$"`
	}{
		AlphaNumTestField:   "12345s",
		ZeroPrefixTestField: "012",
		ValidTestField:      "12345",
	}

	m := validate.Validate(testData)
	NotEqual(t, m, nil)

	fe := m.GetFields()
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 2)
	Equal(t, fe[0].GetName(), "AlphaNumTestField")
	Equal(t, fe[0].GetError().Error(), constants.InvalidString)
	Equal(t, fe[1].GetName(), "ZeroPrefixTestField")
	Equal(t, fe[1].GetError().Error(), constants.InvalidString)
}
