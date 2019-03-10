package metadata

import "reflect"

// IFieldMetadata contains all functions to get field metadata details
type IFieldMetadata interface {
	GetTag() string

	GetName() string

	GetJSONName() string

	GetValue() interface{}

	GetType() reflect.Type

	GetError() error
}

// FieldMetadata stores information about fields
type FieldMetadata struct {
	Name     string
	JSONName string
	Tags     string
	Value    interface{}
	Typ      reflect.Type
	Errors   error
	// Params         string
	// Kinds          reflect.Kind
}

// GetTag returns the validation tag.
func (fm *FieldMetadata) GetTag() string {
	return fm.Tags
}

// GetName returns the name.
func (fm *FieldMetadata) GetName() string {
	return fm.Name
}

// GetJSONName returns the json name.
func (fm *FieldMetadata) GetJSONName() string {
	return fm.JSONName
}

// GetValue returns the value.
func (fm *FieldMetadata) GetValue() interface{} {
	return fm.Value
}

// GetType returns the type.
func (fm *FieldMetadata) GetType() reflect.Type {
	return fm.Typ
}

// GetError returns the error.
func (fm *FieldMetadata) GetError() error {
	return fm.Errors
}

// GenerateFieldMetadata return the instance of FieldMetadata
func GenerateFieldMetadata(f reflect.StructField, val interface{}, err error) FieldMetadata {
	return FieldMetadata{
		Name:     f.Name,
		JSONName: f.Tag.Get("json"),
		Tags:     f.Tag.Get("validate"),
		Value:    val,
		Typ:      f.Type,
		Errors:   err,
	}
}
