package typeref

import "github.com/oleiade/gomme"

type TypeRef interface {
	isTypeRef()
}

func ParseTypeRef(code string) gomme.Result[TypeRef, string] {
	return gomme.Alternative(
		gomme.Map(ParseByte, func(byte ByteType) (TypeRef, error) { return byte, nil }),
		gomme.Map(ParseLong, func(long LongType) (TypeRef, error) { return long, nil }),
		gomme.Map(ParseShort, func(short ShortType) (TypeRef, error) { return short, nil }),
		gomme.Map(ParseBitField, func(bitfield BitFieldType) (TypeRef, error) { return bitfield, nil }),
		gomme.Map(ParseSequence, func(sequence Sequence) (TypeRef, error) { return sequence, nil }),
		gomme.Map(ParseTypeName, func(name TypeName) (TypeRef, error) { return name, nil }),
	)(code)
}
