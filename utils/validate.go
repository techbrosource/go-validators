package utils

import (
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
	metadata "github.com/techbrosource/go-validators/metadata"
	validators "github.com/techbrosource/go-validators/structs/validators"
)

// Validator is Generic data validator.
type Validator interface {
	Validate(interface{}) (bool, error)
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tags reflect.StructTag) Validator {

	switch tags.Get(constants.Validate) {
	case constants.NumTag:
		v := validators.NumberValidator{}
		vp := &v
		vp.SetProperties(tags)
		return vp
	case constants.StringTag:
		v := validators.StringValidator{}
		vp := &v
		vp.SetProperties(tags)
		return vp
	case constants.EnumTag:
		v := validators.EnumValidator{}
		vp := &v
		return vp
	default:
		return &validators.DefaultValidator{}
	}
}

// Validate : to perform data validation using validator definitions on the struct
func Validate(s interface{}) []metadata.FieldSpec {
	var FieldSpecs []metadata.FieldSpec

	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(constants.Validate)
		jsonName := v.Type().Field(i).Tag.Get(constants.JSON)
		if tag == constants.Empty || tag == constants.Hyphen {
			continue
		}
		// logic for nested struct
		if v.Type().Field(i).Tag.Get(constants.Validate) == constants.StructTag {
			netsedFieldSpecs := Validate(v.Field(i).Interface())
			for j := 0; j < len(netsedFieldSpecs); j++ {
				fieldJSONName := (netsedFieldSpecs[j].GetJsonName())
				netsedFieldSpecs[j].SetJsonName(jsonName + constants.Dot + fieldJSONName)
				FieldSpecs = append(FieldSpecs, netsedFieldSpecs[j])
			}
		}
		// logic for nested struct array
		if v.Type().Field(i).Tag.Get(constants.Validate) == "struct_array" {
			nestedArray := reflect.ValueOf(v.Field(i).Interface())
			for k := 0; k < nestedArray.Len(); k++ {
				netsedArrayFieldSpecs := Validate(nestedArray.Index(k).Interface())
				for j := 0; j < len(netsedArrayFieldSpecs); j++ {
					fieldJSONName := (netsedArrayFieldSpecs[j].GetJsonName())
					netsedArrayFieldSpecs[j].SetJsonName(jsonName + "[" + strconv.Itoa(k) + "]" + constants.Dot + fieldJSONName)
					FieldSpecs = append(FieldSpecs, netsedArrayFieldSpecs[j])
				}
			}
		}
		validator := getValidatorFromTag(v.Type().Field(i).Tag)
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			FieldSpecs = append(FieldSpecs, metadata.GenerateFieldSpec(v.Type().Field(i), v.Field(i).Interface(), err))
		}
	}
	return FieldSpecs
}
