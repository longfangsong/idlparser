package typeref

import "github.com/oleiade/gomme"

type BitFieldType struct {
	Width uint8 `json:"width"`
}

func (BitFieldType) isTypeRef() {}

func ParseBitField(code string) gomme.Result[BitFieldType, string] {
	return gomme.Map(
		gomme.Pair(
			gomme.Token[string]("bitfield"),
			gomme.Delimited(
				gomme.Token[string]("<"),
				gomme.UInt8[string](),
				gomme.Token[string](">"),
			),
		),
		func(pair gomme.PairContainer[string, uint8]) (BitFieldType, error) {
			return BitFieldType{Width: uint8(pair.Right)}, nil
		},
	)(code)
}
