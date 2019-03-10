package validators

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	constants "github.com/techbrosource/go-validators/constants"
	validators "github.com/techbrosource/go-validators/structvalidators"
)

// Validator is Generic data validator.
type Validator interface {
	Validate(interface{}) (bool, error)
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, constants.Comma)

	switch args[0] {
	case constants.NumTag:
		validator := validators.NumberValidator{}
		fmt.Println(strings.Join(args[1:], constants.Comma))
		fmt.Sscanf(strings.Join(args[1:], constants.Comma), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case constants.StringTag:
		validator := validators.StringValidator{}
		fmt.Println(strings.Join(args[1:], constants.Comma))
		fmt.Sscanf(strings.Join(args[1:], constants.Comma), "min=%d,max=%d", &validator.Min, &validator.Max)
		fmt.Sscanf(strings.Join(args[1:], constants.Comma), "max=%d", &validator.Max)
		fmt.Println("min = " + strconv.Itoa(validator.Min))
		fmt.Println("max = " + strconv.Itoa(validator.Max))
		return validator
	case constants.StringRegexTag:
		validator := validators.StringRegexValidator{}
		fmt.Println(strings.Join(args[1:], constants.Comma))
		fmt.Sscanf(strings.Join(args[1:], constants.Comma), "regex=%s", &validator.Regex)
		return validator
	case constants.EnumTag:
		validator := validators.EnumValidator{}
		fmt.Sscanf(strings.Join(args[1:], constants.Comma), "name=%s", &validator.Name)
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

type StringTest struct {
	blankStr string `validate:"string,max=10"`
	minStr   string `validate:"string,min=1"`
}

// func main() {
// 	stringTest := StringTest{
// 		Id:   0,
// 		Name: "superlongstring",
// 		Bio:  "",
// 	}

// 	fmt.Println("Errors of string test:")
// 	for i, err := range ValidateStruct(stringTest) {
// 		fmt.Printf("\t%d. %s\n", i+1, err.Error())
// 	}
// }
