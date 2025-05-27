package typeref

import "github.com/oleiade/gomme"

type LongType struct{}

func (LongType) isTypeRef() {}

func ParseLong(code string) gomme.Result[LongType, string] {
	return gomme.Map(
		gomme.Token[string]("long"),
		func(_ string) (LongType, error) { return LongType{}, nil },
	)(code)
}
