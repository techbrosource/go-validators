package utils

import (
	"reflect"

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

	switch tags.Get(constants.ValidatorTag) {
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
func Validate(s interface{}) metadata.Metadatas {
	var fieldMetadatas []metadata.FieldMetadata

	v := reflect.ValueOf(s)
	objectMetadata := metadata.GenerateObjectMetadata(v.Type())

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(constants.ValidatorTag)
		if tag == constants.Empty || tag == constants.Hyphen {
			continue
		}
		validator := getValidatorFromTag(v.Type().Field(i).Tag)
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			fieldMetadatas = append(fieldMetadatas, metadata.GenerateFieldMetadata(v.Type().Field(i), v.Field(i).Interface(), err))
		}
	}
	return metadata.GenerateMetadata(objectMetadata, fieldMetadatas)
}

// ValidateStruct : to perform data validation using validator definitions on the struct
// func ValidateStruct(s interface{}) []error {
// 	var errors []error

// 	v := reflect.ValueOf(s)

// 	for i := 0; i < v.NumField(); i++ {
// 		tag := v.Type().Field(i).Tag.Get(constants.ValidatorTag)
// 		jsonField := v.Type().Field(i).Tag.Get(constants.JSONTag)
// 		if tag == constants.Empty || tag == constants.Hyphen {
// 			continue
// 		}
// 		validator := getValidatorFromTag(tag)
// 		valid, err := validator.Validate(v.Field(i).Interface())
// 		if !valid && err != nil {
// 			errors = append(errors, fmt.Errorf("%s : %s", jsonField, err.Error()))
// 			fmt.Println(jsonField)
// 		}
// 	}
// 	return errors
// }
