package validators

// delimeters
const (
	Hyphen     = "-"
	Empty      = ""
	WhiteSpace = " "
	Comma      = ","
	Dot        = "."
)

// Name of the struct tag used in examples.
const (
	Validate = "validate"
	JSON     = "json"
	MinNum   = "min"
	MaxNum   = "max"
	MinLen   = "minlen"
	MaxLen   = "maxlen"
	Regex    = "regex"
	Example  = "example"
)

// values for 'validate' tag
const (
	StringTag      = "string"
	NumTag         = "number"
	EnumTag        = "enum"
	StructTag      = "struct"
	StructArrayTag = "struct_array"
)

// string tag properties
const (
	Min     = "min"
	Max     = "max"
	Expects = "expects"
	Number  = "number"
)

// validator error messages
const (
	EmptyString         = "must not be blank"
	EmptyNumber         = "must not be zero or blank"
	InvalidDataString   = "must be valid data"
	InvalidPrefixString = "must be valid "
	MinLengthError      = "must be at least %v chars long"
	MaxLengthError      = "must be less than %v chars long"
	MinNumberError      = "must be greater than %v"
	MaxNumberError      = "must be less than %v"
	InvalidEnumString   = "must be valid enum value"
)

// Regex : Regular expression to validate string values
const (
	IntRegex  = `^(0|[1-9][0-9]*)$`
	DateRegex = `(((20|21)\\d\\d)-(0?[1-9]|[1][012])-(0?[1-9]|[12][0-9]|3[01])\\s([01]?\\d|2[0-3]):([0-5]\\d):([0-5]\\d))`
)
