package typeref

import "github.com/oleiade/gomme"

type ByteType struct{}

func (ByteType) isTypeRef() {}

func ParseByte(code string) gomme.Result[ByteType, string] {
	return gomme.Map(
		gomme.Token[string]("byte"),
		func(_ string) (ByteType, error) { return ByteType{}, nil },
	)(code)
}
