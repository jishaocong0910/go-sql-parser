package ast

// 字段项
type I_ColumnItemSyntax interface {
	I_Syntax
	M_2CCD6C894F80() *M_ColumnItemSyntax
	TableAlias() string
	Column() string
	FullColumn() string
}

type M_ColumnItemSyntax struct {
	I I_ColumnItemSyntax
}

func (this *M_ColumnItemSyntax) M_2CCD6C894F80() *M_ColumnItemSyntax {
	return this
}

func ExtendColumnItemSyntax(i I_ColumnItemSyntax) *M_ColumnItemSyntax {
	return &M_ColumnItemSyntax{I: i}
}
