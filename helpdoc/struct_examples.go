package helpdoc

/*
validate:"string"
minlen:"2"
maxlen:"10"
regex:"^(0|[1-9][0-9]*)$"
example:"test string"
expects:"date"
*/
// StringValidators contains all possible validations for string fields
type StringValidators struct {
	// validates whether string is blank or not
	BlankStr string `json:"blank_string" validate:"string" required:"true"`

	// validates minimum length of the string
	MinStr string `json:"min_string" validate:"string" minlen:"2" required:"true"`

	// validates maximum length of the string
	MaxStr string `json:"max_string" validate:"string" maxlen:"10" required:"true"`

	// validates string length both minimum and maximum
	MinMaxStr string `json:"min_max_string" validate:"string" minlen:"2" maxlen:"10" required:"false"`

	// validates string containing only numeric values
	NumericStr string `json:"numeric_string" validate:"string" regex:"^(0|[1-9][0-9]*)$" required:"true"`

	// validates string containing only numeric values with minimum length
	NumericMinStr string `json:"numeric_min_string" validate:"string" regex:"^(0|[1-9][0-9]*)$" minlen:"2" required:"false"`

	// validates string containing only numeric values with minimum length
	NumericRangeStr string `json:"numeric_range_string" validate:"string" regex:"^(0|[1-9][0-9]*)$" minlen:"2" maxlen:"10"`

	// validates string containing only numeric values and mentioned the expectation
	NumericExpectsStr string `json:"numeric_expects_string" validate:"string" regex:"^(0|[1-9][0-9]*)$" expects:"number"`

	// validates string containing only numeric values with example
	NumericExampleStr string `json:"numeric_example_string" validate:"string" regex:"^(0|[1-9][0-9]*)$" example:"213"`
}

/*
validate:"struct"
validate:"struct_array"
*/
// StringValidators contains all possible validations for string fields
type StructValidators struct {
	// validates whether string is blank or not
	BlankStr string `json:"blank_string" validate:"string" required:"true"`

	// validates all cases for the nested struct
	NestedStruct StringValidators `json:"nested_struct" validate:"struct"`

	// validates all cases for the nested array of struct
	NestedArrayStruct StringValidators `json:"nested_array_struct" validate:"struct_array"`
}
