package ast

// 字段项
type I_ColumnItemSyntax interface {
	I_Syntax
	M_ColumnItemSyntax_() *M_ColumnItemSyntax
	TableAlias() string
	Column() string
	FullColumn() string
}

type M_ColumnItemSyntax struct {
	I I_ColumnItemSyntax
}

func (this *M_ColumnItemSyntax) M_ColumnItemSyntax_() *M_ColumnItemSyntax {
	return this
}

func ExtendColumnItemSyntax(i I_ColumnItemSyntax) *M_ColumnItemSyntax {
	return &M_ColumnItemSyntax{I: i}
}
