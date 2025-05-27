package typeref

import "github.com/oleiade/gomme"

type ShortType struct{}

func (ShortType) isTypeRef() {}

func ParseShort(code string) gomme.Result[ShortType, string] {
	return gomme.Map(
		gomme.Token[string]("short"),
		func(_ string) (ShortType, error) { return ShortType{}, nil },
	)(code)
}
