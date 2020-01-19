package reflect

type Type struct {
	name   string
	fields []Field
}

func newType() *Type {
	return &Type{}
}

func (t *Type) GetName() string {
	return t.name
}

func (t *Type) setFields(fields []Field) {
	t.fields = fields
}

func (t *Type) GetFields() []Field {
	return nil
}
