package validator

import (
	"fmt"
	"reflect"
	"strings"

	constants "github.com/techbrosource/go-validators/constants"
	validators "github.com/techbrosource/go-validators/struct-validators"
)

// Validator is Generic data validator.
type Validator interface {
	Validate(interface{}) (bool, error)
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case constants.NumTag:
		validator := validators.NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case constants.StringTag:
		validator := validators.StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case constants.StringRegexTag:
		validator := validators.StringRegexValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "regex=%s", &validator.Regex)
		return validator
	case constants.EnumTag:
		validator := validators.EnumValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "name=%s", &validator.Name)
		return validator
	default:
		return validators.DefaultValidator{}
	}
}

// ValidateStruct : to perform data validation using validator definitions on the struct
func ValidateStruct(s interface{}) []error {
	var errors []error

	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(constants.ValidatorTag)
		jsonField := v.Type().Field(i).Tag.Get(constants.JSONTag)
		if tag == constants.Empty || tag == constants.Hyphen {
			continue
		}
		validator := getValidatorFromTag(tag)
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errors = append(errors, fmt.Errorf("%s : %s", jsonField, err.Error()))
		}
	}
	return errors
}

// type User struct {
// 	Id   int `validate:"number,min=1,max=1000"`
// 	Name string
// 	Bio  string
// }

// func main() {
// 	user := User{
// 		Id:   0,
// 		Name: "superlongstring",
// 		Bio:  "",
// 	}

// 	fmt.Println("Errors:")
// 	for i, err := range ValidateStruct(user) {
// 		fmt.Printf("\t%d. %s\n", i+1, err.Error())
// 	}
// }
