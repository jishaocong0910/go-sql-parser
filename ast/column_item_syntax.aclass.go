package ast

// 字段项
type ColumnItemSyntax_ interface {
	ColumnItemSyntax_() *ColumnItemSyntax__
	Syntax_

	TableAlias() string
	Column() string
	FullColumn() string
}

type ColumnItemSyntax__ struct {
	I ColumnItemSyntax_
}

func (this *ColumnItemSyntax__) ColumnItemSyntax_() *ColumnItemSyntax__ {
	return this
}

func ExtendColumnItemSyntax(i ColumnItemSyntax_) *ColumnItemSyntax__ {
	return &ColumnItemSyntax__{I: i}
}
