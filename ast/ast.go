package ast

import (
	"github.com/longfangsong/idl-parser/ast/bitset"
	"github.com/longfangsong/idl-parser/ast/struct_type"
	"github.com/longfangsong/idl-parser/ast/utils"
	"github.com/oleiade/gomme"
)

type ModuleContent interface {
	IsModuleContent()
}

type Module struct {
	Name    string          `json:"name"`
	Content []ModuleContent `json:"content"`
}

func Parse(code string) gomme.Result[Module, string] {
	bitsetTokenResult := gomme.Token[string]("module")(code)
	if bitsetTokenResult.Err != nil {
		return gomme.Failure[string, Module](bitsetTokenResult.Err, code)
	}
	nameResult :=
		utils.InEmpty(
			gomme.Recognize(gomme.Pair(gomme.Alpha1[string](), gomme.Alphanumeric0[string]())),
		)(bitsetTokenResult.Remaining)
	if nameResult.Err != nil {
		return gomme.Failure[string, Module](nameResult.Err, code)
	}
	contentResult := gomme.Delimited(
		utils.InEmpty(gomme.Token[string]("{")),
		gomme.Many0(utils.InEmpty(
			gomme.Terminated(gomme.Alternative(
				gomme.Map(bitset.Parse, func(output bitset.BitSet) (ModuleContent, error) { return output, nil }),
				gomme.Map(struct_type.Parse, func(output struct_type.Struct) (ModuleContent, error) { return output, nil }),
			),
				gomme.Optional(utils.InEmpty(gomme.Token[string](";"))),
			),
		)),
		utils.InEmpty(gomme.Token[string]("}")),
	)(nameResult.Remaining)
	if contentResult.Err != nil {
		return gomme.Failure[string, Module](contentResult.Err, code)
	}
	return gomme.Success(Module{
		Name:    nameResult.Output,
		Content: contentResult.Output,
	}, contentResult.Remaining)
}
