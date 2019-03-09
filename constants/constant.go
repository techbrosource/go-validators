package validators

// delimeters
const (
	Hyphen     = "-"
	Empty      = ""
	WhiteSpace = " "
)

// Name of the struct tag used in examples.
const (
	ValidatorTag = "validate"
	JSONTag      = "json"
)

// validator tag types
const (
	StringTag       = "string"
	StringLengthTag = "length_of_string"
	StringRegexTag  = "string_regex"
	NumTag          = "number"
	EnumTag         = "enum"
)

// validator error messages
const (
	EmptyString       = "must not be blank"
	EmptyNumber       = "must not be zero or blank"
	InvalidString     = "Invalid"
	MinLengthError    = "must be at least %v chars long"
	MaxLengthError    = "must be less than %v chars long"
	MinNumberError    = "must be greater than %v"
	MaxNumberError    = "must be less than %v"
	InvalidEnumString = "must be valid enum value"
)

// Regex : Regular expression to validate string values
const (
	IntRegex  = `^(0|[1-9][0-9]*)$`
	DateRegex = `(((20|21)\\d\\d)-(0?[1-9]|[1][012])-(0?[1-9]|[12][0-9]|3[01])\\s([01]?\\d|2[0-3]):([0-5]\\d):([0-5]\\d))`
)
