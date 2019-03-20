package utils

import (
	"reflect"
	"strconv"

	constants "github.com/techbrosource/go-validators/constants"
	models "github.com/techbrosource/go-validators/models"
	validators "github.com/techbrosource/go-validators/structs/validators"
)

// Validator is Generic data validator.
type Validator interface {
	Validate(interface{}) (bool, error)
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(field reflect.StructField) Validator {
	tags := field.Tag
	types := field.Type.Kind()
	switch types {
	case reflect.String:
		v := validators.StringValidator{}
		vp := &v
		vp.SetProperties(tags)
		return vp
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := validators.IntegerNumberValidator{}
		vp := &v
		vp.SetProperties(tags, types)
		return vp
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := validators.IntegerNumberValidator{}
		vp := &v
		vp.SetProperties(tags, types)
		return vp
	case reflect.Float32, reflect.Float64:
		v := validators.FloatNumberValidator{}
		vp := &v
		vp.SetProperties(tags, types)
		return vp
	default:
		return &validators.DefaultValidator{}
	}
	// switch tags.Get(constants.Validate) {
	// case constants.NumTag:
	// 	v := validators.NumberValidator{}
	// 	vp := &v
	// 	vp.SetProperties(tags, types)
	// 	return vp
	// case constants.StringTag:
	// 	v := validators.StringValidator{}
	// 	vp := &v
	// 	vp.SetProperties(tags)
	// 	return vp
	// default:
	// 	return &validators.DefaultValidator{}
	// }
}

func validateNestedStruct(v reflect.Value, fs *[]models.FieldSpec, jsonName string) {
	netsedFieldSpecs := Validate(v.Interface())
	for j := 0; j < len(netsedFieldSpecs); j++ {
		fieldJSONName := (netsedFieldSpecs[j].GetJsonName())
		netsedFieldSpecs[j].SetJsonName(jsonName + constants.Dot + fieldJSONName)
		*fs = append(*fs, netsedFieldSpecs[j])
	}
}

func validateNestedArrayStruct(v reflect.Value, fs *[]models.FieldSpec, jsonName string, k int) {
	netsedArrayFieldSpecs := Validate(v.Interface())
	for j := 0; j < len(netsedArrayFieldSpecs); j++ {
		fieldJSONName := (netsedArrayFieldSpecs[j].GetJsonName())
		netsedArrayFieldSpecs[j].SetJsonName(jsonName + "[" + strconv.Itoa(k) + "]" + constants.Dot + fieldJSONName)
		*fs = append(*fs, netsedArrayFieldSpecs[j])
	}
}
