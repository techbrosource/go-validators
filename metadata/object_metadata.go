package metadata

import "reflect"

// IMetadata contains all functions to get Metadata Details
type IMetadata interface {
	GetObject() ObjectMetadata

	GetFields() []FieldMetadata
}

// GetObject returns the object metadata.
func (m *Metadatas) GetObject() ObjectMetadata {
	return m.Object
}

// GetFields returns the object metadata.
func (m *Metadatas) GetFields() []FieldMetadata {
	return m.Fields
}

// Metadatas stores information about all types of objects
type Metadatas struct {
	Object ObjectMetadata
	Fields []FieldMetadata
}

// GenerateMetadata returns the instance of Metadata
func GenerateMetadata(objectMetadata ObjectMetadata, fieldMetadatas []FieldMetadata) Metadatas {
	return Metadatas{
		Object: objectMetadata,
		Fields: fieldMetadatas,
	}
}

// IObjectMetadata contains all functions to get Object Metadata details
type IObjectMetadata interface {
	GetName() string

	GetType() string

	GetNumOfFields() int
}

// GetName returns the object name.
func (m *ObjectMetadata) GetName() string {
	return m.Name
}

// GetType returns the object type.
func (m *ObjectMetadata) GetType() string {
	return m.Type
}

// GetNumOfFields returns the number of fields.
func (m *ObjectMetadata) GetNumOfFields() int {
	return m.NumOfFields
}

// ObjectMetadata stores information about object
type ObjectMetadata struct {
	Name        string
	Type        string
	NumOfFields int
}

// GenerateObjectMetadata returns the instance of ObjectMetadata
func GenerateObjectMetadata(t reflect.Type) ObjectMetadata {
	return ObjectMetadata{
		Name:        t.Name(),
		Type:        t.Kind().String(),
		NumOfFields: t.NumField(),
	}
}
