package tests

import (
	"testing"

	constants "github.com/techbrosource/go-validators/constants"
	validate "github.com/techbrosource/go-validators/utils"
	. "gopkg.in/go-playground/assert.v1"
)

func TestFloatNumberRequired(t *testing.T) {
	testData := struct {
		TestField float32 `json:"test_field" required:"true"`
	}{
		TestField: 0.0,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), constants.EmptyNumber)
}

func TestGteFloatNumber(t *testing.T) {
	testData := struct {
		TestField float64 `json:"test_field" gte:"1200.01"`
	}{
		TestField: 1200.00,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be greater than equal to 1200.01")
}

func TestLteFloatNumber(t *testing.T) {
	testData := struct {
		TestField float32 `json:"test_field" lte:"1200.01"`
	}{
		TestField: 1201,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than equal to 1200.01")
}

func TestGtFloatNumber(t *testing.T) {
	testData := struct {
		TestField float64 `json:"test_field" gt:"1200.02"`
	}{
		TestField: 1200.01,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be greater than 1200.02")
}

func TestLtFloatNumber(t *testing.T) {
	testData := struct {
		TestField float32 `json:"test_field" lt:"1200.05"`
	}{
		TestField: 1200.06,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 1200.05")
}

func TestGtAndLtFloatNumber(t *testing.T) {
	testData := struct {
		TestField float32 `json:"test_field" gt:"1100" lt:"1200.105"`
	}{
		TestField: 1200.106,
	}

	fe := validate.Validate(testData)
	Equal(t, len(fe), 1)
	Equal(t, fe[0].GetName(), "TestField")
	Equal(t, fe[0].GetError().Error(), "must be less than 1200.105")
}
