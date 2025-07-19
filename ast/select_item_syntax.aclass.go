package ast

// 查询项
type SelectItemSyntax_ interface {
	SelectItemSyntax_() *SelectItemSyntax__
	Syntax_
}

type SelectItemSyntax__ struct {
	I SelectItemSyntax_
}

func (this *SelectItemSyntax__) SelectItemSyntax_() *SelectItemSyntax__ {
	return this
}

func ExtendSelectItemSyntax(i SelectItemSyntax_) *SelectItemSyntax__ {
	return &SelectItemSyntax__{I: i}
}
