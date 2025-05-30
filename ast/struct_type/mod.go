package struct_type

import (
	"github.com/longfangsong/idl-parser/ast/typeref"
	"github.com/longfangsong/idl-parser/ast/utils"
	"github.com/oleiade/gomme"
)

type Annotation struct {
	Name   string            `json:"name"`
	Values map[string]string `json:"values"`
}

type Field struct {
	Annotation Annotation           `json:"annotation"`
	Type       typeref.BitFieldType `json:"type"`
	Name       string               `json:"name"`
}

type Struct struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

func (Struct) IsModuleContent() {}

func parseKVPairs(code string) gomme.Result[map[string]string, string] {
	return gomme.Map(gomme.SeparatedList0(
		gomme.SeparatedPair(
			gomme.Recognize(
				gomme.Pair(
					gomme.Alpha1[string](),
					gomme.Alphanumeric0[string](),
				)),
			utils.InEmpty(gomme.Token[string]("=")),
			gomme.Recognize(
				gomme.Pair(
					gomme.Alpha1[string](),
					gomme.Alphanumeric0[string](),
				)),
		),
		utils.InEmpty(gomme.Token[string](",")),
	),
		func(pairs []gomme.PairContainer[string, string]) (map[string]string, error) {
			values := make(map[string]string)
			for _, pair := range pairs {
				values[pair.Left] = pair.Right
			}
			return values, nil
		})(code)
}

func parseAnnotation(code string) gomme.Result[Annotation, string] {
	return gomme.Map(
		gomme.SeparatedPair(
			gomme.Preceded(
				gomme.Token[string]("@"),
				utils.Identifier,
			),
			gomme.Whitespace0[string](),
			gomme.Delimited(
				gomme.Token[string]("("),
				parseKVPairs,
				gomme.Token[string](")"),
			),
		),
		func(output gomme.PairContainer[string, map[string]string]) (Annotation, error) {
			return Annotation{
				Name:   output.Left,
				Values: output.Right,
			}, nil
		},
	)(code)
}

func parseField(code string) gomme.Result[Field, string] {
	var bitFieldParser gomme.Parser[string, typeref.BitFieldType] = typeref.ParseBitField
	var annotationParser gomme.Parser[string, Annotation] = parseAnnotation
	var emptyParser gomme.Parser[string, string] = utils.ParseEmpty1
	return gomme.Map(
		gomme.SeparatedPair(
			annotationParser,
			emptyParser,
			gomme.SeparatedPair(
				bitFieldParser,
				gomme.Whitespace1[string](),
				gomme.Recognize(
					gomme.Pair(gomme.Alpha1[string](), gomme.Alphanumeric0[string]()),
				),
			),
		),
		func(output gomme.PairContainer[
			Annotation,
			gomme.PairContainer[typeref.BitFieldType, string],
		]) (Field, error) {
			return Field{
				Annotation: output.Left,
				Type:       output.Right.Left,
				Name:       output.Right.Right,
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
