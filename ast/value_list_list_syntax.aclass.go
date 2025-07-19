package ast

// 值列表的列表
type ValueListListSyntax_ interface {
	ValueListListSyntax_() *ValueListListSyntax__
	ListSyntax_[*ValueListSyntax]
}

type ValueListListSyntax__ struct {
	i ValueListListSyntax_
}

func (this *ValueListListSyntax__) ValueListListSyntax_() *ValueListListSyntax__ {
	return this
}

func ExtendValueListListSyntax(i ValueListListSyntax_) *ValueListListSyntax__ {
	return &ValueListListSyntax__{i: i}
}
