package reflect

type Member interface {
	GetDeclaringType() *Type
	GetName() string
	GetModifiers() int
}
