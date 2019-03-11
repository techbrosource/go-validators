package structs

import (
	"testing"

	constants "github.com/techbrosource/go-validators/constants"
	validate "github.com/techbrosource/go-validators/utils"
	. "gopkg.in/go-playground/assert.v1"
)

type NetsedStruct struct {
	ParentField1 string      `json:"parent_field1" validate:"string"`
	Child        ChildStruct `json:"child" validate:"struct"`
}

type NetsedArrayStruct struct {
	ParentField1 string        `json:"parent_field1" validate:"string"`
	Child        []ChildStruct `json:"child" validate:"struct_array"`
}

type ChildStruct struct {
	ChildField1 string `json:"child_field1" validate:"string"`
	ChildField2 int32  `json:"child_field2" validate:"number"`
}

func TestNestedStructRequired(t *testing.T) {
	childData := ChildStruct{
		ChildField1: "test",
	}

	parentData := NetsedStruct{
		Child: childData,
	}

	fe := validate.Validate(parentData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 2)
	Equal(t, fe[0].GetJSONName(), "parent_field1")
	Equal(t, fe[0].GetError().Error(), constants.EmptyString)
	Equal(t, fe[1].GetJSONName(), "child.child_field2")
	Equal(t, fe[1].GetError().Error(), constants.EmptyNumber)
}

func TestNestedArrayStructRequired(t *testing.T) {
	childData := ChildStruct{
		ChildField1: "test",
	}

	childDatas := []ChildStruct{childData, childData}

	parentData := NetsedArrayStruct{
		Child: childDatas,
	}

	fe := validate.Validate(parentData)
	NotEqual(t, len(fe), 0)
	Equal(t, len(fe), 3)
	Equal(t, fe[0].GetJSONName(), "parent_field1")
	Equal(t, fe[0].GetError().Error(), constants.EmptyString)
	Equal(t, fe[1].GetJSONName(), "child[0].child_field2")
	Equal(t, fe[1].GetError().Error(), constants.EmptyNumber)
	Equal(t, fe[2].GetJSONName(), "child[1].child_field2")
	Equal(t, fe[2].GetError().Error(), constants.EmptyNumber)
}
