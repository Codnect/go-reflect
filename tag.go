package reflect

import "github.com/codnect/go-one"

type Tag interface {
	GetName() string
	GetValue() one.One
}

type TaggedElement interface {
	IsTagPresent(tag Tag) bool
	IsTagPresentByName(name string) bool
	GetTag(tag Tag) (Tag, error)
	GetTags() []Tag
	GetTagByName(name string) (Tag, error)
}

type BaseTaggedElement struct {
	TaggedElement
}

func (t BaseTaggedElement) IsTagPresent(tag Tag) bool {
	_, err := t.GetTag(tag)
	if err != nil {
		return false
	}
	return true
}

func (t BaseTaggedElement) IsTagPresentByName(name string) bool {
	_, err := t.GetTagByName(name)
	if err != nil {
		return false
	}
	return true
}

type FieldTag struct {
	name  string
	value string
}

func newFieldTag(name string, value string) FieldTag {
	return FieldTag{
		name:  name,
		value: value,
	}
}

func (tag FieldTag) GetName() string {
	return tag.name
}

func (tag FieldTag) GetValue() one.One {
	return tag.value
}

/*
GetBool()
GetByte()
GetInt8()
GetShort()
GetInt16()
GetInt()
GetInt32()
GetInt64()
GetString()
GetFloat32()
GetFloat64()
GetUint8()
GetUInt16()
GetUInt()
GetUInt32()
GetUInt64()
*/
