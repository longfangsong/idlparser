package struct_type

import (
	"github.com/longfangsong/idl-parser/ast/typeref"
	"github.com/longfangsong/idl-parser/ast/utils"
	"github.com/oleiade/gomme"
)

type Field struct {
	Type typeref.BitFieldType `json:"type"`
	Name string               `json:"name"`
}

type Struct struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

func (Struct) IsModuleContent() {}

func parseField(code string) gomme.Result[Field, string] {
	var bitFieldParser gomme.Parser[string, typeref.BitFieldType] = typeref.ParseBitField
	return gomme.Map(
		gomme.SeparatedPair(
			bitFieldParser,
			gomme.Whitespace1[string](),
			gomme.Recognize(
				gomme.Pair(gomme.Alpha1[string](), gomme.Alphanumeric0[string]()),
			),
		),
		func(output gomme.PairContainer[typeref.BitFieldType, string]) (Field, error) {
			return Field{
				Type: output.Left,
				Name: output.Right,
			}, nil
		},
	)(code)
}

func Parse(code string) gomme.Result[Struct, string] {
	bitsetTokenResult := gomme.Token[string]("struct")(code)
	if bitsetTokenResult.Err != nil {
		return gomme.Failure[string, Struct](bitsetTokenResult.Err, code)
	}
	nameResult :=
		utils.InEmpty(
			gomme.Recognize(gomme.Pair(gomme.Alpha1[string](), gomme.Alphanumeric0[string]())),
		)(bitsetTokenResult.Remaining)
	if nameResult.Err != nil {
		return gomme.Failure[string, Struct](nameResult.Err, code)
	}
	fieldsResult := utils.InEmpty(
		gomme.Delimited(
			utils.InEmpty(gomme.Token[string]("{")),
			gomme.SeparatedList0(parseField, utils.InEmpty(gomme.Token[string](";"))),
			gomme.Pair(
				gomme.Optional(utils.InEmpty(gomme.Token[string](";"))),
				utils.InEmpty(gomme.Token[string]("}")),
			),
		))(nameResult.Remaining)
	if fieldsResult.Err != nil {
		return gomme.Failure[string, Struct](fieldsResult.Err, code)
	}
	return gomme.Success(
		Struct{
			Name:   nameResult.Output,
			Fields: fieldsResult.Output,
		},
		fieldsResult.Remaining,
	)
}
