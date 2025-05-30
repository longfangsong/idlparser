package typeref

import "github.com/oleiade/gomme"

type UnsignedLongLongType struct{}

func (UnsignedLongLongType) isTypeRef() {}

func ParseUnsignedLongLong(code string) gomme.Result[UnsignedLongLongType, string] {
	return gomme.Map(
		gomme.SeparatedPair(
			gomme.SeparatedPair(
				gomme.Token[string]("unsigned"),
				gomme.Whitespace1[string](),
				gomme.Token[string]("long"),
			),
			gomme.Whitespace1[string](),
			gomme.Token[string]("long"),
		),
		func(_ gomme.PairContainer[gomme.PairContainer[string, string], string]) (UnsignedLongLongType, error) {
			return UnsignedLongLongType{}, nil
		},
	)(code)
}
