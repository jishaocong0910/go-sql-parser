package ast

import "go-sql-parser/enum"

// DELETE语法
type I_DeleteSyntax interface {
	I_StatementSyntax
	M_775679059A28() *M_DeleteSyntax
}

type M_DeleteSyntax struct {
	I              I_DeleteSyntax
	TableReference I_TableReferenceSyntax
	Where          *WhereSyntax
	Hint           *HintSyntax
}

func (this *M_DeleteSyntax) M_775679059A28() *M_DeleteSyntax {
	return this
}

func ExtendDeleteSyntax(i I_DeleteSyntax) *M_DeleteSyntax {
	i.M_5CF6320E8474().ParenthesizeType = enum.ParenthesizeTypes.NOT_SUPPORT
	return &M_DeleteSyntax{I: i}
}
