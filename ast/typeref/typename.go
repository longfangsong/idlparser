package typeref

import "github.com/oleiade/gomme"

type TypeName struct {
	Name string `json:"name"`
}

func (TypeName) isTypeRef() {}

func ParseTypeName(name string) gomme.Result[TypeName, string] {
	return gomme.Map(
		(gomme.Recognize(
			gomme.Pair(
				gomme.Alpha1[string](),
				gomme.Alphanumeric0[string](),
			),
		)),
		func(_ string) (TypeName, error) { return TypeName{Name: name}, nil },
	)(name)
}
