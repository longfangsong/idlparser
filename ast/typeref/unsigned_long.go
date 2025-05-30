package typeref

import "github.com/oleiade/gomme"

type UnsignedLongType struct{}

func (UnsignedLongType) isTypeRef() {}

func ParseUnsignedLong(code string) gomme.Result[UnsignedLongType, string] {
	return gomme.Map(
		gomme.SeparatedPair(
			gomme.Token[string]("unsigned"),
			gomme.Whitespace1[string](),
			gomme.Token[string]("long"),
		),
		func(_ gomme.PairContainer[string, string]) (UnsignedLongType, error) { return UnsignedLongType{}, nil },
	)(code)
}
