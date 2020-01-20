package reflect

import (
	"errors"
	"strings"
)

type Tag interface {
	GetTagName() string
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
	Name  string
	Value string
}

func (tag FieldTag) GetTagName() string {
	return tag.Name
}

type TagValueParser interface {
	Parse(value string) error
}

type Json struct {
	Value     string
	OmitEmpty bool
	Ignore    bool
}

func (tag Json) GetTagName() string {
	return "json"
}

func (tag Json) Parse(value string) error {
	value = strings.Trim(value, " ")
	if value == "" {
		return nil
	}
	result := strings.Split(value, ",")
	if len(result) == 1 {
		if strings.Trim(result[0], " ") == "-" {
			tag.Ignore = true
		} else {
			tag.Value = strings.Trim(result[0], " ")
		}
	} else if len(result) == 2 {
		tag.Value = strings.Trim(result[0], " ")
		if strings.Trim(result[0], " ") == "omitempty" {
			tag.OmitEmpty = true
		} else {
			return errors.New("json tag has wrong format")
		}
	} else {
		return errors.New("json tag has wrong format")
	}
	return nil
}

type Yaml struct {
	Value string
}

func (tag Yaml) GetTagName() string {
	return "yaml"
}

func (tag Yaml) Parse(value string) error {
	value = strings.Trim(value, " ")
	tag.Value = value
	return nil
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
