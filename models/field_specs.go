package metadata

import (
	"reflect"

	constants "github.com/techbrosource/go-validators/constants"
)

// IFieldSpec contains all functions to get field metadata details
type IFieldSpec interface {
	GetTag() string

	GetName() string

	GetJsonName() string

	GetValue() interface{}

	GetType() reflect.Type

	GetError() error

	GetExample() string

	SetJSONName(n string)
}

// FieldSpec stores information about fields
type FieldSpec struct {
	name     string
	jsonName string
	tags     string
	value    interface{}
	typ      reflect.Type
	errors   error
	example  string
}

// GetTag returns the validation tag.
func (fm *FieldSpec) GetTag() string {
	return fm.tags
}

// GetName returns the name.
func (fm *FieldSpec) GetName() string {
	return fm.name
}

// GetJsonName returns the json name.
func (fm *FieldSpec) GetJsonName() string {
	return fm.jsonName
}

// GetValue returns the value.
func (fm *FieldSpec) GetValue() interface{} {
	return fm.value
}

// GetType returns the type.
func (fm *FieldSpec) GetType() reflect.Type {
	return fm.typ
}

// GetError returns the error.
func (fm *FieldSpec) GetError() error {
	return fm.errors
}

// GetExample returns the example.
func (fm *FieldSpec) GetExample() string {
	return fm.example
}

// SetJsonName is to set JsonName.
func (fm *FieldSpec) SetJsonName(s string) {
	fm.jsonName = s
}

// GenerateFieldSpec return the instance of FieldSpec
func GenerateFieldSpec(f reflect.StructField, val interface{}, err error) FieldSpec {
	return FieldSpec{
		name:     f.Name,
		jsonName: f.Tag.Get(constants.JSON),
		tags:     f.Tag.Get(constants.Validate),
		value:    val,
		typ:      f.Type,
		errors:   err,
		example:  f.Tag.Get(constants.Example),
	}
}
