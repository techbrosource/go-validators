package utils

import (
	"reflect"

	constants "github.com/techbrosource/go-validators/constants"
	models "github.com/techbrosource/go-validators/models"
)

// Validate : to perform data validation using validator definitions on the struct
func Validate(s interface{}) []models.FieldSpec {
	var fieldSpecs []models.FieldSpec

	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// tag := v.Type().Field(i).Tag.Get(constants.Validate)
		// if tag == constants.Empty || tag == constants.Hyphen {
		// 	continue
		// }
		jsonName := v.Type().Field(i).Tag.Get(constants.JSON)

		// logic to validate nested struct
		if v.Type().Field(i).Tag.Get(constants.Validate) == constants.StructTag {
			validateNestedStruct(v.Field(i), &fieldSpecs, jsonName)
		}

		// logic to validate nested struct array
		if v.Type().Field(i).Tag.Get(constants.Validate) == "struct_array" {
			nestedArray := reflect.ValueOf(v.Field(i).Interface())
			for k := 0; k < nestedArray.Len(); k++ {
				validateNestedArrayStruct(nestedArray.Index(k), &fieldSpecs, jsonName, k)
			}
		}

		// logic to validate struct fields
		validator := getValidatorFromTag(v.Type().Field(i))
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			fieldSpecs = append(fieldSpecs, models.GenerateFieldSpec(v.Type().Field(i), v.Field(i).Interface(), err))
		}
	}
	return fieldSpecs
}
