package tests

import (
	"testing"

	constants "github.com/techbrosource/go-validators/constants"
	validate "github.com/techbrosource/go-validators/utils"
	. "gopkg.in/go-playground/assert.v1"
)

func TestNumberRequired(t *testing.T) {
	testData := struct {
		TestField int16 `json:"test_field" required:"true"`
	}{
		TestField: 0,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), constants.EmptyNumber)
}

func TestGteNumber(t *testing.T) {
	testData := struct {
		TestField int64 `json:"test_field" gte:"1200"`
	}{
		TestField: 1199,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be greater than equal to 1200")
}

func TestLteNumber(t *testing.T) {
	testData := struct {
		TestField int32 `json:"test_field" lte:"1200"`
	}{
		TestField: 1201,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than equal to 1200")
}

func TestGtNumber(t *testing.T) {
	testData := struct {
		TestField int64 `json:"test_field" gt:"1200"`
	}{
		TestField: 1200,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be greater than 1200")
}

func TestLtNumber(t *testing.T) {
	testData := struct {
		TestField int32 `json:"test_field" lt:"1200"`
	}{
		TestField: 1200,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 1200")
}

func TestGtAndLtNumber(t *testing.T) {
	testData := struct {
		TestField int32 `json:"test_field" gt:"1100" lt:"1200"`
	}{
		TestField: 1200,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 1200")
}
