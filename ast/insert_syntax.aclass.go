package ast

// INSERT语法
type I_InsertSyntax interface {
	I_StatementSyntax
	M_8548F7993C94() *M_InsertSyntax
}

type M_InsertSyntax struct {
	I                  I_InsertSyntax
	NameTableReference I_NameTableReferenceSyntax
	InsertColumnList   *InsertColumnListSyntax
	ValueListList      I_ValueListListSyntax
	Hint               *HintSyntax
}

func (this *M_InsertSyntax) M_8548F7993C94() *M_InsertSyntax {
	return this
}

func ExtendInsertSyntax(i I_InsertSyntax) *M_InsertSyntax {
	return &M_InsertSyntax{I: i}
}
