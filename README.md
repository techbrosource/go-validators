# go-validators
Value validations for fields of struct based on custom tags

![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/validator/branches/v9/badge.svg)](https://semaphoreci.com/joeybloggs/validator)
[![Coverage Status](https://coveralls.io/repos/go-playground/validator/badge.svg?branch=v9&service=github)](https://coveralls.io/github/techbrosource/go-validators?branch=v1.1)
[![Go Report Card](https://goreportcard.com/badge/github.com/techbrosource/go-validators)](https://goreportcard.com/report/github.com/techbrosource/go-validators)
[![GoDoc](https://godoc.org/gopkg.in/go-playground/validator.v9?status.svg)](https://github.com/techbrosource/go-validators/blob/master/README.md)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package go-validators provides validations for structs and individual fields based on tags.

It has the following **unique** features:

-   String type fields validations with length validation support (minimum length as well as maximum length)
-   String type fields validation with regex validation support
-   Numeric type (int32 as of now) fields validation with range validation support (minimum as well maximum number range)
-   Ability to validate Date type field
-   Ability to validate nested struct defined in the parent struct as a field
-   Ability to validate nested struct array defined in the parent struct as a field
-   Capability to return array of errors based on tags given with field metadata informations. Like : Field Name, JSON Name of the field, Field Type, Field Value, Field Tags, Example value for the field etc.
-   Ability to dive into both map keys and values for validation  

Installation
------------

Use go get.

	go get github.com/techbrosource/go-validators@1.0.0

Then import the validator package into your own code.

	import "github.com/techbrosource/go-validators/utils"

Error Return Value
-------

Validation functions return type []models.FieldSpecs

Validator return InvalidValidationError for bad validation input along with other metadata of field; so, in your code all you need to do is check if the length of array returned is not 0, like so:

```go
import (
  validator "github.com/techbrosource/go-validators/utils"
  models "github.com/techbrosource/go-validators/models"
)

var fieldErrors []models.FieldSpecs
fieldErrors = validator.Validate(myStruct)
	if len(fieldErrors) != 0 {
		// iterate fieldErrors to get details...
	}

 ```
 
 Usage and documentation
------

Please see https://github.com/techbrosource/go-validators/tree/master/helpdoc for detailed usage docs.

##### Examples:

- [All](https://github.com/techbrosource/go-validators/blob/master/helpdoc/struct_examples.go)

How to Contribute
------

Make a pull request...

License
------
Distributed under MIT License, please see license file within the code for more details.
