package helpdoc

// StringValidators contains all possible validations for string fields
type StringValidators struct {
	// validates whether string is blank or not
	BlankStr string `json:"blank_string" validate:"string"`

	// validates minimum length of the string
	MinStr string `json:"min_string" validate:"string,min=2"`

	// validates maximum length of the string
	MaxStr string `json:"max_string" validate:"string,max=10"`

	// validates string length both minimum and maximum
	MinMaxStr string `json:"min_max_string" validate:"string,min=2,max=10"`

	// validates string containing only numeric values
	NumericStr string `json:"numeric_string" validate:"string,regex=^(0|[1-9][0-9]*)$"`
}
