package ast

// 值列表的列表
type I_ValueListListSyntax interface {
	I_ListSyntax[*ValueListSyntax]
	M_ValueListListSyntax_() *M_ValueListListSyntax
}

type M_ValueListListSyntax struct {
	i I_ValueListListSyntax
}

func (this *M_ValueListListSyntax) M_ValueListListSyntax_() *M_ValueListListSyntax {
	return this
}

func ExtendValueListListSyntax(i I_ValueListListSyntax) *M_ValueListListSyntax {
	return &M_ValueListListSyntax{i: i}
}
