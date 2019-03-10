package validators_test

import (
	"testing"

	validate "github.com/techbrosource/go-validators"
	. "gopkg.in/go-playground/assert.v1"
)

type StringTest struct {
	BlankStr string `json:"blank_string" validate:"string"`
	MinStr   string `json:"min_string" validate:"string,min=1"`
	// MaxStr    string `json:"max_string" validate:"string,max=10"`
	MinMaxStr string `json:"min_max_string" validate:"string,min=2,max=10"`
}

func TestStructLevelInvalidError(t *testing.T) {

	var test StringTest

	err := validate.ValidateStruct(test)
	NotEqual(t, err, nil)

	// append()

	// errs, ok := err.(validate.ValidationErrors)
	// Equal(t, ok, true)

	// fe := errs[0]
	// Equal(t, fe.Field(), "Value")
	// Equal(t, fe.StructField(), "Value")
	// Equal(t, fe.Namespace(), "StructLevelInvalidErr.Value")
	// Equal(t, fe.StructNamespace(), "StructLevelInvalidErr.Value")
	// Equal(t, fe.Tag(), "required")
	// Equal(t, fe.ActualTag(), "required")
	// Equal(t, fe.Kind(), reflect.Invalid)
	// Equal(t, fe.Type(), reflect.TypeOf(nil))
}
