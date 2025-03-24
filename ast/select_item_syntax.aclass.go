package ast

// 查询项
type I_SelectItemSyntax interface {
	I_Syntax
	M_SelectItemSyntax_() *M_SelectItemSyntax
}

type M_SelectItemSyntax struct {
	I I_SelectItemSyntax
}

func (this *M_SelectItemSyntax) M_SelectItemSyntax_() *M_SelectItemSyntax {
	return this
}

func ExtendSelectItemSyntax(i I_SelectItemSyntax) *M_SelectItemSyntax {
	return &M_SelectItemSyntax{I: i}
}
