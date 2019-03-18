package structs

import (
	"testing"

	constants "github.com/techbrosource/go-validators/constants"
	validate "github.com/techbrosource/go-validators/utils"
	. "gopkg.in/go-playground/assert.v1"
)

func TestNumberRequired(t *testing.T) {
	testData := struct {
		TestField int32 `json:"test_field" validate:"number" required:"true"`
	}{
		TestField: 0,
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), constants.EmptyNumber)
}

func TestMinNumber(t *testing.T) {
	testData := struct {
		TestField int32 `json:"test_field" validate:"number" min:"2"`
	}{
		TestField: 1,
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be greater than 2")
}

func TestMaxNumber(t *testing.T) {
	testData := struct {
		TestField int32 `json:"test_field" validate:"number" max:"5"`
	}{
		TestField: 123456,
	}

	fe := validate.Validate(testData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 5")
}
