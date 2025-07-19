package ast

// GROUP BY语法
type GroupBySyntax_ interface {
	GroupBySyntax_() *GroupBySyntax__
	Syntax_
}

type GroupBySyntax__ struct {
	I                GroupBySyntax_
	OrderingItemList *OrderingItemListSyntax
}

func (this *GroupBySyntax__) GroupBySyntax_() *GroupBySyntax__ {
	return this
}

func (this *GroupBySyntax__) accept(v_ Visitor_) {
	v_.visitor_().visitGroupBySyntax__(this)
}

func ExtendGroupBySyntax(i GroupBySyntax_) *GroupBySyntax__ {
	return &GroupBySyntax__{I: i}
}
