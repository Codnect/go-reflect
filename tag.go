package reflect

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
