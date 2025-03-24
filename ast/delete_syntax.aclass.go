package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// DELETE语法
type I_DeleteSyntax interface {
	I_StatementSyntax
	M_DeleteSyntax_() *M_DeleteSyntax
}

type M_DeleteSyntax struct {
	I              I_DeleteSyntax
	TableReference I_TableReferenceSyntax
	Where          *WhereSyntax
	Hint           *HintSyntax
}

func (this *M_DeleteSyntax) M_DeleteSyntax_() *M_DeleteSyntax {
	return this
}

func ExtendDeleteSyntax(i I_DeleteSyntax) *M_DeleteSyntax {
	i.M_Syntax_().ParenthesizeType = enum.ParenthesizeTypes.NOT_SUPPORT
	return &M_DeleteSyntax{I: i}
}
