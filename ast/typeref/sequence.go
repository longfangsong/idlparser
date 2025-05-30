package typeref

import (
	"github.com/longfangsong/idl-parser/ast/utils"
	"github.com/oleiade/gomme"
)

type Sequence struct {
	InnerType TypeRef `json:"inner_type"`
}

func (Sequence) isTypeRef() {}

func ParseSequence(code string) gomme.Result[Sequence, string] {
	return gomme.Map(
		gomme.Preceded(
			gomme.Token[string]("sequence"),
			utils.InEmpty(gomme.Delimited(
				gomme.Token[string]("<"),
				utils.InEmpty(ParseTypeRef),
				gomme.Token[string](">"),
			))),
		func(innerType TypeRef) (Sequence, error) {
			return Sequence{InnerType: innerType}, nil
		},
	)(code)
}
