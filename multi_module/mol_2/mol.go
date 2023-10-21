package mol_2

func NewModule(value string) *Module {
	return &Module{Field: value}
}

type Module struct {
	Field string
}
